@echo off
setlocal

echo Running go vet...
go vet ./... || exit /b 1

echo Running tests...
go test ./... || exit /b 1

echo Running golangci-lint...
golangci-lint run ./... || exit /b 1

echo All checks passed.
