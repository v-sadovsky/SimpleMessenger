#!/usr/bin/env bash

# This script removes all vendor *.proto files and empty directories
CUR_DIR=$(pwd)
VENDOR_PROTO_PATH="$CUR_DIR/vendor.protobuf"

find "$VENDOR_PROTO_PATH" -type f ! -name "*.proto" -delete
find "$VENDOR_PROTO_PATH" -empty -type d -delete