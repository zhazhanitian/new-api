#!/bin/bash
# 本地开发启动脚本（Go 直接运行，无需 Docker 构建）
# 用法：./dev.sh
# ./dev.sh --skip-build  跳过前端构建（前端代码未改动时可加速启动）

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

export PATH="$HOME/.bun/bin:$PATH"

# 解析参数
SKIP_BUILD=false
for arg in "$@"; do
  case $arg in
    --skip-build) SKIP_BUILD=true ;;
  esac
done

echo "▶ 启动 PostgreSQL 和 Redis 容器..."
docker compose -f docker-compose.dev.yml up -d postgres redis

echo "▶ 等待 PostgreSQL 就绪..."
until docker exec new-api-dev-pg pg_isready -U root -d new-api -q 2>/dev/null; do
  sleep 1
done

echo "▶ 停止 Docker 中的 new-api 容器（如果在运行）..."
docker compose -f docker-compose.dev.yml stop new-api 2>/dev/null || true

# 构建前端（Go embed 需要 dist/ 存在才能编译）
if [ "$SKIP_BUILD" = false ]; then
  echo "▶ 构建前端（web/classic）..."
  cd "$SCRIPT_DIR/web/classic"
  bun run build
  cd "$SCRIPT_DIR"
  echo "✔ 前端构建完成"
else
  echo "⚡ 跳过前端构建（--skip-build）"
fi

echo "▶ 编译并启动 Go 后端..."
export SQL_DSN="postgresql://root:123456@localhost:5432/new-api"
export REDIS_CONN_STRING="redis://localhost:6379"
export TZ="Asia/Shanghai"
export BATCH_UPDATE_ENABLED="true"
export PORT="9006"

go run main.go &
BACKEND_PID=$!

echo "▶ 等待 Go 后端就绪（:9006）..."
until curl -sf http://localhost:9006/api/status >/dev/null 2>&1; do
  sleep 1
  kill -0 $BACKEND_PID 2>/dev/null || { echo "✘ 后端启动失败，请检查日志"; exit 1; }
done

echo "✅ 全部就绪！访问 http://localhost:9006"

# Ctrl+C 时关闭后端
trap "echo '▶ 停止服务...'; kill $BACKEND_PID 2>/dev/null; exit" INT TERM

wait $BACKEND_PID
