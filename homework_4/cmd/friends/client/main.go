package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/v_sadovsky/simple_messenger/homework_4/pkg/api/profiles"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewProfilesServiceClient(conn)

	// /SaveProfile
	{
		resp, err := cli.SaveProfile(context.Background(), &pb.SaveProfileRequest{
			Info: &pb.UserInfo{
				Name:     "client_name",
				Email:    "client_email@gmail.com",
				Password: "very strong password",
			},
		})
		if err != nil {
			log.Fatalf("SaveProfile error: %v", err)
		} else {
			log.Printf("Profile id is %d", resp.GetId())
		}
	}

	// GetProfile
	{
		resp, err := cli.GetProfile(context.Background(), &pb.GetProfileRequest{Id: 1})
		if err != nil {
			log.Fatalf("GetProfile error: %v", err)
		} else {
			notes, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf(" protojson.Marshal error: %v", err)
			} else {
				log.Printf("profile: %s", string(notes))
			}
		}
	}

	// ListProfiles
	{
		resp, err := cli.ListProfiles(context.Background(), &pb.ListProfilesRequest{})
		if err != nil {
			log.Fatalf("ListProfiles error: %v", err)
		} else {
			notes, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf(" protojson.Marshal error: %v", err)
			} else {
				log.Printf("profiles: %s", string(notes))
			}
		}
	}
}
