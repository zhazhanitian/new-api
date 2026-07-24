# New API

基于 [new-api](https://github.com/QuantumNous/new-api) 二次开发的 LLM 网关与 AI 资产管理系统。

---

## 本地开发

### 依赖

- Go
- [Bun](https://bun.sh/)（前端构建）
- Docker（运行 PostgreSQL & Redis）

### 启动

```bash
# 首次启动（含前端构建，耗时较长）
./dev.sh

# 前端未改动时跳过构建，加速启动
./dev.sh --skip-build
```

启动后访问：`http://localhost:9006`

脚本做了什么：
1. 用 `docker-compose.dev.yml` 启动 PostgreSQL 和 Redis
2. 构建前端（`web/classic`，使用 bun）
3. 用 `go run main.go` 启动后端，自动注入本地数据库/Redis 连接串

### 停止

```bash
./stop.sh
```

停止 Go 后端进程 + Docker 的 postgres/redis 容器。

---

## 部署文档

| 文档 | 说明 |
|------|------|
| [宝塔配置域名与反向代理](./docs/宝塔配置域名与反向代理.md) | 宝塔面板添加站点、申请 SSL、配置 Nginx 反向代理（含流式超时配置） |
| [海外构建镜像并部署到生产服务器](./docs/海外构建镜像并部署到生产服务器.md) | 在海外服务器构建 Docker 镜像 → 导出 tar → 上传到生产服务器 → 启动 |
| [数据库迁移（国内到海外）](./docs/数据库迁移（国内到海外）.md) | 将国内 PostgreSQL 数据完整迁移到海外服务器 |
| [价格配置说明](./docs/价格配置说明.md) | 按量/按次/表达式计费配置方法，含图像生成、音频、缓存价格示例 |

---

## 环境变量（常用）

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `SQL_DSN` | 数据库连接串（PostgreSQL / MySQL） | — |
| `REDIS_CONN_STRING` | Redis 连接串 | — |
| `SESSION_SECRET` | Session 密钥（多机部署必须设置） | — |
| `CRYPTO_SECRET` | 加密密钥（使用 Redis 时必须设置） | — |
| `PORT` | 监听端口 | `3000` |
| `TZ` | 时区 | — |
| `STREAMING_TIMEOUT` | 流式响应超时（秒） | `300` |
| `BATCH_UPDATE_ENABLED` | 批量写入开关 | `false` |

完整环境变量参考上游文档：[Environment Variables](https://docs.newapi.pro/en/docs/installation/config-maintenance/environment-variables)

---

## 上游项目

- [QuantumNous/new-api](https://github.com/QuantumNous/new-api)
- [songquanpeng/one-api](https://github.com/songquanpeng/one-api)
