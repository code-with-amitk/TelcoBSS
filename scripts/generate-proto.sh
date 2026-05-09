#!/usr/bin/env bash
set -euo pipefail

PROTO_DIR="docs/api/proto"
OUT_DIR="pkg/api"

mkdir -p "$OUT_DIR"

protoc \
  --go_out="$OUT_DIR" \
  --go-grpc_out="$OUT_DIR" \
  --proto_path="$PROTO_DIR" \
  "$PROTO_DIR"/*.proto

echo "Generated Go protobuf stubs in $OUT_DIR"
