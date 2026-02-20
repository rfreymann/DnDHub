#!/usr/bin/env bash
set -euo pipefail

# Default values
export PORT=${PORT:-8080}
export DATABASE_URL=${DATABASE_URL:-"postgres://myappuser:secretpw@localhost:15432/myappdb?sslmode=disable"}
export JWT_SECRET=${JWT_SECRET:-"supersecretlongrandom"}

echo "Starting DnDHub API on port $PORT..."
go run ./cmd/api
