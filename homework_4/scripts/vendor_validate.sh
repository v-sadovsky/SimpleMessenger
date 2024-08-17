#!/usr/bin/env bash

# This script removes all vendor *.proto files and empty directories
CUR_DIR=$(pwd)
VENDOR_PROTO_PATH="$CUR_DIR/vendor.protobuf"

git clone -b main --single-branch --depth=1 --filter=tree:0 \
		https://github.com/bufbuild/protovalidate "$VENDOR_PROTO_PATH"/protovalidate && \
cd "$VENDOR_PROTO_PATH"/protovalidate && \
git checkout

mv "$VENDOR_PROTO_PATH"/protovalidate/proto/protovalidate/buf "$VENDOR_PROTO_PATH"
rm -rf "$VENDOR_PROTO_PATH"/protovalidate