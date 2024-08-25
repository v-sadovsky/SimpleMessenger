package main

import (
	"context"
	"github.com/v_sadovsky/simple_messenger/homework_6/internal/app/adapters/chats_repository"
	"log"

	"google.golang.org/grpc"

	"github.com/v_sadovsky/simple_messenger/homework_6/internal/app/server"
	"github.com/v_sadovsky/simple_messenger/homework_6/internal/app/usecase/chats"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// =========================
	// adapters
	// =========================

	// repository

	messagesRepo := chats_repository.NewRepository()

	// services

	//wmsAdapter := warehousemanagementsystem.NewClient()

	// =========================
	// usecases
	// =========================

	chatsUsecase := chats.NewUsecase(chats.Deps{
		MessagesRepository: messagesRepo,
	})

	// =========================
	// delivery
	// =========================

	config := server.Config{
		GRPCPort:               ":8082",
		GRPCGatewayPort:        ":8080",
		ChainUnaryInterceptors: []grpc.UnaryServerInterceptor{
			//middleware.ErrorsUnaryInterceptor(),
		},
	}

	srv, err := server.New(ctx, config, server.Deps{
		ChatsUsecase: chatsUsecase,
		// Dependency injection
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("run: %v", err)
	}
}
