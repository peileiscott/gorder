#!/usr/bin/env bash

set -e

if ! [[ $0 =~ scripts/genopenapi.sh ]]; then
  echo "Please run this script from the root of the repository"
  exit 255
fi

# For more supported servers, please refer to https://github.com/oapi-codegen/oapi-codegen
readonly server=(
  "gin-server"
#  "echo-server"
#  "chi-server"
)

if [ ${#server[@]} -ne 1 ]; then
  echo "Can only provide one server implementation, please check scripts/genopenapi.sh"
  exit 255
fi

readonly service="$1"
readonly api_file="api/openapi/$service.yaml"
readonly api_out_dir="internal/$service/ports"
readonly client_out_dir="internal/common/client/$service"

mkdir -p "$api_out_dir" "$client_out_dir"

oapi-codegen -generate types -o "$api_out_dir/openapi_types.gen.go" -package ports "$api_file"
oapi-codegen -generate "${server[0]}" -o "$api_out_dir/openapi_api.gen.go" -package ports "$api_file"

oapi-codegen -generate types -o "$client_out_dir/openapi_types.gen.go" -package "$service" "$api_file"
oapi-codegen -generate client -o "$client_out_dir/openapi_client.gen.go" -package "$service" "$api_file"

echo "Successfully generated openapi files for $service service"