#!/usr/bin/env bash

# This script prepares and run swagger-ui service
CUR_DIR=$(pwd)
LOCAL_BIN="$CUR_DIR/bin"
API_SCHEMA_PATH="$CUR_DIR/api/profiles"
SWAGGER_UI_PATH="$CUR_DIR/swaggerui"
SCHEMA_FILE_NAME="service.swagger.json"

find "$SWAGGER_UI_PATH" -type f -name "$SCHEMA_FILE_NAME" -delete
cp "$API_SCHEMA_PATH"/"$SCHEMA_FILE_NAME" "$SWAGGER_UI_PATH"
go run ./cmd/friends/swaggerui/main.go
