#!/usr/bin/env bash

# This script installs necessary binary dependencies

	echo "Installing binary dependencies..."

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest   # to generate models
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest  # to generate client and server
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/bufbuild/buf/cmd/buf@v1.32.2
	go install github.com/yoheimuta/protolint/cmd/protolint@latest