@echo off
setlocal

set BINARY_NAME=poca
set BUILD_DIR=%~dp0..\build
set CMD_DIR=%~dp0..\cmd\poca

if not exist "%BUILD_DIR%" mkdir "%BUILD_DIR%"

echo Building %BINARY_NAME%...
go build -o "%BUILD_DIR%\%BINARY_NAME%.exe" "%CMD_DIR%"
echo Built to %BUILD_DIR%\%BINARY_NAME%.exe
