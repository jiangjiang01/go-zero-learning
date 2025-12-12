#!/bin/bash

# 开启严格 bash 模式：脚本遇错即停（-e），引用未定义变量时报错（-u），管道一出错即失败（-o pipefail）；可提高脚本健壮性与安全性
set -euo pipefail

echo "🛑 停止 Go-Zero Learning 项目..."

docker compose down

echo ""
echo "✅ 服务已停止"
echo ""
echo "💡 提示："
echo "   - 如需删除数据卷，运行: docker compose down -v"
echo "   - 如需查看日志，运行: docker compose logs [服务名]"