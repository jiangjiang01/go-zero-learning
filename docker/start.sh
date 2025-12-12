#!/bin/bash

# 开启严格 bash 模式：脚本遇错即停（-e），引用未定义变量时报错（-u），管道一出错即失败（-o pipefail）；可提高脚本健壮性与安全性
set -euo pipefail

echo "🚀 启动 Go-Zero Learning 项目..."

# 确保在项目根目录执行（避免从其他目录执行导致 docker-compose 相对路径失效）
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
cd "${PROJECT_ROOT}"

# 检查 .env 文件是否存在
if [ ! -f .env ]; then
    echo "❌ 错误：.env 文件不存在！"
    echo ""
    echo "请按照以下步骤操作："
    echo "  1. 复制 .env.example 为 .env："
    echo "     cp .env.example .env"
    echo "  2. 编辑 .env 文件，设置正确的环境变量值"
    echo "  3. 重新运行启动脚本"
    echo ""
    exit 1
fi

# 检查 Docker 是否运行
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker 未运行，请先启动 Docker。"
    exit 1
fi

# 构建并启动服务
echo "📦 构建镜像并启动服务..."
docker compose up -d --build

# 等待服务启动
echo "等待服务启动"
# 最长等 60 秒
for i in {1..12}; do
  if docker compose ps --status running | grep -q "healthy"; then
    break
  fi
  sleep 5
done

# 检查服务状态
echo ""
echo "服务状态："
docker compose ps

echo ""
echo "服务启动完成！"
# 初始化数据库
echo ""
echo "初始化数据库"
echo "docker exec -i -e MYSQL_PWD=123456 go-zero-mysql mysql -uroot --default-character-set=utf8mb4 testdb < scripts/init_test_data.sql"
echo "--------------------------------"
echo ""
echo "访问地址："
echo "  - 前端：http://localhost"
echo "  - API：http://localhost:8888"
echo ""
echo "查看日志："
echo "  docker compose logs -f [服务名]"
echo ""
echo "停止服务："
echo "  docker compose down"
