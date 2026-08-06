# iMac 本地构建镜像并部署到百度智能云 Ubuntu 服务器

> 适用场景：本地 iMac 构建 `new-api` Docker 镜像，压缩后上传到百度智能云 Ubuntu 服务器，生产服务器只负责导入镜像和重启服务，不再拉代码、不在服务器上构建。

## 一、前提确认

本地 iMac 需要安装并启动 Docker Desktop。

生产服务器需要已经安装 Docker 和 Docker Compose，并且已有项目部署目录：

```bash
/www/wwwroot/newapi/new-api-main
```

本文默认：

```bash
# 本地项目目录
/Users/biaodi/workspace/company/whaleClub/mapiV2/new-api

# 生产服务器项目目录
/www/wwwroot/newapi/new-api-main

# 生产服务端口
9006
```

如果实际路径或端口不同，按服务器实际情况替换。

镜像名沿用服务器历史配置：

```bash
new-api-new-api:latest
```

这样本地构建和海外服务器构建都交付同一个镜像名，生产服务器的 `docker-compose.yml` 不需要在两套流程之间来回修改。

## 二、本地拉取最新代码

在 iMac 终端执行：

```bash
cd /Users/biaodi/workspace/company/whaleClub/mapiV2/new-api
git pull
```

确认当前代码就是要发布的版本：

```bash
git status
```

如果有本地未提交修改，也会一起进入镜像。发布前请确认这些修改是你想发布的内容。

## 三、本地构建 Ubuntu 服务器可运行的镜像

百度智能云 Ubuntu 服务器通常是 `linux/amd64` 架构。iMac 如果是 M 系列芯片，默认会构建 `arm64` 镜像，服务器可能无法运行，所以必须指定 `--platform linux/amd64`。

```bash
cd /Users/biaodi/workspace/company/whaleClub/mapiV2/new-api

# 生成一个本次发布 tag，方便回滚和排查
IMAGE_TAG=$(date +%Y%m%d-%H%M)

# 构建 linux/amd64 镜像，并加载到本地 Docker
docker buildx build \
  --platform linux/amd64 \
  -t new-api-new-api:latest \
  -t new-api:${IMAGE_TAG} \
  --load \
  .
```

构建完成后确认镜像：

```bash
docker images | grep new-api
```

应该能看到类似：

```text
new-api-new-api   latest          ...
new-api           20260731-1130   ...
```

## 四、本地导出并压缩镜像

为了减少上传时间，建议直接导出成 `tar.gz`：

```bash
cd /Users/biaodi/workspace/company/whaleClub/mapiV2/new-api

docker save new-api-new-api:latest new-api:${IMAGE_TAG} | gzip > new-api-image.tar.gz

ls -lh new-api-image.tar.gz
```

如果你重新打开了一个终端，`IMAGE_TAG` 变量可能已经丢失，可以先用下面命令查到本次 tag：

```bash
docker images | grep new-api
```

然后手动替换导出命令里的 tag，例如：

```bash
docker save new-api-new-api:latest new-api:20260731-1130 | gzip > new-api-image.tar.gz
```

## 五、上传镜像到百度智能云 Ubuntu 服务器

把 `生产服务器IP` 替换为真实 IP：

```bash
scp ./new-api-image.tar.gz root@生产服务器IP:/www/wwwroot/newapi/new-api-main/
```

如果服务器 SSH 不是默认 22 端口，例如端口是 `2222`：

```bash
scp -P 2222 ./new-api-image.tar.gz root@生产服务器IP:/www/wwwroot/newapi/new-api-main/
```

## 六、确认生产服务器 docker-compose.yml

登录百度智能云 Ubuntu 服务器：

```bash
ssh root@生产服务器IP
cd /www/wwwroot/newapi/new-api-main
```

确认 `docker-compose.yml` 里的 `new-api` 服务使用历史镜像名：

```yaml
    image: new-api-new-api:latest
```

如果之前已经是这样，就不用改：

```yaml
services:
  new-api:
    image: new-api-new-api:latest
    container_name: new-api
    restart: always
    command: --log-dir /app/logs
```

如果还是构建模式，例如：

```yaml
    build:
      context: .
      dockerfile: Dockerfile
```

才需要改成 `image: new-api-new-api:latest`。其余配置不要动，尤其是 `ports`、`volumes`、`env_file`、`environment`、`depends_on`。

## 七、生产服务器导入镜像并重启

在百度智能云 Ubuntu 服务器执行：

```bash
cd /www/wwwroot/newapi/new-api-main

# 导入本地上传的镜像
docker load -i new-api-image.tar.gz

# 确认镜像已经存在
docker images | grep new-api

# 重启服务，不要加 --build
docker compose down
docker compose up -d
```

注意：这里不要执行 `docker compose up -d --build`，否则又会在服务器上重新构建。

## 八、检查服务

```bash
docker compose ps
curl http://127.0.0.1:9006/api/status
```

如果状态异常，查看日志：

```bash
docker compose logs -f new-api
```

## 九、下次更新流程

以后每次发布只需要重复下面几步。

本地 iMac：

```bash
cd /Users/biaodi/workspace/company/whaleClub/mapiV2/new-api
git pull

IMAGE_TAG=$(date +%Y%m%d-%H%M)

docker buildx build \
  --platform linux/amd64 \
  -t new-api-new-api:latest \
  -t new-api:${IMAGE_TAG} \
  --load \
  .

docker save new-api-new-api:latest new-api:${IMAGE_TAG} | gzip > new-api-image.tar.gz

scp ./new-api-image.tar.gz root@生产服务器IP:/www/wwwroot/newapi/new-api-main/
```

生产服务器：

```bash
ssh root@生产服务器IP
cd /www/wwwroot/newapi/new-api-main

docker load -i new-api-image.tar.gz
docker compose down
docker compose up -d
docker compose ps
curl http://127.0.0.1:9006/api/status
```

## 十、常见问题

### 1. iMac 构建很慢

M 系列 iMac 构建 `linux/amd64` 镜像时会走跨架构构建，比原生 `arm64` 慢一些，但通常仍然比把代码上传到云端再构建更稳定。

### 2. 服务器提示 exec format error

一般是镜像架构不对。重新在本地构建，确保命令里有：

```bash
--platform linux/amd64
```

### 3. docker compose 还是在服务器上构建

检查生产服务器的 `docker-compose.yml`，`new-api` 服务里不能再有：

```yaml
build:
```

应该使用：

```yaml
image: new-api-new-api:latest
```

### 4. 想节省服务器磁盘

确认新版本运行正常后，可以清理悬空镜像：

```bash
docker image prune
```

不要随便执行 `docker system prune -a --volumes`，它可能删除未使用镜像和数据卷。

### 5. 本地 iMac 镜像越积越多，如何清理

每次构建都会生成一个带日期 tag 的镜像（例如 `new-api:20260806-1004`），时间久了本地会积累很多旧版本。

**查看当前所有 new-api 相关镜像：**

```bash
docker images | grep new-api
```

**删除指定 tag（用于精确清理某几个版本）：**

```bash
docker rmi new-api:20260731-1150
docker rmi new-api:20260805-1214
```

注意：如果多个 tag 指向同一个镜像 ID（`docker images` 里 IMAGE ID 列相同），删除其中一个 tag 不会删除镜像本身，只有最后一个 tag 被删除时镜像才会真正释放磁盘。

**批量删除所有旧的日期 tag（保留 `new-api-new-api:latest` 和最新的日期 tag）：**

先确认当前最新的日期 tag：

```bash
docker images | grep '^new-api '
```

把你确认要保留的 tag 记下来，再批量删除其余的日期 tag：

```bash
# 删除所有 new-api:年份开头 的旧 tag（按需调整正则）
docker images --format '{{.Repository}}:{{.Tag}}' | grep '^new-api:2' | xargs docker rmi
```

**清理所有本地不再使用的镜像（悬空 + 未被任何容器引用的镜像）：**

```bash
# 只清理悬空镜像（无 tag 的 <none>）
docker image prune

# 清理所有未被容器使用的镜像（更彻底，操作前确认没有容器依赖它们）
docker image prune -a
```

日常建议：每次确认新版本在服务器上运行正常后，在本地删掉 2～3 个版本之前的旧日期 tag，保留最近 1～2 个版本用于应急回滚。
