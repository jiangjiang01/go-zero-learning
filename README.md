# Go-Zero 学习项目

> **项目名称**：go-zero-learning
>
> **学习目标**：从 0 到 1，在 AI 指导下一步一步模仿 go-zero-looklook 项目，掌握 go-zero 框架开发
>
> **技术栈**：go-zero v1.9.3 + GORM + MySQL + Redis + JWT + gRPC
>
> **参考项目**：[go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook)

---

## 🎯 核心教学原则（AI 必须遵守）

> **⚠️ 这是最高优先级的原则，每次回答都必须遵守！**

### 角色定位

- **AI 角色**：资深开发，负责引导和教学
- **用户角色**：学习者，跟着 AI 一步步实现功能
- **学习方式**：边做边学，通过"踩坑"加深理解

### 代码输出策略（必须遵守）

> **核心原则**：分步学习，先暴露问题，再优化方案

#### 1. 分小步输出

- ❌ **禁止**：一次性给出完整功能的代码
- ✅ **必须**：每次只输出一小步，逐步实现

#### 2. 先给"low 的写法"（让学习者踩坑）

- ✅ 先给出能运行但有问题的版本（如静态检查警告、性能问题、设计缺陷）
- ✅ 让学习者先遇到问题，自己思考原因
- ✅ 不要一上来就提供最优解

#### 3. 暴露问题

- ✅ 让学习者看到警告/错误
- ✅ 让学习者意识到这种写法的问题

#### 4. 解释问题

- ✅ 说明为什么会有警告/错误
- ✅ 解释这种写法的风险和问题
- ✅ 解释为什么这种写法不好

#### 5. 给出优化方案

- ✅ 提供更好的实现方式
- ✅ 解释为什么优化方案更好
- ✅ 对比两种写法的差异

### 学习流程示例

```
步骤1：给出简单实现（有警告/问题）
  ↓
步骤2：学习者遇到警告，思考原因
  ↓
步骤3：解释问题（为什么不好）
  ↓
步骤4：给出优化方案（为什么更好）
```

**慢一点没关系，重点是学到更多知识点！**

---

## 🚀 新会话启动指南

### 对于 AI

**新会话开始时，必须：**

1. ✅ 立即读取 `README.md` 文档，了解项目状态
2. ✅ 查看"当前项目进度"，了解已完成的功能
3. ✅ 确认下一步要做什么
4. ✅ **严格遵守"代码输出策略"**（最高优先级）
5. ✅ 以"资深开发"的身份，引导学习者一步步实现

### 对于用户

**继续项目时，可以这样说：**

```
我在学习 go-zero 项目，请先查看 README.md 了解当前进度。
当前需要实现 [具体功能]。
```

---

## 📁 项目结构

```
go-zero-learning/
├── backend/                    # 后端项目（go.mod 在 backend/）
│   ├── common/                # 公共模块
│   │   ├── db/               # 数据库连接
│   │   ├── jwt/              # JWT 工具
│   │   ├── middleware/       # 中间件
│   │   ├── errorx/          # 错误处理
│   │   ├── response/         # 统一响应
│   │   └── ctxdata/          # Context 数据管理
│   ├── model/                # 数据模型
│   └── service/               # 服务
│       └── user/
│           └── api/          # 用户 API 服务
├── frontend/                  # 前端项目
│   ├── src/
│   │   ├── api/              # API 接口
│   │   ├── views/            # 页面组件
│   │   │   └── system/       # 系统管理页面
│   │   │       ├── user/     # 用户管理
│   │   │       └── role/     # 角色管理
│   │   └── router/           # 路由配置
└── learn/                     # 学习代码（独立项目）
```

**项目特点**：

- 单模块结构（参考 go-zero-looklook）
- `backend/` 作为后端根目录
- `go.mod` 在 `backend/` 目录下

---

## 🌐 RESTful API 设计规范

### 设计原则

1. **使用复数资源名**：`/api/users` 而不是 `/api/user`
2. **使用 HTTP 方法表示操作**：POST（创建）、GET（获取）、PUT（更新）、DELETE（删除）
3. **使用路径参数表示资源 ID**：`/api/users/:id`
4. **使用特殊资源表示当前用户**：`/api/users/me`

### 当前 API 路径规范

#### 用户管理 API

| 功能             | HTTP 方法 | 路径               | 说明                           |
| ---------------- | --------- | ------------------ | ------------------------------ |
| 用户注册         | POST      | `/api/users`       | 创建新用户                     |
| 用户登录         | POST      | `/api/users/login` | 登录（子资源操作）             |
| 获取当前用户信息 | GET       | `/api/users/me`    | 获取当前认证用户信息           |
| 获取用户列表     | GET       | `/api/users`       | 获取用户列表（支持分页和搜索） |
| 获取指定用户详情 | GET       | `/api/users/:id`   | 获取指定用户信息               |
| 更新当前用户信息 | PUT       | `/api/users/me`    | 更新当前认证用户信息           |
| 更新指定用户信息 | PUT       | `/api/users/:id`   | 更新指定用户信息（防止自更新） |
| 删除用户         | DELETE    | `/api/users/:id`   | 删除用户（防止自删除）         |

#### 角色管理 API

| 功能         | HTTP 方法 | 路径             | 说明                                     |
| ------------ | --------- | ---------------- | ---------------------------------------- |
| 创建角色     | POST      | `/api/roles`     | 创建新角色（需要认证）                   |
| 获取角色列表 | GET       | `/api/roles`     | 获取角色列表（支持分页和搜索，需要认证） |
| 获取角色详情 | GET       | `/api/roles/:id` | 获取指定角色信息（需要认证）             |
| 更新角色     | PUT       | `/api/roles/:id` | 更新角色信息（需要认证）                 |
| 删除角色     | DELETE    | `/api/roles/:id` | 删除角色（需要认证）                     |

---

## 🔧 统一响应格式和错误处理机制

### 统一响应格式

```json
{
  "code": 0,              // 状态码，0 表示成功，非 0 表示失败
  "message": "success",   // 消息
  "data": {...},         // 数据（成功时返回，失败时为空）
  "timestamp": 1705939200 // 时间戳
}
```

### 错误码分类

- `0`：成功
- `1000-1999`：通用错误码
- `2000-2999`：用户相关错误码
- `3000-3999`：权限相关错误码（后续扩展）
- `4000-4999`：商品相关错误码（后续扩展）

### 错误码列表

#### 通用错误码（1000-1999）

- `1001` - 参数错误
- `1002` - 未授权
- `1003` - 禁止访问
- `1004` - 资源不存在
- `1005` - 内部错误
- `1006` - Token 无效或已过期

#### 用户相关错误码（2000-2999）

- `2001` - 用户不存在
- `2002` - 用户已存在
- `2003` - 用户名已存在
- `2004` - 邮箱已存在
- `2005` - 密码错误
- `2006` - 密码长度不足
- `2007` - 邮箱格式不正确
- `2008` - 不能删除自己
- `2009` - 没有需要更新的字段
- `2010` - 未找到用户信息

#### 角色相关错误码（3000-3999）

- `3001` - 角色不存在
- `3002` - 角色已存在
- `3003` - 角色名称已存在
- `3004` - 角色代码已存在
- `3005` - 没有需要更新的字段

### 使用方式

#### 在 Logic 中使用错误处理

```go
import "go-zero-learning/common/errorx"

// 返回预定义错误
return nil, errorx.ErrUserNotFound

// 创建自定义错误
return nil, errorx.NewBusinessError(errorx.CodeInvalidParam, "参数错误")
```

#### 在 Handler 中使用

```go
import (
    "go-zero-learning/common/errorx"
    "go-zero-learning/common/response"
)

// 错误处理
if err != nil {
    errorx.HandleError(w, r, err)
    return
}

// 成功响应
response.OkJson(w, r, resp)
```

---

## 📝 当前项目进度（实时更新）

### 项目结构

- ✅ 项目根目录：`go-zero-learning/`
- ✅ 后端根目录：`backend/`（go.mod 在 backend/）
- ✅ 模块结构：单模块结构（参考 go-zero-looklook）

### 已完成功能

#### ✅ 项目基础搭建

- [x] 目录结构创建（backend/common、backend/model、backend/service）
- [x] go.mod 配置（模块名：go-zero-learning）
- [x] 数据库连接模块（backend/common/db）
- [x] JWT 工具模块（backend/common/jwt）
- [x] 用户模型（backend/model/user）

#### ✅ 用户 API 服务框架

- [x] API 定义（user.api）
- [x] 代码生成（handler、logic、svc）
- [x] ServiceContext 配置（数据库连接、自动迁移）
- [x] 服务能正常运行（端口 8888）

#### ✅ 用户认证功能

- [x] 用户注册逻辑（密码加密 bcrypt）
- [x] 用户登录逻辑（JWT Token 生成）
- [x] 参数验证（go-zero 自动验证）
- [x] 错误处理（用户名/邮箱重复检查）

#### ✅ 认证中间件和用户信息

- [x] 认证中间件（JWT 验证）
- [x] Context 数据管理（ctxdata 包，避免 key 冲突）
- [x] 获取用户信息逻辑（从 context 获取用户 ID）

#### ✅ 用户管理功能

- [x] 用户列表 API（分页、搜索）
- [x] 用户更新 API（更新邮箱和密码）
- [x] 用户详情 API（根据 ID 获取）
- [x] 用户删除 API（防止自删除）
- [x] RESTful API 重构（统一使用 RESTful 规范）

#### ✅ 错误处理和响应格式

- [x] 统一错误处理模块（backend/common/errorx）
- [x] 错误码定义（通用错误码 1000-1999，用户错误码 2000-2999，角色错误码 3000-3999）
- [x] 统一错误响应格式（code + message + timestamp）
- [x] 自动 HTTP 状态码映射
- [x] 所有 handler 和 logic 使用统一错误处理

#### ✅ 角色管理功能（后端）

- [x] 角色模型定义（backend/model/role）
- [x] 角色 API 定义（user.api 中添加角色相关类型和路由）
- [x] 角色错误码定义（3001-3005）
- [x] 角色 CRUD API（创建、列表、详情、更新、删除）
- [x] 角色列表分页和搜索功能
- [x] 角色名称和代码唯一性检查
- [x] 所有 API 测试通过

#### ✅ 角色管理功能（前端）

- [x] 角色管理 API 接口（frontend/src/api/role.ts）
- [x] 角色管理页面（frontend/src/views/system/role/index.vue）
- [x] 角色新增/编辑对话框（frontend/src/views/system/role/components/RoleDialog.vue）
- [x] 路由配置（/system/role）
- [x] 表单验证（名称、代码、描述）
- [x] 所有功能测试通过

### 待完成功能

#### 阶段二：权限管理（当前阶段）

- [x] 角色管理（角色 CRUD）✅
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

- [x] 错误处理完善 ✅
- [ ] 日志系统
- [ ] Docker 部署
- [ ] 性能优化

### 当前问题/注意事项

- 配置文件字段名：使用 `dataSource`（小写驼峰）
- 运行方式：`cd backend/service/user/api && go run user-api.go -f etc/user-api.yaml`
- 数据库：MySQL 3307 端口，数据库名 testdb
- go-zero 参数验证：可选字段（optional）在 JSON 中缺失时会报错，需要在请求中包含所有字段（临时方案）

### 下一步计划

**当前阶段**：阶段二 - 权限管理

**下一步**：开始实现权限管理功能（权限 CRUD）

**最后更新**：2025-01-22  
**当前状态**：

- 阶段一（用户认证和管理）全部完成 ✅
- 阶段二（权限管理）进行中：
  - 角色管理（后端+前端）已完成 ✅
  - 权限管理（待实现）
  - 菜单管理（待实现）
  - 权限中间件（待实现）

---

## 📚 最小必要知识

### 1. go-zero 项目结构

```
go-zero 项目：
├── service/
│   └── user/
│       └── api/
│           ├── user.api          # API 定义文件
│           ├── user.go           # 入口
│           └── internal/
│               ├── config/       # 配置
│               ├── handler/      # HTTP 处理（自动生成）
│               ├── logic/        # 业务逻辑（你写这里）
│               └── svc/           # ServiceContext（依赖注入）
```

### 2. 开发流程

```
go-zero：
1. 写 .api 文件：定义 API
2. 运行 goctl api go：生成 handler、logic 骨架
3. 在 logic 中写业务逻辑
```

### 3. ServiceContext 模式（依赖注入）

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

---

## 🔗 学习资源

1. **go-zero 官方文档**：https://go-zero.dev/
2. **GORM 文档**：https://gorm.io/
3. **go-zero-looklook**：https://github.com/Mikaelemmmm/go-zero-looklook
