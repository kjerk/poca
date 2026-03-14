#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

BINARY_NAME="poca"
BUILD_DIR="$ROOT_DIR/build"
CMD_DIR="$ROOT_DIR/cmd/poca"

mkdir -p "$BUILD_DIR"

echo "Building $BINARY_NAME..."
go build -o "$BUILD_DIR/$BINARY_NAME" "$CMD_DIR"
echo "Built to $BUILD_DIR/$BINARY_NAME"
