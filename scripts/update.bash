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
    --branch "$_tag" \
    git@github.com:pomerium/pomerium-console \
    "$_temp_dir"
(cd "$_temp_dir" && git sparse-checkout set pkg/pb)
cp -f "$_temp_dir/pkg/pb/"*.pb.go "$_script_dir/../pb/"

# update pomerium
cd "$_script_dir/.."
go get -u github.com/pomerium/pomerium@"$_tag"
go mod tidy
