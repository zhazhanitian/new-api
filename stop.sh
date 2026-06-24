#!/bin/bash
# 本地开发停止脚本
# 用法：./stop.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

echo "▶ 停止 Go 后端进程..."
pkill -f "go run main.go" 2>/dev/null || true
pkill -f "new-api" 2>/dev/null || true

echo "▶ 停止前端开发服务器..."
pkill -f "bun run dev" 2>/dev/null || true
pkill -f "vite" 2>/dev/null || true

echo "▶ 停止 Docker 容器..."
docker compose -f "$SCRIPT_DIR/docker-compose.dev.yml" stop postgres redis 2>/dev/null || true

echo "✅ 所有服务已停止"
