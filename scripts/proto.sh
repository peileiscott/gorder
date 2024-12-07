#!/usr/bin/env bash

set -e

if [[ $1 =~ scripts/proto.sh ]]; then
  echo "Please run this script from the root of the repository"
  exit 1
fi

readonly service_name=$1
readonly out_dir="internal/common/genproto/${service_name}pb"

mkdir -p "$out_dir"

protoc \
  --proto_path="api/protobuf" \
  --go_out="$out_dir" --go_opt=paths=source_relative \
  --go-grpc_out="$out_dir" --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  "api/protobuf/$service_name.proto"

echo "Generated protobuf files for $service_name service in $out_dir"