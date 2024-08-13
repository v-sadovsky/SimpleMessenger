#!/usr/bin/env bash

# This script setups google/protobuf proto description
CUR_DIR=$(pwd)
VENDOR_PROTO_PATH="$CUR_DIR/vendor.protobuf"

git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf "$VENDOR_PROTO_PATH"/protobuf &&\
	  cd "$VENDOR_PROTO_PATH"/protobuf &&\
  	git sparse-checkout set --no-cone src/google/protobuf &&\
	  git checkout

mkdir -p "$VENDOR_PROTO_PATH"/google
mv "$VENDOR_PROTO_PATH"/protobuf/src/google/protobuf "$VENDOR_PROTO_PATH"/google
rm -rf "$VENDOR_PROTO_PATH"/protobuf