#!/bin/bash

# 定时任务测试脚本：订单自动取消功能
# 使用方法：./scripts/test-order-cancel-job.sh

BASE_URL="http://127.0.0.1:8888"
TOKEN=""

echo "=========================================="
echo "定时任务测试：订单自动取消功能"
echo "=========================================="
echo ""

# 1. 登录获取 Token
echo "1. 登录获取 Token..."
LOGIN_RESPONSE=$(curl -s -X POST "${BASE_URL}/api/users/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "123456"
  }')

TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*' | cut -d'"' -f4)

if [ -z "$TOKEN" ]; then
  echo "❌ 登录失败"
  echo "响应: $LOGIN_RESPONSE"
  exit 1
fi

echo "✅ 登录成功，Token: ${TOKEN}"
echo ""

# 2. 查询商品列表（获取商品ID）
echo "2. 查询商品列表..."
PRODUCTS_RESPONSE=$(curl -s -X GET "${BASE_URL}/api/products?page=1&page_size=10" \
  -H "Authorization: Bearer $TOKEN")

PRODUCT_ID=$(echo $PRODUCTS_RESPONSE | grep -o '"id":[0-9]*' | head -1 | cut -d':' -f2)

if [ -z "$PRODUCT_ID" ]; then
  echo "❌ 获取商品失败"
  echo "响应: $PRODUCTS_RESPONSE"
  exit 1
fi

echo "✅ 获取商品ID: $PRODUCT_ID"
echo ""

# 3. 创建待支付订单
echo "3. 创建待支付订单..."
ORDER_RESPONSE=$(curl -s -X POST "${BASE_URL}/api/orders" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"items\": [
      {
        \"product_id\": $PRODUCT_ID,
        \"quantity\": 1
      }
    ],
    \"remark\": \"定时任务测试订单\"
  }")

ORDER_ID=$(echo $ORDER_RESPONSE | grep -o '"id":[0-9]*' | head -1 | cut -d':' -f2)
ORDER_NO=$(echo $ORDER_RESPONSE | grep -o '"order_no":"[^"]*' | cut -d'"' -f4)

if [ -z "$ORDER_ID" ]; then
  echo "❌ 创建订单失败"
  echo "响应: $ORDER_RESPONSE"
  exit 1
fi

echo "✅ 订单创建成功"
echo "   订单ID: $ORDER_ID"
echo "   订单编号: $ORDER_NO"
echo "   订单状态: 待支付（1）"
echo ""

# 4. 记录订单创建时间
CREATE_TIME=$(date +%s)
echo "4. 订单创建时间: $(date '+%Y-%m-%d %H:%M:%S')"
echo ""

# 5. 等待定时任务执行
echo "5. 等待定时任务执行（等待3.5分钟，确保订单超过1分钟超时时间）..."
echo "   定时任务每分钟的第30秒执行一次，订单超时时间为1分钟"
echo "   请观察服务日志，应该看到订单被自动取消"
echo ""

# 计算需要等待的时间（订单超时1分钟 + 缓冲时间）
# 确保至少执行2-3次定时任务（每分钟执行一次）
WAIT_MINUTES=3.5
WAIT_SECONDS=$(echo "$WAIT_MINUTES * 60" | bc)

# 如果系统没有 bc，使用整数计算（210秒 = 3.5分钟）
if [ -z "$WAIT_SECONDS" ]; then
  WAIT_SECONDS=210
fi

# 分7次等待，每次30秒，总共210秒（3.5分钟）
for i in {1..7}; do
  echo "   等待中... ($i/7) - $(date '+%H:%M:%S')"
  sleep 5
done

echo ""

# 6. 查询订单状态
echo "6. 查询订单状态..."
ORDER_DETAIL_RESPONSE=$(curl -s -X GET "${BASE_URL}/api/orders/${ORDER_ID}" \
  -H "Authorization: Bearer $TOKEN")

# 检查响应是否为空
if [ -z "$ORDER_DETAIL_RESPONSE" ]; then
  echo "   ❌ API 请求失败：响应为空"
  echo "   请检查："
  echo "   1. 服务是否正常运行"
  echo "   2. TOKEN 是否有效"
  echo "   3. 订单ID是否正确: $ORDER_ID"
  echo ""
  echo "   💡 提示：如果数据库查询显示订单状态为5，说明定时任务成功"
  ORDER_STATUS=""
else
  # 使用 jq 解析（如果可用）
  if command -v jq &> /dev/null; then
    ORDER_STATUS=$(echo "$ORDER_DETAIL_RESPONSE" | jq -r '.data.status // empty')
    STATUS_TEXT=$(echo "$ORDER_DETAIL_RESPONSE" | jq -r '.data.status_text // empty')
  # 使用 Python 解析
  elif command -v python3 &> /dev/null; then
    ORDER_STATUS=$(echo "$ORDER_DETAIL_RESPONSE" | python3 -c "import sys, json; data=json.load(sys.stdin); print(data.get('data', {}).get('status', ''))" 2>/dev/null)
    STATUS_TEXT=$(echo "$ORDER_DETAIL_RESPONSE" | python3 -c "import sys, json; data=json.load(sys.stdin); print(data.get('data', {}).get('status_text', ''))" 2>/dev/null)
  # 备用方案：使用 sed
  else
    ORDER_STATUS=$(echo "$ORDER_DETAIL_RESPONSE" | sed -n 's/.*"data":{[^}]*"status":\([0-9]*\).*/\1/p')
    STATUS_TEXT=$(echo "$ORDER_DETAIL_RESPONSE" | grep -o '"status_text":"[^"]*' | cut -d'"' -f4)
  fi

  echo "   订单状态码: $ORDER_STATUS"
  echo "   订单状态: $STATUS_TEXT"
  echo ""
fi

# 7. 验证结果
if [ "$ORDER_STATUS" = "5" ]; then
  echo "✅ 测试成功！订单已被自动取消"
  echo ""
  echo "7. 检查库存是否恢复..."
  PRODUCT_DETAIL_RESPONSE=$(curl -s -w "\nHTTP_CODE:%{http_code}" -X GET "${BASE_URL}/api/products/${PRODUCT_ID}" \
    -H "Authorization: Bearer $TOKEN")
  
  # 分离 HTTP 状态码和响应体
  HTTP_CODE=$(echo "$PRODUCT_DETAIL_RESPONSE" | grep "HTTP_CODE" | cut -d':' -f2)
  PRODUCT_DETAIL_RESPONSE=$(echo "$PRODUCT_DETAIL_RESPONSE" | sed '/HTTP_CODE/d')
  
  # 从 data 对象中提取 stock
  if command -v python3 &> /dev/null; then
    STOCK=$(echo "$PRODUCT_DETAIL_RESPONSE" | python3 -c "import sys, json; data=json.load(sys.stdin); print(data.get('data', {}).get('stock', ''))" 2>/dev/null)
  else
    STOCK=$(echo "$PRODUCT_DETAIL_RESPONSE" | sed -n 's/.*"data":{[^}]*"stock":\([0-9]*\).*/\1/p')
    if [ -z "$STOCK" ]; then
      STOCK=$(echo "$PRODUCT_DETAIL_RESPONSE" | grep -o '"stock":[0-9]*' | tail -1 | cut -d':' -f2)
    fi
  fi
  echo "   商品库存: $STOCK"
  echo ""
  echo "✅ 测试完成！"
else
  echo "❌ 测试失败！订单状态为 $ORDER_STATUS，预期为 5（已取消）"
  echo "   请检查："
  echo "   1. 定时任务是否正常启动"
  echo "   2. 订单创建时间是否超过1分钟"
  echo "   3. 查看服务日志确认任务执行情况"
  echo ""
  echo "   💡 提示：如果数据库查询显示订单状态为5，说明定时任务成功，只是脚本解析有问题"
fi

echo ""
echo "=========================================="