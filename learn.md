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

### 待完成功能

#### 阶段一：用户认证和管理
- [ ] 用户注册逻辑（密码加密 bcrypt）
- [ ] 用户登录逻辑（JWT Token 生成）
- [ ] 获取用户信息逻辑（需要认证中间件）
- [ ] 认证中间件（JWT 验证）
- [ ] 用户列表 API（分页、搜索）
- [ ] 用户详情 API
- [ ] 用户更新 API
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
1. 实现用户注册和登录逻辑
2. 添加认证中间件
3. 完善错误处理
4. 测试所有 API

**最后更新**：2025-01-22  
**当前状态**：服务已运行，待实现业务逻辑

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

**项目信息**：
- 项目结构：backend/ 作为后端根目录，go.mod 在 backend/
- 参考项目：go-zero-looklook
- 技术栈：go-zero v1.9.3 + GORM + MySQL + Redis + JWT
