# Go 后端开发学习计划

> **项目名称**：go-zero-admin（或 go-zero-learning）
>
> **学习目标**：从 0 到 1，在 AI 指导下一步一步模仿 go-zero-looklook 项目，掌握 go-zero 框架开发
>
> **技术栈要求**：go-zero v1.5.2 + GORM + MySQL/PostgreSQL + Redis + JWT + gRPC + 定时任务
>
> **参考项目**：
>
> - 主要参考：go-zero-looklook (https://github.com/Mikaelemmmm/go-zero-looklook)
> - 公司项目：用于对比学习和扩展场景

---

## 🚀 新会话使用说明

**如果你是第一次使用这个文档**：

1. 你已经创建了空目录，并将此 `learn.md` 文件放入其中
2. 直接复制下面的"开始提示"给 AI，开始学习
3. AI 会参考此文档和参考项目，指导你从零开始实现

**如果你要继续之前的项目**：

1. 查看下方的"学习进度跟踪"，了解当前进度
2. 告诉 AI 你完成了哪个阶段，下一步要做什么
3. AI 会继续指导你实现下一个功能

---

## 📋 目录

- [学习方式建议](#学习方式建议)
- [短期学习计划（2-3 周）](#短期学习计划2-3周压缩版)
- [中期学习计划（1-2 个月）](#中期学习计划1-2个月边学边看公司项目)
- [项目一：用户管理系统](#项目一用户管理系统)
- [项目二：微服务电商系统](#项目二微服务电商系统)
- [学习资源](#学习资源)

---

## 🎯 学习方式建议

### 推荐：边做边学

**为什么直接开始项目？**

1. ✅ 你已经有 gin/gorm 基础，go-zero 差异不大
2. ✅ 有实际项目驱动，学习更有针对性
3. ✅ 遇到问题再查文档，记忆更深刻
4. ✅ 实践中学比纯理论学习快 3-5 倍

**前置准备（30 分钟）**：

1. 安装工具：`go install github.com/zeromicro/go-zero/tools/goctl@latest`
2. 了解核心概念（见下方"最小必要知识"）
3. 然后直接开始项目！

### 最小必要知识（30 分钟快速了解）

#### 1. go-zero 是什么？

- 微服务框架（类似 Spring Cloud）
- 支持 REST API 和 gRPC RPC
- 使用 `goctl` 工具生成代码

#### 2. 项目结构（和 gin 的差异）

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

#### 3. 开发流程（和 gin 的差异）

```
gin：
1. 写路由：router.POST("/user", handler)
2. 写处理函数：func handler(c *gin.Context) { ... }

go-zero：
1. 写 .api 文件：定义 API
2. 运行 goctl api go：生成 handler、logic 骨架
3. 在 logic 中写业务逻辑
```

#### 4. ServiceContext 模式（依赖注入）

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

#### 5. .api 文件语法（类似 OpenAPI）

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

**学习建议**：

- ✅ 先了解这 5 个核心概念（30 分钟）
- ✅ 然后直接开始项目
- ✅ 遇到问题再查文档：https://go-zero.dev/
- ✅ 对比公司项目的代码，加深理解

---

## ⏱️ 学习时间评估（基于你的学习速度）

> **参考**：你一周完成了 gin-vue-admin（登录、用户管理、角色管理、菜单管理、日志、字典、Dashboard）
>
> **你的优势**：10 年前端经验 + 已掌握 gin/gorm 基础 + 编程概念理解快

### 🎯 实际学习时间预估

| 项目阶段                   | 原计划 | **实际预估**     | 说明                                   |
| -------------------------- | ------ | ---------------- | -------------------------------------- |
| **项目一：用户管理系统**   | 8 周   | **2-3 周**       | 你已会 gin/gorm，主要是学 go-zero 差异 |
| **项目二：微服务电商系统** | 4 个月 | **1-2 个月**     | 边学边看公司项目，加速理解             |
| **总计**                   | 6 个月 | **1.5-2.5 个月** | 专注学习的话可以更快                   |

### 📍 关键节点：什么时候可以开始看公司项目？

**✅ 可以开始看公司项目的节点：完成项目一的阶段一、二、三后（约 2-3 周）**

**原因**：

- ✅ 已掌握 go-zero REST API 开发
- ✅ 已掌握 GORM + MySQL 在 go-zero 中的使用
- ✅ 已掌握 Redis + JWT 认证
- ✅ 已掌握中间件开发
- ✅ **此时可以看懂公司项目的 `service/admin/api` 部分（API 网关服务）**

**学习策略**：

1. **第 1-2 周**：专注学习，完成项目一的基础框架和用户功能
2. **第 3 周开始**：可以开始看公司项目的 API 服务部分，同时继续学习 RPC
3. **第 4 周开始**：学习 RPC 的同时，看公司项目的 RPC 服务，对比学习

---

## 短期学习计划（2-3 周，压缩版）

### 🚀 学习方式：边做边学（推荐）

**建议**：直接开始项目，遇到问题再查文档。这样学习更高效！

**前置准备（30 分钟）**：

1. 安装 go-zero 和 goctl：`go install github.com/zeromicro/go-zero/tools/goctl@latest`
2. 快速了解核心概念（看下面"最小必要知识"）
3. 然后直接开始项目，边做边学

---

### 第 1 周：go-zero 基础 + 用户功能

**目标**：掌握 go-zero 核心概念，实现基础 CRUD

**学习方式**：直接开始项目，遇到问题再查文档

| 学习内容         | 最小必要知识（30 分钟）                                                                                                                                                    | 实践项目                         |
| ---------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------- |
| go-zero 框架基础 | 1. go-zero 是什么：微服务框架<br>2. REST API vs RPC：REST 对外，RPC 对内<br>3. goctl 工具：代码生成器                                                                      | 创建第一个 go-zero REST API 服务 |
| 项目结构理解     | 1. 标准结构：`service/xxx/api/internal/{config,handler,logic,svc}`<br>2. config：配置<br>3. svc：ServiceContext（依赖注入）<br>4. logic：业务逻辑<br>5. handler：HTTP 处理 | 搭建项目骨架                     |
| API 定义         | 1. `.api` 文件：定义 API（类似 OpenAPI）<br>2. `goctl api go`：生成代码<br>3. handler → logic：handler 调用 logic                                                          | 定义用户注册/登录 API            |
| GORM 集成        | 1. 在 ServiceContext 中初始化 DB<br>2. 在 logic 中使用 DB<br>3. 和 gin 中使用 GORM 一样                                                                                    | 实现用户注册、登录、列表功能     |

**关键差异（gin vs go-zero）**：

- gin：直接写 `router.POST("/user", handler)`
- go-zero：写 `.api` 文件 → `goctl` 生成 → 在 `logic` 中写业务

### 第 2 周：认证 + Redis + 中间件

**目标**：掌握认证和缓存

**最小必要知识（15 分钟）**：

#### Redis 基础（5 分钟）

```go
// 1. 连接 Redis（在 ServiceContext 中）
redis := redis.NewRedisInstance(host, username, password, db)

// 2. 基本操作
redis.Set("key", "value", 3600)  // 设置，过期时间 3600 秒
redis.Get("key")                  // 获取
redis.DelKey("key")               // 删除
```

#### JWT 认证（5 分钟）

```go
// 1. 生成 Token
token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

// 2. 验证 Token
token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte(secret), nil
})

// 3. 在中间件中验证（类似 gin 的中间件）
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 验证 token
        // 通过后调用 next(w, r)
    }
}
```

#### go-zero 中间件（5 分钟）

```go
// 在 main.go 中注册中间件
server.Use(authMiddleware.Handle)

// 中间件格式（和 gin 类似）
func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 验证逻辑
        next(w, r)
    }
}
```

| 学习内容   | 具体任务                                                             | 实践项目                   |
| ---------- | -------------------------------------------------------------------- | -------------------------- |
| Redis 基础 | 1. Redis 连接池配置<br>2. 基本操作（Get/Set/Del）<br>3. 过期时间设置 | 实现登录会话管理           |
| JWT 认证   | 1. JWT 生成和验证<br>2. Token 刷新机制<br>3. 中间件实现              | 实现完整的认证系统         |
| 中间件开发 | 1. 认证中间件<br>2. 错误处理中间件<br>3. 请求日志中间件              | 添加认证和日志中间件       |
| 查询优化   | 1. 分页查询<br>2. 条件过滤<br>3. 完善用户管理功能                    | 完善用户列表（分页、搜索） |

### 第 3 周：gRPC RPC 服务

**目标**：掌握微服务间通信

**最小必要知识（20 分钟）**：

#### Protocol Buffers（10 分钟）

```protobuf
// user.proto
syntax = "proto3";

service User {
    rpc GetUser(GetUserReq) returns (GetUserResp);
}

message GetUserReq {
    string id = 1;
}

message GetUserResp {
    string id = 1;
    string username = 2;
}
```

**关键点**：

- `.proto` 文件：定义 RPC 服务和消息
- `goctl rpc protoc`：生成 Go 代码
- 类似 `.api` 文件，但是用于 RPC

#### gRPC 服务开发（5 分钟）

```go
// 1. 创建 RPC 服务（类似 REST API）
// goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.

// 2. 实现服务方法（在 logic 中）
func (l *GetUserLogic) GetUser(in *pb.GetUserReq) (*pb.GetUserResp, error) {
    // 业务逻辑
    return &pb.GetUserResp{...}, nil
}
```

#### 服务调用（5 分钟）

```go
// 在 API 服务的 ServiceContext 中配置 RPC 客户端
type ServiceContext struct {
    UserRpc user.UserClient  // RPC 客户端
}

// 在 logic 中调用 RPC
func (l *SomeLogic) SomeMethod() {
    resp, err := l.svcCtx.UserRpc.GetUser(ctx, &pb.GetUserReq{Id: "123"})
}
```

**关键差异**：

- REST API：对外提供 HTTP 接口
- RPC 服务：对内提供 gRPC 接口，服务间调用

| 学习内容         | 具体任务                                                         | 实践项目              |
| ---------------- | ---------------------------------------------------------------- | --------------------- |
| Protocol Buffers | 1. .proto 文件编写<br>2. 使用 protoc 生成代码<br>3. 理解消息定义 | 定义用户服务 proto    |
| gRPC 服务开发    | 1. 创建 RPC 服务<br>2. 实现服务方法<br>3. 服务注册               | 实现用户 RPC 服务     |
| 服务调用         | 1. RPC 客户端配置<br>2. 跨服务调用<br>3. 错误处理                | API 服务调用 RPC 服务 |

---

## 中期学习计划（1-2 个月，边学边看公司项目）

### 第 4-5 周：定时任务 + 高级特性（边学边看公司项目）

**学习策略**：此时可以开始看公司项目，对比学习

**最小必要知识（15 分钟）**：

#### 定时任务（5 分钟）

```go
// 1. 创建 cron 实例
c := cron.New()

// 2. 添加定时任务
c.AddFunc("0 0 * * *", func() {
    // 每天凌晨执行
})

// 3. 启动
c.Start()

// 定时表达式：秒 分 时 日 月 周
// "0 0 * * *" = 每天 00:00:00
// "0 */1 * * *" = 每小时
```

#### 文件处理（5 分钟）

```go
// Excel 导出（excelize）
f := excelize.NewFile()
f.SetCellValue("Sheet1", "A1", "用户名")
f.SaveAs("users.xlsx")

// 文件上传（go-zero）
// 在 .api 文件中定义文件上传接口
```

#### 错误处理（5 分钟）

```go
// 统一错误返回
type CodeError struct {
    Code int
    Msg  string
}

// 在中间件中捕获错误
func ErrorMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                // 统一错误处理
            }
        }()
        next(w, r)
    }
}
```

| 学习内容 | 具体任务                                                        | 实践项目             | 公司项目对照                   |
| -------- | --------------------------------------------------------------- | -------------------- | ------------------------------ |
| 定时任务 | 1. robfig/cron 使用<br>2. 定时任务服务设计<br>3. 任务调度管理   | 实现数据统计定时任务 | 看 `service/*/schedule/`       |
| 文件处理 | 1. 文件上传下载<br>2. Excel 导入导出（excelize）<br>3. 图片处理 | 实现文件管理功能     | 看 `common/tool/excel_util.go` |
| 错误处理 | 1. 统一错误处理<br>2. 错误码设计<br>3. 错误拦截器               | 完善错误处理机制     | 看 `common/error/`             |

### 第 6-8 周：完整项目实战（参考公司项目）

**目标**：完成一个完整的微服务项目，同时深入理解公司项目

**实践项目**：电商后台管理系统（参考公司项目结构）

**功能模块**：

1. 用户管理（用户、角色、权限）
2. 商品管理（商品、分类、库存）
3. 订单管理（订单、支付、物流）
4. 数据统计（Dashboard、报表）
5. 系统管理（字典、日志、配置）

**技术要点**：

- 微服务拆分（API + 多个 RPC 服务）
- 完整的认证授权
- Redis 缓存
- 定时任务
- 文件管理
- 完整的错误处理

**学习方式**：

- 实现一个功能 → 对比公司项目相同功能 → 理解差异和优化点

---

## 完整项目：模仿 go-zero-looklook 项目

> **项目名称**：go-zero-admin（或 go-zero-learning）
>
> **学习周期**：第 1-8 周（分阶段实现，逐步完善）
>
> **关键节点**：完成阶段一、二、三后（约 2-3 周），可以开始看公司项目的 API 服务部分
>
> **技术栈**：go-zero REST API + GORM + MySQL + Redis + JWT + gRPC
>
> **参考项目**：
>
> - **主要参考**：go-zero-looklook (https://github.com/Mikaelemmmm/go-zero-looklook)
> - **公司项目参考**：当前工作区的项目（用于对比学习和扩展场景）
>
> **学习方式**：从 0 到 1，在 AI 指导下一步一步模仿 go-zero-looklook 的功能实现

### 📋 项目概述

这是一个**完整的项目**，分阶段实现 go-zero-looklook 的功能模块。每个阶段专注一个模块，逐步完善，最终得到一个完整的微服务系统。

### 📋 go-zero-looklook 项目功能模块分析

**核心功能模块**（按学习顺序，分阶段实现）：

**第一阶段（第 1-3 周）**：基础功能

1. **项目基础搭建**：项目结构、配置、数据库连接
2. **用户认证模块**：注册、登录、JWT Token
3. **用户管理模块**：用户信息、用户列表、用户详情
4. **认证中间件**：JWT 验证、会话管理

**第二阶段（第 4-5 周）**：权限和商品 5. **权限管理模块**：角色、权限、菜单管理 6. **商品管理模块**：商品 CRUD、分类管理

**第三阶段（第 6-7 周）**：订单和高级功能 7. **订单管理模块**：订单创建、支付、状态管理 8. **文件管理模块**：文件上传、下载 9. **数据统计模块**：Dashboard、报表

**第四阶段（第 8 周）**：微服务和部署 10. **RPC 服务**：微服务拆分、服务间调用 11. **部署优化**：Docker、CI/CD

**学习策略**：

- 先分析 go-zero-looklook 的某个功能实现
- 然后从零开始，在 AI 指导下逐步实现相同功能
- 对比学习，理解差异和优化点
- 每个阶段完成后，项目可以独立运行

### 📁 项目初始目录结构

```
user-service/
├── go.mod
├── go.sum
├── .env                    # 环境变量配置（可选）
├── service/
│   └── user/
│       └── api/
│           ├── user.api    # API 定义文件
│           ├── user.go     # 服务入口
│           ├── etc/
│           │   └── user-api.yaml  # 配置文件
│           └── internal/
│               ├── config/
│               │   └── config.go
│               ├── handler/
│               │   └── routes.go
│               ├── logic/
│               │   ├── registerlogic.go
│               │   ├── loginlogic.go
│               │   └── userlistlogic.go
│               └── svc/
│                   └── servicecontext.go
├── common/
│   ├── db/
│   │   └── db.go          # 数据库连接
│   ├── redis/
│   │   └── redis.go       # Redis 连接
│   ├── tool/
│   │   ├── jwt.go         # JWT 工具
│   │   └── snowflake.go   # 雪花ID生成
│   └── middleware/
│       └── authmiddleware.go  # 认证中间件
└── model/
    └── user.go            # 用户模型
```

### 🚀 开始提示（复制此内容给 AI）

```
我要开始学习项目：go-zero-admin（模仿 go-zero-looklook）

## 项目背景
这是一个学习项目，目标是**从0到1，在AI指导下一步一步模仿 go-zero-looklook 的功能实现**。
通过模仿这个开源项目，掌握 go-zero 框架的 REST API 开发，为后续理解公司项目做准备。

**重要**：请先阅读当前目录下的 learn.md 文件，了解完整的学习计划和目标。

## 参考项目（必须参考）
1. **主要参考项目**：go-zero-looklook
   - GitHub 地址：https://github.com/Mikaelemmmm/go-zero-looklook
   - 这是一个基于 go-zero 的完整微服务项目
   - 技术栈与公司项目高度一致
   - **学习方式**：分析它的某个功能实现 → 在AI指导下从零实现相同功能 → 对比学习

2. **公司项目参考**：当前工作区的项目（如果有的话）
   - service/admin/api - API 网关服务结构
   - common/ - 公共模块（db、redis、tool、middleware）
   - orm/ - 数据模型和查询
   - 可以参考其配置方式、ServiceContext 设计、中间件实现

## 学习方式
1. **先分析**：分析 go-zero-looklook 的某个功能是如何实现的（项目结构、代码组织、业务逻辑）
2. **再实现**：在AI指导下，从零开始实现相同的功能
3. **对比学习**：对比 go-zero-looklook 的实现方式，理解差异和优化点

## 当前状态
已创建空目录，准备从零开始模仿 go-zero-looklook 的项目。
当前目录下已有 learn.md 文件，包含了完整的学习计划。

## 第一步：分析 go-zero-looklook 的项目结构
请帮我：
1. 分析 go-zero-looklook 的项目结构（service 目录、common 目录、model 目录等）
2. 分析它的配置文件结构（yaml 配置、环境变量等）
3. 分析它的数据库连接方式（GORM + MySQL）
4. 分析它的 Redis 连接方式
5. 分析它的基础工具（JWT、雪花ID等）

然后指导我创建相同的项目结构。

## 技术栈要求（必须匹配公司项目）
- go-zero v1.5.2（不要使用最新版本）
- GORM（gorm.io/gorm）
- MySQL（gorm.io/driver/mysql）
- Redis（github.com/gomodule/redigo/redis）
- JWT（github.com/golang-jwt/jwt/v4）
- 雪花算法 ID 生成（github.com/bwmarrin/snowflake）

## 开发规范
1. **模仿学习**：先看 go-zero-looklook 的实现，再自己实现
2. **理解代码**：不要直接复制，要理解每一行代码的作用
3. **对比学习**：实现后对比 go-zero-looklook 的代码，找出差异
4. **逐步实现**：一次一个小功能，确保每个功能都能正常运行
5. **参考 learn.md**：按照 learn.md 中的学习计划逐步实现

请从第一步开始，先分析 go-zero-looklook 的项目结构，然后指导我创建相同的结构。
```

### 📝 功能清单（模仿 go-zero-looklook）

#### 阶段一：项目基础搭建（第 1 周，前 2 天）

**目标**：搭建项目骨架，参考 go-zero-looklook 的项目结构

- [ ] 1.1 分析 go-zero-looklook 的项目结构
- [ ] 1.2 初始化项目（go.mod、目录结构）
- [ ] 1.3 配置管理（config.go、yaml 配置，参考 go-zero-looklook）
- [ ] 1.4 数据库连接（GORM + MySQL，参考 go-zero-looklook 的 common/db）
- [ ] 1.5 Redis 连接（参考 go-zero-looklook 的 common/redis）
- [ ] 1.6 基础工具（雪花 ID、JWT，参考 go-zero-looklook 的 common/utils）

#### 阶段二：用户认证模块（第 1 周，后 3 天）

**目标**：模仿 go-zero-looklook 的用户注册和登录功能

- [ ] 2.1 分析 go-zero-looklook 的用户模型和注册逻辑
- [ ] 2.2 实现用户模型（参考 go-zero-looklook 的 model/user）
- [ ] 2.3 实现用户注册 API（参考 go-zero-looklook 的 register 实现）
- [ ] 2.4 实现用户登录 API（参考 go-zero-looklook 的 login 实现）
- [ ] 2.5 实现 JWT Token 生成和验证（参考 go-zero-looklook 的 jwt 实现）

#### 阶段三：用户管理模块（第 2 周，前 3 天）

**目标**：模仿 go-zero-looklook 的用户管理功能

- [ ] 3.1 分析 go-zero-looklook 的用户列表实现
- [ ] 3.2 实现用户列表 API（分页、搜索，参考 go-zero-looklook）
- [ ] 3.3 实现用户详情 API（参考 go-zero-looklook）
- [ ] 3.4 实现用户更新 API（参考 go-zero-looklook）
- [ ] 3.5 实现用户删除 API（参考 go-zero-looklook）

#### 阶段四：认证中间件和高级功能（第 2 周，后 2 天）

**目标**：模仿 go-zero-looklook 的中间件和会话管理

- [ ] 4.1 分析 go-zero-looklook 的认证中间件实现
- [ ] 4.2 实现 JWT 认证中间件（参考 go-zero-looklook 的 middleware）
- [ ] 4.3 实现 Redis 会话管理（参考 go-zero-looklook 的 session 实现）
- [ ] 4.4 实现错误处理中间件（参考 go-zero-looklook 的 error handling）
- [ ] 4.5 实现请求日志中间件（参考 go-zero-looklook 的 logging）

#### 阶段五：RPC 服务（第 3 周）

**目标**：模仿 go-zero-looklook 的微服务架构

- [ ] 5.1 分析 go-zero-looklook 的 RPC 服务结构
- [ ] 5.2 定义用户 RPC 服务（.proto，参考 go-zero-looklook 的 rpc 定义）
- [ ] 5.3 实现用户 RPC 服务（参考 go-zero-looklook 的 rpc 实现）
- [ ] 5.4 API 服务调用 RPC 服务（参考 go-zero-looklook 的服务调用方式）
- [ ] 5.5 服务间错误处理（参考 go-zero-looklook 的错误处理）

### 🎯 里程碑：可以开始看公司项目

**完成阶段一、二、三后（约 2 周），你可以：**

- ✅ 看懂公司项目的 `service/admin/api` 部分
- ✅ 理解 go-zero REST API 的开发模式
- ✅ 理解 GORM + Redis + JWT 的使用
- ✅ 开始看公司项目的业务逻辑，熟悉业务

**学习建议**：

- 白天看公司项目，理解业务
- 晚上继续学习 RPC，完成阶段四
- 对比公司项目的实现方式，加深理解

---

---

## 项目阶段划分说明

> **注意**：这是一个完整的项目，只是分阶段实现不同功能模块
>
> 每个阶段都是在前一阶段的基础上继续完善，最终得到一个完整的系统
>
> **技术栈**：go-zero（REST + RPC）+ GORM + MySQL + Redis + 定时任务
>
> **参考项目**：go-zero-looklook、公司项目

### 📁 项目初始目录结构

```
mall-service/
├── go.mod
├── go.sum
├── .env
├── service/
│   ├── admin/
│   │   └── api/              # API 网关服务
│   │       ├── admin.api
│   │       ├── admin.go
│   │       └── internal/
│   ├── user/
│   │   └── rpc/              # 用户 RPC 服务
│   │       ├── user.proto
│   │       ├── user.go
│   │       └── internal/
│   ├── product/
│   │   └── rpc/              # 商品 RPC 服务
│   │       ├── product.proto
│   │       ├── product.go
│   │       └── internal/
│   ├── order/
│   │   └── rpc/              # 订单 RPC 服务
│   │       ├── order.proto
│   │       ├── order.go
│   │       └── internal/
│   └── schedule/
│       └── schedule.go       # 定时任务服务
├── common/                   # 公共模块
│   ├── db/
│   ├── redis/
│   ├── tool/
│   └── middleware/
└── model/                    # 数据模型
    ├── user.go
    ├── product.go
    └── order.go
```

### 🚀 继续项目的提示（当完成第一阶段后使用）

```
我要继续完善项目：模仿 go-zero-looklook 项目

## 项目背景
这是一个完整的项目，分阶段实现 go-zero-looklook 的功能模块。
已完成第一阶段（项目基础、用户认证、用户管理），现在要继续实现其他功能模块。

## 参考项目（必须参考）
1. **主要参考项目**：go-zero-looklook
   - GitHub 地址：https://github.com/Mikaelemmmm/go-zero-looklook
   - 继续分析它的其他功能模块实现
   - 在AI指导下从零实现相同功能

2. **公司项目参考**：当前工作区的 /Users/losan/Desktop/xiaoiron/admin
   - 对比公司项目的实现方式
   - 理解差异和优化点

## 当前状态
已完成项目一的基础功能（用户注册、登录、用户管理），现在要继续模仿 go-zero-looklook 的其他功能。

## 下一步要模仿的功能
请告诉我下一步应该模仿 go-zero-looklook 的哪个功能模块，然后：
1. 先分析 go-zero-looklook 的这个功能是如何实现的
2. 在AI指导下从零实现相同功能
3. 对比学习，理解差异

## 技术栈要求（必须匹配公司项目）
- go-zero v1.5.2（REST + RPC）
- GORM + MySQL
- Redis
- JWT 认证
- 定时任务（robfig/cron）
- Protocol Buffers

请告诉我下一步应该模仿哪个功能模块。
```

### 📝 功能清单（继续模仿 go-zero-looklook）

> **说明**：这是项目的第一阶段之后的继续，在同一个项目中实现更多功能模块

#### 阶段五：权限管理模块（第 4 周）

**目标**：模仿 go-zero-looklook 的权限管理功能

- [ ] 1.1 分析 go-zero-looklook 的角色和权限模型
- [ ] 1.2 实现角色管理（参考 go-zero-looklook 的 role 实现）
- [ ] 1.3 实现权限管理（参考 go-zero-looklook 的 permission 实现）
- [ ] 1.4 实现菜单管理（参考 go-zero-looklook 的 menu 实现）
- [ ] 1.5 实现权限中间件（参考 go-zero-looklook 的权限验证）

#### 阶段六：商品管理模块（第 5 周）

**目标**：模仿 go-zero-looklook 的商品管理功能

- [ ] 2.1 分析 go-zero-looklook 的商品模型和业务逻辑
- [ ] 2.2 实现商品 CRUD（参考 go-zero-looklook 的 product 实现）
- [ ] 2.3 实现商品分类管理（参考 go-zero-looklook 的 category 实现）
- [ ] 2.4 实现商品库存管理（参考 go-zero-looklook 的 inventory 实现）

#### 阶段七：订单管理模块（第 6 周）

**目标**：模仿 go-zero-looklook 的订单管理功能

- [ ] 3.1 分析 go-zero-looklook 的订单模型和流程
- [ ] 3.2 实现订单创建（参考 go-zero-looklook 的 order 实现）
- [ ] 3.3 实现订单支付（参考 go-zero-looklook 的 payment 实现）
- [ ] 3.4 实现订单状态管理（参考 go-zero-looklook 的 order status）
- [ ] 3.5 实现购物车功能（参考 go-zero-looklook 的 cart 实现）

#### 阶段八：高级功能（第 7 周）

**目标**：模仿 go-zero-looklook 的高级功能

- [ ] 4.1 分析 go-zero-looklook 的文件管理实现
- [ ] 4.2 实现文件上传下载（参考 go-zero-looklook 的 file 实现）
- [ ] 4.3 实现数据统计 Dashboard（参考 go-zero-looklook 的 dashboard）
- [ ] 4.4 实现定时任务（参考 go-zero-looklook 的 schedule，对比公司项目）

#### 阶段九：优化和部署（第 8 周）

**目标**：学习部署和优化

- [ ] 5.1 分析 go-zero-looklook 的 Docker 配置
- [ ] 5.2 实现 Docker 部署（参考 go-zero-looklook 的 docker-compose）
- [ ] 5.3 性能优化（对比公司项目的优化方式）
- [ ] 5.4 错误处理完善（对比 `common/error/`）
- [ ] 5.5 CI/CD 配置（可选）

---

## 推荐开源项目

### 1. go-zero-looklook ⭐⭐⭐⭐⭐

- **仓库**：https://github.com/Mikaelemmmm/go-zero-looklook
- **Star**: 3.5k+
- **技术栈匹配度**：95%
- **推荐理由**：
  - 技术栈与公司项目高度一致
  - 代码质量好，结构清晰
  - 有文档和教程
  - 适合从零模仿

### 2. zero-examples ⭐⭐⭐⭐

- **仓库**：https://github.com/zeromicro/zero-examples
- **Star**: 1k+
- **推荐理由**：官方维护，代码规范

### 3. go-zero-mall ⭐⭐⭐

- **仓库**：https://github.com/nivin-studio/go-zero-mall
- **Star**: 500+
- **推荐理由**：商城业务场景，微服务拆分示例

---

## 学习资源

1. **go-zero 官方文档**：https://go-zero.dev/
2. **GORM 文档**：https://gorm.io/
3. **go-zero-looklook 教程**：项目 README 和文档
4. **Protocol Buffers 文档**：https://protobuf.dev/
5. **Redis 文档**：https://redis.io/docs/

---

## 实践建议

1. ✅ **从空目录开始**，不要直接复制代码
2. ✅ **一次一个小功能**，像正常开发一样
3. ✅ **理解每一行代码**，不要只是复制粘贴
4. ✅ **对比公司项目**，找出差异和优化点
5. ✅ **记录问题和解决方案**，形成学习笔记
6. ✅ **定期回顾**，巩固知识点

---

## 使用说明

### 如何开始一个新项目？

1. **创建空目录**：按照项目初始目录结构创建文件夹
2. **复制开始提示**：复制对应项目的"开始提示"内容
3. **发送给 AI**：将开始提示发送给 AI，开始学习
4. **逐步实现**：按照功能清单，一次实现一个小功能

### 如何继续之前的项目？

1. **查看 learn.md**：找到对应的项目章节
2. **查看功能清单**：了解当前进度和下一步任务
3. **发送上下文**：告诉 AI 当前项目状态和下一步要做什么

---

## 学习进度跟踪

### 完整项目：go-zero-admin（第 1-8 周）

**第一阶段（第 1-3 周）**：基础功能

- [ ] 阶段一：项目基础搭建
- [ ] 阶段二：用户认证模块
- [ ] 阶段三：用户管理模块
- [ ] 阶段四：认证中间件和 RPC 服务

**第二阶段（第 4-5 周）**：权限和商品

- [ ] 阶段五：权限管理模块
- [ ] 阶段六：商品管理模块

**第三阶段（第 6-7 周）**：订单和高级功能

- [ ] 阶段七：订单管理模块
- [ ] 阶段八：高级功能（文件管理、数据统计）

**第四阶段（第 8 周）**：优化和部署

- [ ] 阶段九：优化和部署

**扩展场景**（完成基本功能后）：

- [ ] 场景一：数据统计和报表系统
- [ ] 场景二：多层级组织架构管理
- [ ] 场景三：消息通知和日志系统

**🎯 里程碑**：完成阶段一、二、三后（约 2 周），可以开始看公司项目

---

## 📊 学习时间总结

| 项目       | 原计划 | **实际预估**     | 关键节点                         |
| ---------- | ------ | ---------------- | -------------------------------- |
| **项目一** | 8 周   | **2-3 周**       | 完成阶段一、二、三后可看公司项目 |
| **项目二** | 4 个月 | **1-2 个月**     | 边学边看公司项目                 |
| **总计**   | 6 个月 | **1.5-2.5 个月** | 专注学习可更快                   |

## 🎯 什么时候可以开始看公司项目？

**✅ 最佳时机：完成项目一的阶段一、二、三后（约 2 周）**

**此时你已经掌握**：

- ✅ go-zero REST API 开发
- ✅ GORM + MySQL 在 go-zero 中的使用
- ✅ Redis + JWT 认证
- ✅ 中间件开发

**可以看懂公司项目的**：

- ✅ `service/admin/api` - API 网关服务
- ✅ `common/` - 公共模块（db、redis、tool、middleware）
- ✅ `orm/` - 数据模型和查询
- ✅ 业务逻辑实现方式

**学习策略**：

1. **第 1-2 周**：专注学习，完成项目一的基础功能
2. **第 3 周开始**：白天看公司项目 API 服务，晚上继续学 RPC
3. **第 4 周开始**：实现项目二的同时，对比公司项目的实现方式

---

---

**最后更新**：2024 年
**当前学习阶段**：准备开始项目
**项目名称**：go-zero-admin
