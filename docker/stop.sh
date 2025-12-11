#!/bin/bash

echo "🛑 停止 Go-Zero Learning 项目..."

docker compose down

echo ""
echo "✅ 服务已停止"
echo ""
echo "💡 提示："
echo "   - 如需删除数据卷，运行: docker compose down -v"
echo "   - 如需查看日志，运行: docker compose logs [服务名]"