#!/usr/bin/env bash

set -e

if [[ $1 =~ scripts/openapi.sh ]]; then
  echo "Please run this script from the root of the repository"
  exit 1
fi

# For all supported servers, please refer to https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#supported-servers
readonly server="gin-server"

readonly service_name=$1
readonly openapi_file="api/openapi/$service_name.yaml"
readonly api_out_dir="internal/$service_name/ports"
readonly client_out_dir="internal/common/client/$service_name"

mkdir -p "$api_out_dir" "$client_out_dir"

oapi-codegen -config "api/openapi/config.yaml" -generate $server -o "$api_out_dir/openapi_api.gen.go" -package ports "$openapi_file"
oapi-codegen -config "api/openapi/config.yaml" -generate client -o "$client_out_dir/openapi_client.gen.go" -package "$service_name" "$openapi_file"
oapi-codegen -config "api/openapi/config.yaml" -generate types -o "$client_out_dir/openapi_types.gen.go" -package "$service_name" "$openapi_file"

echo "Generated openapi files for $service_name service in $client_out_dir and $api_out_dir"