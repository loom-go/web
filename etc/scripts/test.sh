#!/bin/sh

echo "Running wasm tests..."
env -i \
  HOME="$HOME" \
  PATH="$PATH" \
  GOROOT="$(go env GOROOT)" \
  GOPATH="$(go env GOPATH)" \
  GOCACHE="$(go env GOCACHE)" \
  GOOS=js \
  GOARCH=wasm \
  go test -v -exec="bash $(go env GOROOT)/lib/wasm/go_js_wasm_exec" $@
