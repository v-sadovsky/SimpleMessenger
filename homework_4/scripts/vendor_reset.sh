#!/usr/bin/env bash

# This script removes all vendor data
CUR_DIR=$(pwd)
VENDOR_PROTO_PATH="$CUR_DIR/vendor.protobuf"

rm -rf "$VENDOR_PROTO_PATH"
mkdir -p "$VENDOR_PROTO_PATH"