# Render Build Script
#!/bin/bash
echo "Building BookMyGo backend..."
go mod tidy
go build -o main cmd/server/main.go
echo "Build completed!"
