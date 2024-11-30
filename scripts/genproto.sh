#!/usr/bin/env bash

set -e

if ! [[ $0 =~ scripts/genproto.sh ]]; then
  echo "Please run this script from the root of the repository"
  exit 255
fi

readonly service="$1"
readonly out_dir="internal/common/genproto/${service}pb"

mkdir -p "$out_dir"

protoc \
  --proto_path="api/protobuf" \
  --go_out="$out_dir" --go_opt=paths=source_relative \
  --go-grpc_out="$out_dir" --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  "api/protobuf/$service.proto"

echo "Successfully generated protobuf files for $service service"