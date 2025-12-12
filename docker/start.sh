#!/bin/bash

# 开启严格 bash 模式：脚本遇错即停（-e），引用未定义变量时报错（-u），管道一出错即失败（-o pipefail）；可提高脚本健壮性与安全性
set -euo pipefail

echo "🚀 启动 Go-Zero Learning 项目..."

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
echo "docker exec -i go-zero-mysql mysql -uroot -p123456 --default-character-set=utf8mb4 testdb < scripts/init_test_data.sql"
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