#!/bin/bash
set -euxo pipefail

_tag="${1?"tag is required"}"

_script_dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
_temp_dir="$(mktemp -d)"
trap "exit 1" HUP INT PIPE QUIT TERM
trap 'rm -rf $_temp_dir' EXIT

# update protobuf definitions
git clone \
    --depth 1 \
    --filter=blob:none \
    --sparse \
    git@github.com:pomerium/pomerium-console \
    "$_temp_dir"
(cd "$_temp_dir" && git fetch --depth 1 origin "$_tag" && git checkout "$_tag" && git sparse-checkout set pkg/pb)

_files=(
  activity_log
  devices
  external_data_sources
  key_chain
  namespaces
  policy
  route_health_check
  route_redirect_action
  routes
  settings
  types
  users
)

for file in "${_files[@]}"; do
    cp -f "$_temp_dir/pkg/pb/"${file}*.pb*.go "$_script_dir/../pb/"
done

cd "$_script_dir/.."
go mod tidy
