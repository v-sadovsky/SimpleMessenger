#!/usr/bin/env bash

# This script generates .go files with protoc

CUR_DIR=$(pwd)
LOCAL_BIN="$CUR_DIR/bin"
PROTOC="PATH=\"\$PATH:$LOCAL_BIN\" protoc"     # Добавляем bin в текущей директории в PATH при запуске protoc
PROTO_PATH="$CUR_DIR/api"                      # Путь к protobuf файлам
PKG_PROTO_PATH="$CUR_DIR/pkg"                  # Путь к сгенеренным .pb.go файлам
VENDOR_PROTO_PATH="$CUR_DIR/vendor.protobuf"   # Путь к завендореным protobuf файлам

mkdir -p "$PKG_PROTO_PATH"

eval "$PROTOC" -I "$VENDOR_PROTO_PATH" --proto_path="$CUR_DIR" \
	--go_out="$PKG_PROTO_PATH" --go_opt paths=source_relative \
	--go-grpc_out="$PKG_PROTO_PATH" --go-grpc_opt paths=source_relative \
	--grpc-gateway_out="$PKG_PROTO_PATH" --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	"$PROTO_PATH"/profiles/messages.proto \
	"$PROTO_PATH"/profiles/service.proto

eval "$PROTOC" -I "$VENDOR_PROTO_PATH" --proto_path="$CUR_DIR" \
  --openapiv2_out=. --openapiv2_opt logtostderr=true \
  "$PROTO_PATH"/profiles/service.proto
