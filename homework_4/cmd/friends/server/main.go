package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"sync/atomic"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/v_sadovsky/simple_messenger/homework_4/pkg/api/profiles"
)

var idSerial uint64

// server is used to implement pb.ProfilesServiceServer
type server struct {
	pb.UnimplementedProfilesServiceServer

	mx        sync.RWMutex
	profiles  map[uint64]*pb.UserInfo
	validator *protovalidate.Validator
}

func newServer() (*server, error) {
	srv := &server{
		profiles: make(map[uint64]*pb.UserInfo),
	}

	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&pb.SaveProfileRequest{},
			&pb.GetProfileRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize validator: %w", err)
	}

	srv.validator = validator

	return srv, nil
}

func protovalidateVialationsToGoogleViolations(vs []*validate.Violation) []*errdetails.BadRequest_FieldViolation {
	res := make([]*errdetails.BadRequest_FieldViolation, len(vs))
	for i, v := range vs {
		res[i] = &errdetails.BadRequest_FieldViolation{
			Field:       v.FieldPath,
			Description: v.Message,
		}
	}
	return res
}

func convertProtovalidateValidationErrorToErrdetailsBadRequest(valErr *protovalidate.ValidationError) *errdetails.BadRequest {
	return &errdetails.BadRequest{
		FieldViolations: protovalidateVialationsToGoogleViolations(valErr.Violations),
	}
}

func rpcValidationError(err error) error {
	if err == nil {
		return nil
	}

	var valErr *protovalidate.ValidationError
	if ok := errors.As(err, &valErr); ok {
		st, err := status.New(codes.InvalidArgument, codes.InvalidArgument.String()).
			WithDetails(convertProtovalidateValidationErrorToErrdetailsBadRequest(valErr))
		if err == nil {
			return st.Err()
		}
	}

	return status.Error(codes.Internal, err.Error())
}

// SaveProfile implements pb.ProfilesServiceServer
func (s *server) SaveProfile(_ context.Context, req *pb.SaveProfileRequest) (*pb.SaveProfileResponse, error) {
	info := req.GetInfo()
	log.Printf("SaveProfile: received: %s, %s", info.GetName(), info.GetEmail())

	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	id := atomic.AddUint64(&idSerial, 1)

	s.mx.Lock()
	s.profiles[id] = info
	s.mx.Unlock()

	return &pb.SaveProfileResponse{
		Id: id,
	}, nil
}

// GetProfile implements pb.ProfilesServiceServer
func (s *server) GetProfile(_ context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	log.Println("GetProfile: received")

	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	s.mx.RLock()
	defer s.mx.RUnlock()

	id := req.GetId()
	if _, ok := s.profiles[id]; !ok {
		return nil, status.Error(codes.NotFound, "profile not found")
	}

	return &pb.GetProfileResponse{
		Profile: &pb.Profile{
			Id:   id,
			Info: s.profiles[id],
		},
	}, nil
}

// ListProfiles implements pb.ProfilesServiceServer
func (s *server) ListProfiles(_ context.Context, req *pb.ListProfilesRequest) (*pb.ListProfilesResponse, error) {
	log.Println("ListProfiles: received")

	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	s.mx.RLock()
	defer s.mx.RUnlock()

	profiles := make([]*pb.Profile, 0, len(s.profiles))
	for id, info := range s.profiles {
		profiles = append(profiles, &pb.Profile{
			Id:   id,
			Info: info,
		})
	}

	return &pb.ListProfilesResponse{
		Profiles: profiles,
	}, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	server, err := newServer()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		grpcServer := grpc.NewServer()
		pb.RegisterProfilesServiceServer(grpcServer, server)

		reflection.Register(grpcServer)

		lis, err := net.Listen("tcp", ":8082") // gRPR
		if err != nil {
			log.Fatalf("failed to listen gRPC server: %v", err)
		}

		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mux := runtime.NewServeMux()
		if err = pb.RegisterProfilesServiceHandlerServer(ctx, mux, server); err != nil {
			log.Fatalf("failed to serve REST: %v", err)
		}

		corsMux := enableCORS(mux)
		httpServer := &http.Server{
			Handler: corsMux,
		}

		lis, err := net.Listen("tcp", ":8080") // HTTP
		if err != nil {
			log.Fatalf("failed to listen REST server: %v", err)
		}

		// Start HTTP server (and proxy calls to gRPC server endpoint)
		log.Printf("REST server listening at %v", lis.Addr())
		if err := httpServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve REST: %v", err)
		}
	}()

	wg.Wait()
}

// Middleware to add CORS headers
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
