# Repository Guidelines

## Project Structure & Module Organization
后端 Go 代码集中在 `backend/`，公共能力放在 `common/`，业务服务位于 `service/`（如 `service/user/api` 提供 REST、`service/user/user-rpc` 处理 RPC）。数据模型在 `backend/model`，配置文件位于 `service/*/etc`，对应 Dockerfile（示例 `Dockerfile.user-api`）放在 `backend/` 根目录。前端 Vue 项目位于 `frontend/src`，按 `api/`、`views/system/{user,role}/`、`router/` 分类；脚本在 `scripts/`，容器和部署资产在 `docker/`，运行日志写入 `logs/`，而 `learn/` 保存学习性示例代码。

## Build, Test, and Development Commands
`cd backend && go build ./...` 校验 Go 模块是否能完整编译。`cd backend/service/user/api && go run user-api.go` 启动 HTTP 网关，`cd backend/service/user/user-rpc && go run user-rpc.go` 启动 RPC 服务。`docker-compose up -d` 会拉起 MySQL、Redis 等依赖。前端开发执行 `cd frontend && pnpm install && pnpm dev`，构建产物用 `pnpm build`，上线前在 `pnpm preview` 中二次确认。待服务监听 `:8888` 后运行 `./scripts/test-order-cancel-job.sh` 或 `./scripts/test-permission-middleware.sh` 回归关键链路。

## Coding Style & Naming Conventions
提交前必须 `gofmt`+`goimports`，Go 代码使用制表符缩进，导出符号采用驼峰命名，GORM 标签坚持 snake_case。执行 `cd backend && staticcheck ./...`（规则存于 `backend/staticcheck.conf`）。Vue/TS 文件保持两空格缩进，优先 `<script setup>`，文件名使用 kebab-case（例如 `views/system/user/UserList.vue`）。API 客户端按模块划分到 `frontend/src/api`，通过 `pnpm lint`、`pnpm format`、`pnpm type-check` 统一 ESLint、Prettier、Vue TSC 结果。

## Testing Guidelines
后端单测紧邻源码，以 `_test.go` 结尾并以 `Test{Feature}` 命名，使用 `cd backend && go test ./... -cover` 汇总覆盖率并在 CI 中上传。长链路测试依赖脚本与 `scripts/init_test_data.sql`，确保数据可重复。前端组件建议补充 Vitest 套件；在引入自动化前至少保证 `pnpm type-check` 通过，测试文件命名遵循路由语义（如 `system-role.spec.ts`）。

## Commit & Pull Request Guidelines
仓库使用 Conventional Commits（示例 `feat(docker): ...`），格式为 `type(scope): 命令式概要`，正文单行不超过 100 字符。每个 PR 需描述功能改动、列出手动验证命令、关联 issue，并在 UI 变化时附截图。合并前请同步 README/配置中的指引，并至少邀请一位审阅者。

## Security & Configuration Tips
禁止在代码中硬编码账号密钥，统一写入 `service/*/etc/*.yaml` 或 Docker Compose 的环境变量。`.env`、日志文件、`frontend/dist/` 等生成物不得提交。共享环境时用 `scripts/init_test_data.sql` 重新灌入脱敏数据，定期轮换 JWT 密钥，并避免在示例配置中暴露生产地址。
