# Go-Zero 学习项目

> **项目名称**：go-zero-learning
>
> **学习目标**：从 0 到 1，在 AI 指导下一步一步模仿 go-zero-looklook 项目，掌握 go-zero 框架开发
>
> **技术栈**：go-zero v1.9.3 + GORM + MySQL + Redis + JWT + gRPC
>
> **参考项目**：[go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook)

---

## 🚀 新会话使用说明

**继续项目时**：

1. 查看下方的"当前项目进度"，了解已完成的功能
2. 告诉 AI 当前状态和下一步要做什么
3. AI 会继续指导你实现下一个功能

---

## ⚠️ 重要：代码输出策略（必须遵守）

> **核心原则**：分步学习，先暴露问题，再优化方案

### 📋 策略规则

1. **分小步输出**
   - ❌ 不要一次性给出完整功能的代码
   - ✅ 每次只输出一小步，逐步实现

2. **先给简单实现**
   - ✅ 先给出能运行但有问题的版本（如静态检查警告）
   - ✅ 让学习者先遇到问题，自己思考原因

3. **暴露问题**
   - ✅ 让学习者看到警告/错误
   - ✅ 让学习者意识到这种写法的问题

4. **解释问题**
   - ✅ 说明为什么会有警告/错误
   - ✅ 解释这种写法的风险和问题

5. **给出优化方案**
   - ✅ 提供更好的实现方式
   - ✅ 解释为什么优化方案更好

### 🎯 学习目标

通过这种方式：
- ✅ 加深对问题的理解
- ✅ 学会识别和解决代码问题
- ✅ 理解最佳实践的原因
- ✅ 培养代码质量意识

### 📝 示例流程

```
步骤1：给出简单实现（有警告）
  ↓
步骤2：学习者遇到警告，思考原因
  ↓
步骤3：解释问题（为什么不好）
  ↓
步骤4：给出优化方案（为什么更好）
```

**慢一点没关系，重点是学到更多知识点！**

---

## 🎯 学习方式：边做边学

**为什么直接开始项目？**

1. ✅ 已有 gin/gorm 基础，go-zero 差异不大
2. ✅ 有实际项目驱动，学习更有针对性
3. ✅ 遇到问题再查文档，记忆更深刻

**前置准备**：

1. 安装工具：`go install github.com/zeromicro/go-zero/tools/goctl@latest`
2. 了解核心概念（见下方"最小必要知识"）
3. 然后直接开始项目！

---

## 📚 最小必要知识（30 分钟快速了解）

### 1. go-zero 是什么？

- 微服务框架（类似 Spring Cloud）
- 支持 REST API 和 gRPC RPC
- 使用 `goctl` 工具生成代码

### 2. 项目结构（和 gin 的差异）

```
gin 项目：
├── main.go
├── handler/
├── model/
└── router.go

go-zero 项目：
├── service/
│   └── user/
│       └── api/
│           ├── user.api          # API 定义文件（新）
│           ├── user.go           # 入口
│           └── internal/
│               ├── config/       # 配置
│               ├── handler/      # HTTP 处理（自动生成）
│               ├── logic/        # 业务逻辑（你写这里）
│               └── svc/           # ServiceContext（依赖注入）
```

### 3. 开发流程（和 gin 的差异）

```
gin：
1. 写路由：router.POST("/user", handler)
2. 写处理函数：func handler(c *gin.Context) { ... }

go-zero：
1. 写 .api 文件：定义 API
2. 运行 goctl api go：生成 handler、logic 骨架
3. 在 logic 中写业务逻辑
```

### 4. ServiceContext 模式（依赖注入）

```go
// 在 svc/servicecontext.go 中初始化所有依赖
type ServiceContext struct {
    Config config.Config
    DB     *gorm.DB
    Redis  *redis.Redis
}

// 在 logic 中使用
func (l *LoginLogic) Login(req *types.LoginReq) {
    // 通过 l.svcCtx.DB 访问数据库
    // 通过 l.svcCtx.Redis 访问 Redis
}
```

### 5. .api 文件语法（类似 OpenAPI）

```go
syntax = "v1"

type LoginReq {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResp {
    Token string `json:"token"`
}

service user-api {
    @handler Login
    post /api/user/login (LoginReq) returns (LoginResp)
}
```

---

## 📁 项目结构

```
go-zero-learning/
├── backend/                    # 后端项目（go.mod 在 backend/）
│   ├── common/                # 公共模块
│   │   ├── db/               # 数据库连接
│   │   ├── jwt/              # JWT 工具
│   │   └── middleware/       # 中间件
│   ├── model/                # 数据模型
│   └── service/               # 服务
│       └── user/
│           └── api/          # 用户 API 服务
├── frontend/                  # 前端项目（后续）
└── learn/                     # 学习代码（独立项目）
```

**项目特点**：
- 单模块结构（参考 go-zero-looklook）
- `backend/` 作为后端根目录
- `go.mod` 在 `backend/` 目录下

---

## 🌐 RESTful API 设计规范

### 设计原则

本项目严格遵循 RESTful API 设计规范：

1. **使用复数资源名**
   - ✅ `/api/users`（正确）
   - ❌ `/api/user`（错误）

2. **使用 HTTP 方法表示操作**
   - `POST` - 创建资源
   - `GET` - 获取资源
   - `PUT` - 更新资源（完整更新）
   - `DELETE` - 删除资源

3. **使用路径参数表示资源 ID**
   - `/api/users/:id` - 操作指定用户

4. **使用特殊资源表示当前用户**
   - `/api/users/me` - 表示当前认证用户

### 当前 API 路径规范

| 功能 | HTTP 方法 | 路径 | 说明 |
|------|----------|------|------|
| 用户注册 | POST | `/api/users` | 创建新用户 |
| 用户登录 | POST | `/api/users/login` | 登录（子资源操作） |
| 获取当前用户信息 | GET | `/api/users/me` | 获取当前认证用户信息 |
| 获取用户列表 | GET | `/api/users` | 获取用户列表（支持分页和搜索） |
| 更新当前用户信息 | PUT | `/api/users/me` | 更新当前认证用户信息 |

### 路径设计说明

- **资源使用复数**：`/api/users` 而不是 `/api/user`
- **特殊资源**：`/me` 表示当前认证用户，避免暴露用户 ID
- **子资源操作**：`/api/users/login` 表示登录操作（特殊操作使用子资源）
- **路径冲突处理**：`GET /api/users`（列表）和 `POST /api/users`（创建）使用相同路径，通过 HTTP 方法区分

### 后续扩展路径（规划）

| 功能 | HTTP 方法 | 路径 | 说明 |
|------|----------|------|------|
| 获取指定用户详情 | GET | `/api/users/:id` | 获取指定用户信息 |
| 更新指定用户信息 | PUT | `/api/users/:id` | 更新指定用户（需要权限） |
| 删除用户 | DELETE | `/api/users/:id` | 删除用户（需要权限） |

---

## 📝 当前项目进度（实时更新）

### 项目结构
- ✅ 项目根目录：`go-zero-learning/`
- ✅ 后端根目录：`backend/`（go.mod 在 backend/）
- ✅ 模块结构：单模块结构（参考 go-zero-looklook）

### 已完成功能
- ✅ 项目基础搭建
  - [x] 目录结构创建（backend/common、backend/model、backend/service）
  - [x] go.mod 配置（模块名：go-zero-learning）
  - [x] 数据库连接模块（backend/common/db）
  - [x] JWT 工具模块（backend/common/jwt）
  - [x] 用户模型（backend/model/user）
  
- ✅ 用户 API 服务框架
  - [x] API 定义（user.api）
  - [x] 代码生成（handler、logic、svc）
  - [x] ServiceContext 配置（数据库连接、自动迁移）
  - [x] 服务能正常运行（端口 8888）

- ✅ 用户认证功能
  - [x] 用户注册逻辑（密码加密 bcrypt）
  - [x] 用户登录逻辑（JWT Token 生成）
  - [x] 参数验证（go-zero 自动验证）
  - [x] 错误处理（用户名/邮箱重复检查）
  
- ✅ 认证中间件和用户信息
  - [x] 认证中间件（JWT 验证）
  - [x] Context 数据管理（ctxdata 包，避免 key 冲突）
  - [x] 获取用户信息逻辑（从 context 获取用户 ID）

- ✅ 用户管理功能
  - [x] 用户列表 API（分页、搜索）
  - [x] 用户更新 API（更新邮箱和密码）
  - [x] RESTful API 重构（统一使用 RESTful 规范）

### 待完成功能

#### 阶段一：用户认证和管理
- [x] 用户注册逻辑（密码加密 bcrypt）✅
- [x] 用户登录逻辑（JWT Token 生成）✅
- [x] 获取用户信息逻辑（需要认证中间件）✅
- [x] 认证中间件（JWT 验证）✅
- [x] 用户列表 API（分页、搜索）✅
- [x] 用户更新 API ✅
- [x] RESTful API 重构 ✅
- [ ] 用户详情 API（根据 ID 获取）
- [ ] 用户删除 API
- [ ] 错误处理优化

#### 阶段二：权限管理
- [ ] 角色管理（角色 CRUD）
- [ ] 权限管理（权限 CRUD）
- [ ] 菜单管理（菜单 CRUD）
- [ ] 权限中间件（权限验证）

#### 阶段三：商品管理
- [ ] 商品 CRUD
- [ ] 商品分类管理
- [ ] 商品库存管理

#### 阶段四：订单管理
- [ ] 订单创建
- [ ] 订单支付
- [ ] 订单状态管理
- [ ] 购物车功能

#### 阶段五：高级功能
- [ ] 文件上传下载
- [ ] 数据统计 Dashboard
- [ ] 定时任务
- [ ] Redis 缓存集成

#### 阶段六：RPC 服务
- [ ] 用户 RPC 服务
- [ ] 商品 RPC 服务
- [ ] 订单 RPC 服务
- [ ] API 服务调用 RPC 服务

#### 阶段七：优化和部署
- [ ] 错误处理完善
- [ ] 日志系统
- [ ] Docker 部署
- [ ] 性能优化

### 当前问题/注意事项
- 配置文件字段名：使用 `dataSource`（小写驼峰）
- 运行方式：`cd backend/service/user/api && go run user-apic.go`
- 数据库：MySQL 3307 端口，数据库名 testdb

### 下一步计划
1. ✅ 实现用户注册和登录逻辑（已完成）
2. ✅ 实现获取用户信息逻辑（已完成）
3. ✅ 添加认证中间件（JWT 验证）（已完成）
4. ✅ 用户列表 API（分页、搜索）（已完成）
5. ✅ 用户更新 API（已完成）
6. ✅ RESTful API 重构（已完成）
7. 用户详情 API（根据 ID 获取）
8. 用户删除 API
9. 完善错误处理

**最后更新**：2025-01-22  
**当前状态**：用户认证和管理功能基本完成（注册、登录、获取用户信息、用户列表、用户更新），已重构为 RESTful 风格，待实现用户详情和删除功能

---

## 🧪 测试用例

> **注意**：所有 API 路径已更新为 RESTful 风格，请使用新的路径进行测试。

### 用户注册接口 (`POST /api/users`)

#### 成功场景
- [x] **注册新用户成功**
  ```bash
  curl -X POST http://localhost:8888/api/users \
    -H "Content-Type: application/json" \
    -d '{"username":"testuser","email":"test@example.com","password":"123456"}'
  ```
  **预期响应**：返回 Token 和用户信息

#### 失败场景
- [x] **用户名已存在**
  ```bash
  curl -X POST http://localhost:8888/api/users \
    -H "Content-Type: application/json" \
    -d '{"username":"testuser","email":"another@example.com","password":"123456"}'
  ```
  **预期响应**：`{"message": "用户名已存在"}`

- [x] **邮箱已存在**
  ```bash
  curl -X POST http://localhost:8888/api/users \
    -H "Content-Type: application/json" \
    -d '{"username":"newuser","email":"test@example.com","password":"123456"}'
  ```
  **预期响应**：`{"message": "邮箱已存在"}`

- [x] **参数缺失（username）**
  ```bash
  curl -X POST http://localhost:8888/api/users \
    -H "Content-Type: application/json" \
    -d '{"email":"test@example.com","password":"123456"}'
  ```
  **预期响应**：`field "username" is not set`

- [x] **参数缺失（email）**
  ```bash
  curl -X POST http://localhost:8888/api/users \
    -H "Content-Type: application/json" \
    -d '{"username":"testuser","password":"123456"}'
  ```
  **预期响应**：`field "email" is not set`

- [x] **参数缺失（password）**
  ```bash
  curl -X POST http://localhost:8888/api/users \
    -H "Content-Type: application/json" \
    -d '{"username":"testuser","email":"test@example.com"}'
  ```
  **预期响应**：`field "password" is not set`

---

### 用户登录接口 (`POST /api/users/login`)

#### 成功场景
- [x] **登录成功**
  ```bash
  curl -X POST http://localhost:8888/api/users/login \
    -H "Content-Type: application/json" \
    -d '{"username":"testuser","password":"123456"}'
  ```
  **预期响应**：返回 Token 和用户信息

#### 失败场景
- [x] **用户名不存在**
  ```bash
  curl -X POST http://localhost:8888/api/users/login \
    -H "Content-Type: application/json" \
    -d '{"username":"nonexistent","password":"123456"}'
  ```
  **预期响应**：`{"message": "用户名或密码错误"}`

- [x] **密码错误**
  ```bash
  curl -X POST http://localhost:8888/api/users/login \
    -H "Content-Type: application/json" \
    -d '{"username":"testuser","password":"wrongpassword"}'
  ```
  **预期响应**：`{"message": "用户名或密码错误"}`

- [x] **参数缺失（username）**
  ```bash
  curl -X POST http://localhost:8888/api/users/login \
    -H "Content-Type: application/json" \
    -d '{"password":"123456"}'
  ```
  **预期响应**：`field "username" is not set`

- [x] **参数缺失（password）**
  ```bash
  curl -X POST http://localhost:8888/api/users/login \
    -H "Content-Type: application/json" \
    -d '{"username":"testuser"}'
  ```
  **预期响应**：`field "password" is not set`

---

### 获取用户信息接口 (`GET /api/users/me`)

**需要认证**：需要在请求头中提供 `Authorization: Bearer <token>`

#### 成功场景
- [x] **获取用户信息成功**
  ```bash
  # 1. 先登录获取 Token
  curl -X POST http://localhost:8888/api/users/login \
    -H "Content-Type: application/json" \
    -d '{"username":"testuser","password":"123456"}'
  
  # 2. 使用返回的 Token 获取用户信息（将 YOUR_TOKEN 替换为实际 token）
  curl -X GET http://localhost:8888/api/users/me \
    -H "Authorization: Bearer YOUR_TOKEN"
  ```
  **预期响应**：返回用户信息（id、username、email）

#### 失败场景
- [x] **未提供 Token**
  ```bash
  curl -X GET http://localhost:8888/api/users/me
  ```
  **预期响应**：`未提供认证 token` 或类似错误

- [x] **Token 格式错误**
  ```bash
  curl -X GET http://localhost:8888/api/users/me \
    -H "Authorization: invalid-format"
  ```
  **预期响应**：`token 格式错误`

- [x] **Token 无效或已过期**
  ```bash
  curl -X GET http://localhost:8888/api/users/me \
    -H "Authorization: Bearer invalid-token-12345"
  ```
  **预期响应**：`token 无效或已过期`

---

### 完整测试脚本

```bash
#!/bin/bash

echo "=== 1. 注册新用户 ==="
curl -X POST http://localhost:8888/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser1","email":"test1@example.com","password":"123456"}'
echo -e "\n\n"

echo "=== 2. 重复注册（用户名已存在） ==="
curl -X POST http://localhost:8888/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser1","email":"test2@example.com","password":"123456"}'
echo -e "\n\n"

echo "=== 3. 重复注册（邮箱已存在） ==="
curl -X POST http://localhost:8888/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser2","email":"test1@example.com","password":"123456"}'
echo -e "\n\n"

echo "=== 4. 登录成功 ==="
curl -X POST http://localhost:8888/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser1","password":"123456"}'
echo -e "\n\n"

echo "=== 5. 登录失败（密码错误） ==="
curl -X POST http://localhost:8888/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser1","password":"wrong"}'
echo -e "\n\n"

echo "=== 6. 登录失败（用户不存在） ==="
curl -X POST http://localhost:8888/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"nonexistent","password":"123456"}'
echo -e "\n"
```

---

### 用户列表接口 (`GET /api/users`)

**需要认证**：需要在请求头中提供 `Authorization: Bearer <token>`

#### 成功场景
- [x] **获取用户列表（分页）**
  ```bash
  curl -X GET "http://localhost:8888/api/users?page=1&page_size=10" \
    -H "Authorization: Bearer YOUR_TOKEN"
  ```
  **预期响应**：返回用户列表、总数、页码、每页数量

- [x] **获取用户列表（搜索）**
  ```bash
  curl -X GET "http://localhost:8888/api/users?keyword=test" \
    -H "Authorization: Bearer YOUR_TOKEN"
  ```
  **预期响应**：返回匹配关键词的用户列表

- [x] **获取用户列表（默认参数）**
  ```bash
  curl -X GET http://localhost:8888/api/users \
    -H "Authorization: Bearer YOUR_TOKEN"
  ```
  **预期响应**：返回第1页，每页10条的用户列表

---

### 更新用户信息接口 (`PUT /api/users/me`)

**需要认证**：需要在请求头中提供 `Authorization: Bearer <token>`

#### 成功场景
- [x] **更新邮箱**
  ```bash
  curl -X PUT http://localhost:8888/api/users/me \
    -H "Authorization: Bearer YOUR_TOKEN" \
    -H "Content-Type: application/json" \
    -d '{"email":"newemail@example.com"}'
  ```
  **预期响应**：返回更新后的用户信息

- [x] **更新密码**
  ```bash
  curl -X PUT http://localhost:8888/api/users/me \
    -H "Authorization: Bearer YOUR_TOKEN" \
    -H "Content-Type: application/json" \
    -d '{"password":"newpassword123"}'
  ```
  **预期响应**：返回更新后的用户信息

- [x] **同时更新邮箱和密码**
  ```bash
  curl -X PUT http://localhost:8888/api/users/me \
    -H "Authorization: Bearer YOUR_TOKEN" \
    -H "Content-Type: application/json" \
    -d '{"email":"newemail@example.com","password":"newpassword123"}'
  ```
  **预期响应**：返回更新后的用户信息

#### 失败场景
- [x] **未提供任何更新字段**
  ```bash
  curl -X PUT http://localhost:8888/api/users/me \
    -H "Authorization: Bearer YOUR_TOKEN" \
    -H "Content-Type: application/json" \
    -d '{}'
  ```
  **预期响应**：`至少需要提供一个更新字段`

- [x] **邮箱已被其他用户使用**
  ```bash
  curl -X PUT http://localhost:8888/api/users/me \
    -H "Authorization: Bearer YOUR_TOKEN" \
    -H "Content-Type: application/json" \
    -d '{"email":"existing@example.com"}'
  ```
  **预期响应**：`邮箱已被使用`

- [x] **密码长度不足**
  ```bash
  curl -X PUT http://localhost:8888/api/users/me \
    -H "Authorization: Bearer YOUR_TOKEN" \
    -H "Content-Type: application/json" \
    -d '{"password":"123"}'
  ```
  **预期响应**：`密码至少需要6位`

---

## 🔗 学习资源

1. **go-zero 官方文档**：https://go-zero.dev/
2. **GORM 文档**：https://gorm.io/
3. **go-zero-looklook**：https://github.com/Mikaelemmmm/go-zero-looklook
4. **Protocol Buffers 文档**：https://protobuf.dev/
5. **Redis 文档**：https://redis.io/docs/

---

## 💡 实践建议

1. ✅ **从空目录开始**，不要直接复制代码
2. ✅ **一次一个小功能**，像正常开发一样
3. ✅ **理解每一行代码**，不要只是复制粘贴
4. ✅ **对比参考项目**，找出差异和优化点
5. ✅ **记录问题和解决方案**，形成学习笔记

---

## 📌 新会话启动提示

**继续项目时，可以这样说**：

```
我在学习 go-zero 项目，请先查看 learn.md 了解当前进度。
当前需要实现用户注册和登录的业务逻辑。
```

**重要提醒**：
- ⚠️ **必须遵守代码输出策略**（见上方"重要：代码输出策略"章节）
- ✅ 分小步输出，先暴露问题，再优化方案
- ✅ 让学习者先遇到问题，再解释和优化

**项目信息**：
- 项目结构：backend/ 作为后端根目录，go.mod 在 backend/
- 参考项目：go-zero-looklook
- 技术栈：go-zero v1.9.3 + GORM + MySQL + Redis + JWT
