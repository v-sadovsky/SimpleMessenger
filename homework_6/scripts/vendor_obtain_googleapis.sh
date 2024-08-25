#!/usr/bin/env bash

# This script setups googleapis proto descriptions
CUR_DIR=$(pwd)
VENDOR_PROTO_PATH="$CUR_DIR/vendor.protobuf"

git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/googleapis/googleapis "$VENDOR_PROTO_PATH"/googleapis &&\
cd "$VENDOR_PROTO_PATH"/googleapis &&\
git checkout
mv "$VENDOR_PROTO_PATH"/googleapis/google "$VENDOR_PROTO_PATH"
rm -rf "$VENDOR_PROTO_PATH"/googleapis