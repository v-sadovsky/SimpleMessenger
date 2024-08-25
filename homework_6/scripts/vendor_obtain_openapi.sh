#!/usr/bin/env bash

# This script setups openapiv2 proto options
CUR_DIR=$(pwd)
VENDOR_PROTO_PATH="$CUR_DIR/vendor.protobuf"

git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
https://github.com/grpc-ecosystem/grpc-gateway "$VENDOR_PROTO_PATH"/grpc-gateway && \
cd "$VENDOR_PROTO_PATH"/grpc-gateway && \
git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
git checkout
mkdir -p "$VENDOR_PROTO_PATH"/protoc-gen-openapiv2
mv "$VENDOR_PROTO_PATH"/grpc-gateway/protoc-gen-openapiv2/options "$VENDOR_PROTO_PATH"/protoc-gen-openapiv2
rm -rf "$VENDOR_PROTO_PATH"/grpc-gateway