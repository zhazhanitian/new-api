#!/bin/bash
# 本地开发启动脚本（Go 直接运行，无需 Docker 构建）
# 用法：./dev.sh

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

echo "▶ 启动 PostgreSQL 和 Redis 容器..."
docker compose -f docker-compose.dev.yml up -d postgres redis

echo "▶ 等待 PostgreSQL 就绪..."
until docker exec new-api-dev-pg pg_isready -U root -d new-api -q 2>/dev/null; do
  sleep 1
done

echo "▶ 停止 Docker 中的 new-api 容器（如果在运行）..."
docker compose -f docker-compose.dev.yml stop new-api 2>/dev/null || true

echo "▶ 编译并启动 Go 后端..."
export SQL_DSN="postgresql://root:123456@localhost:5432/new-api"
export REDIS_CONN_STRING="redis://localhost:6379"
export TZ="Asia/Shanghai"
export BATCH_UPDATE_ENABLED="true"
export PORT="9006"

go run main.go
