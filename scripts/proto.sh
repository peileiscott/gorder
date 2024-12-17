#!/usr/bin/env bash

set -e

if ! [[ $0 =~ scripts/proto.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

readonly service=$1
readonly out_path="internal/common/genproto/${service}pb"

mkdir -p "$out_path"

protoc \
  --proto_path=api/protobuf "api/protobuf/$service.proto" \
  --go_out="$out_path" --go_opt=paths=source_relative \
  --go-grpc_out="$out_path" --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false
