#!/usr/bin/env bash
set -euo pipefail

echo "Running go vet..."
go vet ./...

echo "Running tests..."
go test ./...

echo "Running golangci-lint..."
golangci-lint run ./...

echo "All checks passed."
