# 构建阶段
FROM node:20-alpine AS builder

WORKDIR /build

# 复制 package.json 和 pnpm-lock.yaml
COPY package.json pnpm-lock.yaml ./

# 安装 pnpm
RUN npm install -g pnpm@9

# 安装依赖
RUN pnpm install

# 复制源代码
COPY . .

# 构建生产版本
RUN pnpm build

# ==================== 运行阶段（最小镜像） ====================
FROM nginx:alpine

# 复制构建产物到 Nginx
COPY --from=builder /build/dist /usr/share/nginx/html

# 复制 Nginx 配置
# 因为构建上下文是 ./frontend，无法访问 ../docker/nginx.conf
# 由于 docker compose.yml 已通过 volume 挂载配置文件，这里可删除。
# COPY docker/nginx.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 80

# 启动 Nginx(以守护进程方式运行)
CMD ["nginx", "-g", "daemon off;"]
