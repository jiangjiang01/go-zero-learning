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

```text
步骤1：给出简单实现（有警告/问题）
  ↓
步骤2：学习者遇到警告，思考原因
  ↓
步骤3：解释问题（为什么不好）
  ↓
步骤4：给出优化方案（为什么更好）
```

**慢一点没关系，重点是学到更多知识点！**

### 🎓 成功教学方法论（基于商品管理实践总结）

#### 核心教学策略

1. **问题驱动学习**
   - 先给出"能运行但有问题"的代码版本
   - 让学习者亲自遇到问题（如类型选择不当、缺少约束等）
   - 引导学习者思考问题原因
   - 最后给出优化方案并详细解释

2. **分层递进教学**
   - 数据模型设计 → API 定义 → 业务逻辑 → 前端实现
   - 每一层都先给基础版本，再优化完善
   - 重点关注层与层之间的关系和数据流转

3. **关键概念重点突破**
   - **可选字段处理**：通过指针类型区分"没传"vs"零值"
   - **错误处理设计**：统一错误码、分类管理、业务语义化
   - **RESTful 设计**：路由规范、HTTP 方法、参数验证
   - **防御性编程**：双重验证、边界检查、业务约束

4. **实战中学习最佳实践**
   - 数据库设计：字段类型选择（int64 vs float64）、约束设计
   - 业务逻辑：参数验证、唯一性检查、事务处理
   - 前端开发：组件设计、状态管理、用户体验

#### 教学节奏控制

- **慢节奏**：每个概念都要让学习者充分理解
- **多互动**：鼓励学习者提问、质疑、思考
- **重实践**：学习者必须亲自动手敲代码
- **善总结**：每个阶段结束后总结关键知识点

#### 错误处理教学重点

- **业务错误 vs 系统错误**：如何设计错误码和错误信息
- **前端错误处理**：避免重复提示，保持用户体验
- **调试技巧**：如何通过错误信息快速定位问题

#### 成功指标

- 学习者能独立分析代码问题
- 学习者掌握 go-zero 核心开发模式
- 学习者具备生产级代码质量意识
- 学习者能举一反三应用到其他模块

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

```text
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
│           ├── api/          # 用户 API 服务
│           └── user-rpc/     # 用户 RPC 服务
├── frontend/                  # 前端项目
│   ├── src/
│   │   ├── api/              # API 接口
│   │   ├── views/            # 页面组件
│   │   │   └── system/       # 系统管理页面
│   │   │       ├── user/     # 用户管理
│   │   │       └── role/     # 角色管理
│   │   └── router/           # 路由配置
├── scripts/                   # 测试脚本目录
│   ├── README.md             # 测试脚本说明文档
│   └── test-*.sh             # 各种测试脚本
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

#### 权限管理 API

| 功能         | HTTP 方法 | 路径                  | 说明                                     |
| ------------ | --------- | --------------------- | ---------------------------------------- |
| 创建权限     | POST      | `/api/permissions`    | 创建新权限（需要认证）                   |
| 获取权限列表 | GET       | `/api/permissions`    | 获取权限列表（支持分页和搜索，需要认证） |
| 获取权限详情 | GET       | `/api/permissions/:id` | 获取指定权限信息（需要认证）             |
| 更新权限     | PUT       | `/api/permissions/:id` | 更新权限信息（需要认证）                 |
| 删除权限     | DELETE    | `/api/permissions/:id` | 删除权限（需要认证）                     |

#### 菜单管理 API

| 功能         | HTTP 方法 | 路径               | 说明                                                           |
| ------------ | --------- | ------------------ | -------------------------------------------------------------- |
| 创建菜单     | POST      | `/api/menus`       | 创建新菜单（需要认证）                                         |
| 获取菜单列表 | GET       | `/api/menus`       | 获取菜单列表（支持分页和搜索，需要认证）                       |
| 获取全部菜单 | GET       | `/api/menus?all=true` | 获取所有菜单（不分页，用于构建树形结构，需要认证）           |
| 获取菜单详情 | GET       | `/api/menus/:id`   | 获取指定菜单信息（需要认证）                                   |
| 更新菜单     | PUT       | `/api/menus/:id`   | 更新菜单信息（所有字段必填，需要认证）                         |
| 删除菜单     | DELETE    | `/api/menus/:id`   | 删除菜单（有子菜单时不能删除，需要认证）                       |

#### 用户角色管理 API

| 功能           | HTTP 方法 | 路径                      | 说明                                       |
| -------------- | --------- | ------------------------- | ------------------------------------------ |
| 给用户分配角色 | POST      | `/api/users/:id/roles`    | 给指定用户分配角色（需要认证）             |
| 获取用户角色   | GET       | `/api/users/:id/roles`    | 获取指定用户的所有角色列表（需要认证）     |
| 移除用户角色   | DELETE    | `/api/users/:id/roles/:role_id` | 移除指定用户的某个角色（需要认证）         |

#### 角色权限管理 API

| 功能           | HTTP 方法 | 路径                            | 说明                                       |
| -------------- | --------- | ------------------------------- | ------------------------------------------ |
| 给角色分配权限 | POST      | `/api/roles/:id/permissions`    | 给指定角色分配权限（需要认证）             |
| 获取角色权限   | GET       | `/api/roles/:id/permissions`    | 获取指定角色的所有权限列表（需要认证）     |
| 移除角色权限   | DELETE    | `/api/roles/:id/permissions/:permission_id` | 移除指定角色的某个权限（需要认证）         |

#### 商品管理 API

| 功能         | HTTP 方法 | 路径                    | 说明                                     |
| ------------ | --------- | ----------------------- | ---------------------------------------- |
| 创建商品     | POST      | `/api/products`         | 创建新商品（需要认证和 product:create 权限） |
| 获取商品列表 | GET       | `/api/products`         | 获取商品列表（支持分页和搜索，需要认证和 product:list 权限） |
| 获取商品详情 | GET       | `/api/products/:id`     | 获取指定商品信息（需要认证和 product:list 权限） |
| 更新商品     | PUT       | `/api/products/:id`     | 更新商品信息（需要认证和 product:update 权限） |
| 更新商品状态 | PUT       | `/api/products/:id/status` | 更新商品状态（需要认证和 product:update 权限） |
| 删除商品     | DELETE    | `/api/products/:id`     | 删除商品（需要认证和 product:delete 权限） |

#### 订单管理 API

| 功能         | HTTP 方法 | 路径                    | 说明                                     |
| ------------ | --------- | ----------------------- | ---------------------------------------- |
| 创建订单     | POST      | `/api/orders`           | 创建新订单（需要认证和 order:create 权限） |
| 获取订单列表 | GET       | `/api/orders`           | 获取订单列表（支持分页、搜索、状态筛选，需要认证和 order:list 权限） |
| 获取订单详情 | GET       | `/api/orders/:id`       | 获取指定订单信息（需要认证和 order:list 权限，只能查看自己的订单） |
| 更新订单状态 | PUT       | `/api/orders/:id/status` | 更新订单状态（需要认证和 order:update 权限） |

#### 商品分类管理 API

| 功能         | HTTP 方法 | 路径                    | 说明                                     |
| ------------ | --------- | ----------------------- | ---------------------------------------- |
| 创建分类     | POST      | `/api/categories`       | 创建新商品分类（需要认证和 category:create 权限） |
| 获取分类列表 | GET       | `/api/categories`       | 获取分类列表（支持分页、搜索、全部模式，需要认证和 category:list 权限） |
| 获取分类详情 | GET       | `/api/categories/:id`   | 获取指定分类信息（需要认证和 category:list 权限） |
| 更新分类     | PUT       | `/api/categories/:id`   | 更新分类信息（需要认证和 category:update 权限） |
| 删除分类     | DELETE    | `/api/categories/:id`   | 删除分类（需要认证和 category:delete 权限） |

#### 购物车管理 API

| 功能         | HTTP 方法 | 路径                    | 说明                                     |
| ------------ | --------- | ----------------------- | ---------------------------------------- |
| 添加商品到购物车 | POST      | `/api/cart/items`       | 添加商品到购物车（需要认证和 cart:add 权限） |
| 获取购物车   | GET       | `/api/cart`             | 获取购物车列表（需要认证和 cart:get 权限） |
| 更新购物车项数量 | PUT       | `/api/cart/items/:item_id` | 更新购物车项数量（需要认证和 cart:update 权限） |
| 删除购物车项 | DELETE    | `/api/cart/items/:item_id` | 删除购物车项（需要认证和 cart:delete 权限） |
| 清空购物车   | DELETE    | `/api/cart`             | 清空购物车（需要认证和 cart:delete 权限） |

#### 数据统计 Dashboard API

| 功能         | HTTP 方法 | 路径                    | 说明                                     |
| ------------ | --------- | ----------------------- | ---------------------------------------- |
| 获取统计数据 | GET       | `/api/dashboard/stats`  | 获取 Dashboard 统计数据（订单、商品、用户统计，需要认证） |

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
- `3000-3999`：角色相关错误码
- `4000-4999`：权限相关错误码
- `5000-5999`：菜单相关错误码
- `6000-6999`：商品相关错误码
- `7000-7999`：订单相关错误码
- `8000-8999`：商品分类相关错误码
- `9000-9999`：购物车相关错误码

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

#### 权限相关错误码（4000-4999）

- `4001` - 权限不存在
- `4002` - 权限已存在
- `4003` - 权限名称已存在
- `4004` - 权限代码已存在
- `4005` - 没有需要更新的字段

#### 菜单相关错误码（5000-5999）

- `5001` - 菜单不存在
- `5002` - 菜单已存在
- `5003` - 菜单代码已存在
- `5004` - 没有需要更新的字段
- `5005` - 菜单下有子菜单，不能删除
- `5006` - 不能将父菜单设置为自己的子菜单

#### 商品相关错误码（6000-6999）

- `6001` - 商品不存在
- `6002` - 商品已存在
- `6003` - 商品名称已存在
- `6004` - 商品价格太低
- `6005` - 商品价格太高
- `6006` - 没有需要更新的字段
- `6007` - 商品状态无效

#### 订单相关错误码（7000-7999）

- `7001` - 订单不存在
- `7002` - 订单状态无效
- `7003` - 订单不能取消
- `7004` - 订单不能支付
- `7005` - 订单数量无效
- `7006` - 商品价格已变更
- `7007` - 库存不足
- `7008` - 订单商品列表不能为空
- `7009` - 订单金额不匹配
- `7010` - 订单不属于当前用户

#### 商品分类相关错误码（8000-8999）

- `8001` - 商品分类不存在
- `8003` - 商品分类名称已存在
- `8004` - 商品分类下有子分类，不能删除
- `8005` - 商品分类下有商品，不能删除
- `8006` - 父分类无效
- `8007` - 没有需要更新的字段

#### 购物车相关错误码（9000-9999）

- `9001` - 购物车项不存在
- `9002` - 购物车项数量无效
- `9003` - 购物车项数量过大
- `9004` - 商品不可用（已下架或不存在）

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
- [x] 错误码定义（通用错误码 1000-1999，用户错误码 2000-2999，角色错误码 3000-3999，权限错误码 4000-4999）
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

#### ✅ 权限管理功能（后端）

- [x] 权限模型定义（backend/model/permission）
- [x] 权限 API 定义（user.api 中添加权限相关类型和路由）
- [x] 权限错误码定义（4001-4005）
- [x] 权限 CRUD API（创建、列表、详情、更新、删除）
- [x] 权限列表分页和搜索功能
- [x] 权限名称和代码唯一性检查
- [x] 所有 API 测试通过

#### ✅ 权限管理功能（前端）

- [x] 权限管理 API 接口（frontend/src/api/permission.ts）
- [x] 权限管理页面（frontend/src/views/system/permission/index.vue）
- [x] 权限新增/编辑对话框（frontend/src/views/system/permission/components/PermissionDialog.vue）
- [x] 路由配置（/system/permission）
- [x] 表单验证（名称、代码、描述）
- [x] 分页功能修复（切换每页数量时重置页码）
- [x] 所有功能测试通过

#### ✅ 菜单管理功能（后端）

- [x] 菜单模型定义（backend/model/menu，支持树形结构）
- [x] 菜单 API 定义（user.api 中添加菜单相关类型和路由）
- [x] 菜单错误码定义（5001-5006）
- [x] 菜单 CRUD API（创建、列表、详情、更新、删除）
- [x] 菜单列表分页和搜索功能（支持全部模式和分页模式）
- [x] 菜单代码唯一性检查
- [x] 同级菜单同名检查
- [x] 父菜单存在性验证和循环引用检查
- [x] 删除菜单时子菜单检查（防止误删）
- [x] 禁用菜单时子菜单检查
- [x] 所有 API 测试通过

#### ✅ 菜单管理功能（前端）

- [x] 菜单管理 API 接口（frontend/src/api/menu.ts）
- [x] 菜单管理页面（frontend/src/views/system/menu/index.vue）
- [x] 菜单新增/编辑对话框（frontend/src/views/system/menu/components/MenuDialog.vue）
- [x] 路由配置（/system/menu）
- [x] 表单验证（名称、代码、类型、状态等）
- [x] 父菜单下拉选择（支持搜索和过滤）
- [x] 菜单类型和状态选择
- [x] 时间戳格式化支持（formatDateTime 支持 Unix 时间戳）
- [x] 所有功能测试通过

#### ✅ 用户角色管理功能（后端）

- [x] 用户角色关联模型定义（backend/model/user_role）
- [x] 用户角色 API 定义（user.api 中添加用户角色相关类型和路由）
- [x] 用户角色 CRUD API（分配角色、获取用户角色列表、移除角色）
- [x] 用户角色关联表唯一性检查（防止重复分配）
- [x] 用户和角色存在性验证
- [x] 所有 API 测试通过

#### ✅ 用户角色管理功能（前端）

- [x] 用户角色管理 API 接口（frontend/src/api/user.ts）
- [x] 用户角色分配对话框（frontend/src/views/system/user/components/UserRoleDialog.vue）
- [x] 用户管理页面添加角色分配按钮
- [x] 显示当前用户已分配角色（标签形式）
- [x] 支持添加和移除角色
- [x] 下拉选择框支持搜索过滤可用角色
- [x] 所有功能测试通过

#### ✅ 角色权限管理功能（后端）

- [x] 角色权限关联模型定义（backend/model/role_permission）
- [x] 角色权限 API 定义（user.api 中添加角色权限相关类型和路由）
- [x] 角色权限 CRUD API（分配权限、获取角色权限列表、移除权限）
- [x] 角色权限关联表唯一性检查（防止重复分配）
- [x] 角色和权限存在性验证
- [x] 所有 API 测试通过
- [x] 权限管理测试脚本（test_permissio.sh）已创建并测试通过

#### ✅ 角色权限管理功能（前端）

- [x] 角色权限管理 API 接口（frontend/src/api/role.ts）
- [x] 角色权限分配对话框（frontend/src/views/system/role/components/RolePermissionDialog.vue）
- [x] 角色管理页面添加权限分配按钮
- [x] 显示当前角色已分配权限（标签形式）
- [x] 支持添加和移除权限
- [x] 下拉选择框支持搜索过滤可用权限
- [x] 所有功能测试通过

#### ✅ 权限中间件功能

- [x] 权限验证中间件实现（backend/common/middleware/permissionmiddleware.go）
- [x] 权限验证逻辑（基于用户角色和角色权限）
- [x] 路由权限配置（按权限代码分组路由）
- [x] 性能优化（避免空切片查询数据库）
- [x] 所有路由权限验证配置完成
- [x] 权限验证测试脚本（scripts/test-permission-middleware.sh）
- [x] 测试覆盖：未登录访问、无角色访问、有权限访问、无权限访问等场景
- [x] 所有测试通过（12/12）

#### ✅ 商品管理功能（后端）

- [x] 商品模型定义（backend/model/product）
- [x] 商品 API 定义（user.api 中添加商品相关类型和路由）
- [x] 商品错误码定义（6000-6999）
- [x] 商品 CRUD API（创建、列表、详情、更新、删除）
- [x] 商品列表分页和搜索功能
- [x] 商品名称唯一性检查
- [x] 商品状态管理
- [x] 商品库存管理（添加库存字段，支持创建和更新时设置库存）
- [x] 商品图片管理（支持多图片上传，图片URL存储在数据库中）
- [x] 所有 API 测试通过

#### ✅ 商品管理功能（前端）

- [x] 商品管理 API 接口（frontend/src/api/product.ts）
- [x] 商品管理页面（frontend/src/views/system/product/index.vue）
- [x] 商品新增/编辑对话框（frontend/src/views/system/product/components/ProductDialog.vue）
- [x] 路由配置（/system/product）
- [x] 表单验证（名称、价格、状态）
- [x] 价格单位转换（元 ↔ 分）
- [x] 库存管理（商品列表显示库存，编辑对话框支持设置库存，库存颜色提示）
- [x] **修复 bug：ProductDialog.vue resetForm 函数初始化问题**（函数定义必须在 watch 之前）
- [x] **修复 UI 问题：移除重复错误提示**（只保留一个权限验证错误信息）
- [x] 商品图片上传功能（支持多图片上传、预览、删除）
- [x] 所有功能测试通过

#### ✅ 订单管理功能（后端）

- [x] 订单模型定义（backend/model/order、backend/model/order_item）
- [x] 订单 API 定义（user.api 中添加订单相关类型和路由）
- [x] 订单错误码定义（7000-7999）
- [x] 订单 CRUD API（创建、列表、详情、状态更新）
- [x] 订单列表分页和搜索功能（支持状态筛选、订单编号搜索）
- [x] 订单状态流转验证（防止非法状态转换）
- [x] 订单金额计算和验证
- [x] 订单项快照存储（商品信息快照）
- [x] 用户权限验证（只能查看自己的订单）
- [x] 库存验证和扣减（创建订单时检查库存，使用数据库原子操作扣减库存，防止并发问题）
- [x] 所有 API 测试通过

#### ✅ 订单管理功能（前端）

- [x] 订单管理 API 接口（frontend/src/api/order.ts）
- [x] 订单管理页面（frontend/src/views/system/order/index.vue）
- [x] 订单详情对话框（frontend/src/views/system/order/components/OrderDetailDialog.vue）
- [x] 订单状态更新对话框（frontend/src/views/system/order/components/OrderStatusDialog.vue）
- [x] 路由配置（/system/order）
- [x] 订单列表展示（分页、搜索、状态筛选）
- [x] 订单状态管理（支付、取消等操作）
- [x] 价格格式化显示
- [x] 所有功能测试通过

#### ✅ 商品分类管理功能（后端）

- [x] 分类模型定义（backend/model/category，支持树形结构）
- [x] 分类 API 定义（user.api 中添加分类相关类型和路由）
- [x] 分类错误码定义（8000-8999）
- [x] 分类 CRUD API（创建、列表、详情、更新、删除）
- [x] 分类列表分页和搜索功能（支持全部模式和分页模式）
- [x] 分类名称同级唯一性检查
- [x] 循环引用检查（防止将父分类设置为自己的子分类）
- [x] 删除分类时子分类和商品关联检查
- [x] 商品模型添加分类关联（CategoryID）
- [x] 路由权限配置（category:list、category:create、category:update、category:delete）
- [x] 所有 API 测试通过

#### ✅ 商品分类管理功能（前端）

- [x] 分类管理 API 接口（frontend/src/api/category.ts）
- [x] 分类管理页面（frontend/src/views/system/category/index.vue）
- [x] 分类新增/编辑对话框（frontend/src/views/system/category/components/CategoryDialog.vue）
- [x] 路由配置（/system/category）
- [x] 表单验证（名称、父分类、排序、状态）
- [x] 父分类下拉选择（支持搜索和过滤）
- [x] 分类列表展示（分页、搜索）
- [x] 所有功能测试通过

#### ✅ 购物车管理功能（后端）

- [x] 购物车模型定义（backend/model/cart、backend/model/cart_item）
- [x] 购物车 API 定义（user.api 中添加购物车相关类型和路由）
- [x] 购物车错误码定义（9000-9999）
- [x] 购物车 CRUD API（添加商品、获取购物车、更新数量、删除商品、清空购物车）
- [x] 参数验证（数量验证、商品状态检查）
- [x] 性能优化（解决 N+1 查询问题，批量查询商品）
- [x] 路由权限配置（cart:get、cart:add、cart:update、cart:delete）
- [x] 所有 API 测试通过

#### ✅ 购物车管理功能（前端）

- [x] 购物车管理 API 接口（frontend/src/api/cart.ts）
- [x] 购物车管理页面（frontend/src/views/system/cart/index.vue）
- [x] 路由配置（/system/cart）
- [x] 购物车信息展示（商品种类数、总金额）
- [x] 数量管理功能（加、减按钮，输入框修改）
- [x] 删除购物车项功能
- [x] 清空购物车功能
- [x] 商品列表页面"加入购物车"功能（带数量选择对话框）
- [x] 所有功能测试通过

#### ✅ 数据统计 Dashboard 功能（后端+前端）

- [x] Dashboard API 定义（dashboard.api）
- [x] 统计逻辑实现（订单统计、商品统计、用户统计）
- [x] 路由配置（需要认证）
- [x] 前端 API 接口（dashboard.ts）
- [x] Dashboard 页面（统计卡片、详细统计、ECharts 图表）
- [x] ECharts 图表展示（订单状态分布、商品状态分布）
- [x] 响应式设计（支持移动端和桌面端）
- [x] 所有功能测试通过

#### ✅ Redis 缓存集成功能（后端）

- [x] Redis 连接模块（backend/common/redis/redis.go）
- [x] Redis 配置集成（配置文件和服务上下文）
- [x] 商品列表缓存实现（Cache-Aside 模式）
- [x] 随机过期时间优化（防止缓存雪崩，300-360秒随机）
- [x] 缓存键规范化（MD5 哈希处理关键词，防止特殊字符问题）
- [x] 缓存清除机制（商品创建/更新/删除时自动清除相关缓存）
- [x] 空结果缓存（防止缓存穿透）
- [x] 所有功能测试通过

#### ✅ 定时任务功能（后端）

- [x] Cron 管理模块（backend/common/cron/cron.go）
- [x] 订单自动取消任务（OrderCancelJob）
- [x] 超时未支付订单自动取消（默认30分钟超时，测试环境1分钟）
- [x] 事务处理确保数据一致性（订单状态更新+库存恢复）
- [x] 详细的 cron 表达式注释
- [x] 数据库查询优化（只查询超时订单）
- [x] 任务完成等待机制（优雅关闭）
- [x] 定时任务测试脚本（scripts/test-order-cancel-job.sh）
- [x] 所有功能测试通过

#### ✅ 用户 RPC 服务（后端）

- [x] RPC 服务创建（backend/service/user/user-rpc）
- [x] Proto 文件定义（user-rpc.proto）
- [x] RPC 服务配置（数据库连接、端口 8081）
- [x] 用户查询 RPC（GetUser、ListUsers 含 keyword 搜索）
- [x] 用户创建 RPC（CreateUser，包含密码加密、唯一性检查）
- [x] 用户更新 RPC（UpdateUser，支持邮箱和密码更新）
- [x] 用户删除 RPC（DeleteUser）
- [x] gRPC 错误处理（使用标准错误码：InvalidArgument、NotFound、AlreadyExists）
- [x] API 服务集成 RPC 客户端（ServiceContext 中添加 UserRpc）
- [x] 所有用户接口迁移到 RPC（GetUserDetail、GetUserList、Register、UpdateUser、UpdateUserDetail、DeleteUser）
- [x] 错误码映射（gRPC 错误码 → 业务错误码）
- [x] 所有功能测试通过

#### ✅ 商品 RPC 服务（后端）

- [x] RPC 服务创建（backend/service/product/product-rpc）
- [x] Proto 文件定义（product-rpc.proto）
- [x] RPC 服务配置（数据库连接、端口 8082）
- [x] 商品查询 RPC（GetProduct、ListProducts 含 keyword 搜索和分页）
- [x] 商品创建 RPC（CreateProduct，包含名称唯一性检查、价格验证）
- [x] 商品更新 RPC（UpdateProduct，支持部分字段更新、图片更新）
- [x] 商品状态更新 RPC（UpdateProductStatus）
- [x] 商品删除 RPC（DeleteProduct，支持硬删除和软删除）
- [x] gRPC 错误处理（使用标准错误码：InvalidArgument、NotFound、AlreadyExists）
- [x] API 服务集成 RPC 客户端（ServiceContext 中添加 ProductRpc）
- [x] 所有商品接口迁移到 RPC（GetProductDetail、GetProductList、CreateProduct、UpdateProduct、UpdateProductStatus、DeleteProduct）
- [x] 错误码映射（gRPC 错误码 → 业务错误码）
- [x] 缓存清除机制（商品创建/更新/删除时自动清除相关缓存）
- [x] 所有功能测试通过

### 待完成功能

#### 阶段三：商品管理扩展

- [x] 商品分类管理 ✅
- [x] 商品库存管理 ✅
- [x] 商品图片上传 ✅

#### 阶段四：订单管理扩展

- [x] 订单创建 ✅
- [x] 订单状态管理 ✅
- [x] 购物车功能 ✅

#### 阶段五：高级功能

- [x] 文件上传下载 ✅
- [x] 数据统计 Dashboard ✅
- [x] Redis 缓存集成 ✅
- [x] 定时任务 ✅

#### 阶段六：RPC 服务

- [x] 用户 RPC 服务 ✅
  - [x] RPC 服务创建和配置 ✅
  - [x] 用户查询接口（GetUser、ListUsers）✅
  - [x] 用户创建接口（CreateUser）✅
  - [x] 用户更新接口（UpdateUser）✅
  - [x] 用户删除接口（DeleteUser）✅
  - [x] API 服务集成 RPC 客户端 ✅
  - [x] 所有用户接口迁移到 RPC ✅
- [x] 商品 RPC 服务 ✅
  - [x] RPC 服务创建和配置 ✅
  - [x] 商品查询接口（GetProduct、ListProducts）✅
  - [x] 商品创建接口（CreateProduct）✅
  - [x] 商品更新接口（UpdateProduct、UpdateProductStatus）✅
  - [x] 商品删除接口（DeleteProduct）✅
  - [x] API 服务集成 RPC 客户端 ✅
  - [x] 所有商品接口迁移到 RPC ✅
  - [x] 缓存清除机制集成 ✅
- [ ] 订单 RPC 服务

#### 阶段七：优化和部署

- [x] 错误处理完善 ✅
- [ ] 日志系统
- [ ] Docker 部署
- [ ] 性能优化

### 测试脚本

测试脚本统一存放在 `scripts/` 目录下，使用 kebab-case 命名规范。

- **订单自动取消测试**：`scripts/test-order-cancel-job.sh`
  - 测试定时任务的订单自动取消功能
  - 验证超时未支付订单自动取消、库存恢复
  - 使用方法：`./scripts/test-order-cancel-job.sh`

- **权限验证测试**：`scripts/test-permission-middleware.sh`
  - 测试权限中间件的功能
  - 覆盖未登录访问、无角色访问、有权限访问、无权限访问等场景
  - 使用方法：`./scripts/test-permission-middleware.sh`

详见：`scripts/README.md`

### 当前问题/注意事项

#### 常规配置
- 配置文件字段名：使用 `dataSource`（小写驼峰）
- **API 服务运行方式**：`cd backend/service/user/api && go run user-api.go -f etc/user-api.yaml`（端口 8888）
- **用户 RPC 服务运行方式**：`cd backend/service/user/user-rpc && go run userrpc.go -f etc/userrpc.yaml`（端口 8081）
- **商品 RPC 服务运行方式**：`cd backend/service/product/product-rpc && go run productrpc.go -f etc/productrpc.yaml`（端口 8082）
- **注意**：API 服务依赖 RPC 服务，需要先启动 RPC 服务，再启动 API 服务
- 数据库：MySQL 3307 端口，数据库名 testdb
- Redis 配置：默认使用 `127.0.0.1:6379`，使用 Docker 启动：`docker run -d --name redis-dev -p 6379:6379 redis:7-alpine`
- go-zero 参数验证：可选字段（optional）在 JSON 中缺失时会报错，需要在请求中包含所有字段（临时方案）
- 网络配置：如果系统无法解析 `localhost`，配置文件已使用 `127.0.0.1` 替代
- 测试脚本：所有测试脚本统一存放在 `scripts/` 目录，使用 kebab-case 命名规范

#### ⚠️ 二次提示问题（AI 需要避免）
**问题描述**：在前端错误处理中，如果 API 返回错误（如权限验证失败），会出现两个错误提示：
1. 来自服务器的原始错误信息（如 "没有权限访问"）
2. 前端 catch 块中捕获异常后的重复提示（如 "获取商品列表失败"）

**解决方案**：
- 在 catch 块中**不应该再显示错误提示**（已通过 ElMessage.error 捕获了 API 返回的错误）
- 只在 catch 块中输出日志用于调试
- 让用户只看到来自服务器的原始、清晰的错误信息

**示例代码修改**：
```typescript
const fetchProductList = async () => {
  try {
    loading.value = true
    const response = await getProductList(params)

    if (response.code === 0) {
      // 成功处理
    } else {
      ElMessage.error(response.message || '操作失败')  // ✅ 显示服务器错误
    }
  } catch (error) {
    console.error('错误:', error)  // ✅ 只在控制台输出
    // ❌ 不要再 ElMessage.error('操作失败')
  } finally {
    loading.value = false
  }
}
```

**最后出现**：2025-01-22 商品管理功能修复

**AI 提醒**：下次遇到类似 API 错误处理时，记得检查是否有重复的错误提示

### 下一步计划

**当前阶段**：阶段六 - RPC 服务（进行中）

**完成情况**：
- 阶段一（用户认证和管理）全部完成 ✅
- 阶段二（权限管理）全部完成 ✅
- 阶段三（商品管理）全部完成 ✅
- 阶段四（订单管理）全部完成 ✅
- 阶段五（高级功能）全部完成 ✅
- 阶段六（RPC 服务）部分完成：
  - 用户 RPC 服务 ✅
  - 商品 RPC 服务 ✅
  - API 服务调用 RPC 服务 ✅

**下一步选择**：
1. **继续 RPC 迁移**：商品 RPC 服务、订单 RPC 服务
2. **优化和部署**：日志系统、Docker 部署、性能优化
3. **服务发现**：集成 Etcd 实现服务注册与发现

**教学方法论已总结**：基于商品管理和 RPC 迁移的成功实践，已将核心教学策略整理到文档中

**最后更新**：2025-12-04（完成商品 RPC 服务迁移）
**当前状态**：

- 阶段一（用户认证和管理）全部完成 ✅
- 阶段二（权限管理）全部完成 ✅：
  - 角色管理（后端+前端）✅
  - 权限管理（后端+前端）✅
  - 菜单管理（后端+前端）✅
  - 用户角色管理（后端+前端）✅
  - 角色权限管理（后端+前端）✅
  - 权限中间件（权限验证）✅
- 阶段三（商品管理）全部完成 ✅：
  - 商品 CRUD（后端+前端）✅
  - 商品列表分页搜索✅
  - UI bug 修复（重复提示）✅
  - 商品分类管理（后端+前端）✅
  - 商品库存管理（后端+前端）✅
  - 商品图片上传（后端+前端）✅
- 阶段四（订单管理）全部完成 ✅：
  - 订单创建（后端+前端）✅
  - 订单列表（后端+前端）✅
  - 订单详情（后端+前端）✅
  - 订单状态管理（后端+前端）✅
  - 购物车功能（后端+前端）✅
  - 订单库存验证和扣减（后端）✅
- 阶段五（高级功能）进行中：
  - 文件上传下载（后端+前端）✅
    - 文件上传 API（支持分类、大小限制、类型验证）✅
    - 静态文件服务（使用 go-zero 的 rest.WithFileServer）✅
    - 前端文件上传组件（支持拖拽、预览、进度显示）✅
    - 商品图片上传集成（多图片上传、预览、删除）✅
  - 数据统计 Dashboard（后端+前端）✅
    - 订单统计（总数、今日订单、总金额、今日销售额、状态分布）✅
    - 商品统计（总数、上架/下架、库存预警、总库存）✅
    - 用户统计（总数、今日新增、活跃用户）✅
    - ECharts 图表展示（订单状态分布、商品状态分布）✅
  - Redis 缓存集成（后端）✅
    - Redis 连接模块和配置集成✅
    - 商品列表缓存实现（Cache-Aside 模式）✅
    - 随机过期时间优化（防止缓存雪崩）✅
    - 缓存键规范化（MD5 哈希处理）✅
    - 缓存清除机制（数据更新时自动清除）✅
    - 空结果缓存（防止缓存穿透）✅
  - 定时任务（后端）✅
    - Cron 管理模块（common/cron）✅
    - 订单自动取消任务（超时未支付订单自动取消）✅
    - 事务处理确保数据一致性（订单状态更新+库存恢复）✅
    - 详细的 cron 表达式注释✅
    - 定时任务测试脚本（scripts/test-order-cancel-job.sh）✅

---

## 🚀 RPC 服务架构与迁移经验

### 为什么需要 RPC？

在微服务架构中，RPC 服务用于：

1. **领域边界清晰**：将业务逻辑按领域拆分（用户、商品、订单等）
2. **服务复用**：多个 API 服务可以共享同一个 RPC 服务的能力
3. **独立扩展**：可以单独扩容某个领域的服务，而不影响其他服务
4. **技术异构**：不同领域可以使用不同的技术栈或数据库
5. **安全隔离**：数据库只对 RPC 服务暴露，其他服务通过 RPC 访问

### 当前架构

```
前端 → API 服务（8888） → UserRpc 服务（8081） → MySQL
       │                    │
       │                    └─ 用户领域业务逻辑
       │
       └─ ProductRpc 服务（8082） → MySQL
                      │
                      └─ 商品领域业务逻辑
       └─ HTTP 协议、认证、权限、错误映射
```

### 已迁移的用户接口

| API 接口 | HTTP 方法 | 路径 | RPC 方法 | 状态 |
|---------|----------|------|---------|------|
| 用户详情 | GET | `/api/users/:id` | `GetUser` | ✅ |
| 用户列表 | GET | `/api/users` | `ListUsers` | ✅ |
| 创建用户 | POST | `/api/users` | `CreateUser` | ✅ |
| 更新当前用户 | PUT | `/api/users/me` | `UpdateUser` | ✅ |
| 更新指定用户 | PUT | `/api/users/:id` | `UpdateUser`（复用） | ✅ |
| 删除用户 | DELETE | `/api/users/:id` | `DeleteUser` | ✅ |

### 已迁移的商品接口

| API 接口 | HTTP 方法 | 路径 | RPC 方法 | 状态 |
|---------|----------|------|---------|------|
| 商品详情 | GET | `/api/products/:id` | `GetProduct` | ✅ |
| 商品列表 | GET | `/api/products` | `ListProducts` | ✅ |
| 创建商品 | POST | `/api/products` | `CreateProduct` | ✅ |
| 更新商品 | PUT | `/api/products/:id` | `UpdateProduct` | ✅ |
| 更新商品状态 | PUT | `/api/products/:id/status` | `UpdateProductStatus` | ✅ |
| 删除商品 | DELETE | `/api/products/:id` | `DeleteProduct` | ✅ |

### RPC 服务开发步骤

#### 1. 创建 RPC 服务

```bash
cd backend/service/user
goctl rpc new user-rpc --module go-zero-learning
```

#### 2. 定义 Proto 文件

```protobuf
syntax = "proto3";

package userrpc;
option go_package="./userrpc";

message GetUserReq {
  int64 id = 1;
}

message GetUserResp {
  int64 id = 1;
  string username = 2;
  string email = 3;
}

service UserRpc {
  rpc GetUser(GetUserReq) returns (GetUserResp);
}
```

**注意事项**：
- `package` 和 `service` 名称不能包含连字符（`-`），使用驼峰命名
- `go_package` 指定生成的 Go 包路径

#### 3. 生成代码

```bash
cd backend/service/user/user-rpc
goctl rpc protoc user-rpc.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=gozero
```

#### 4. 实现业务逻辑

在 `internal/logic/` 中实现 RPC 方法：

```go
func (l *GetUserLogic) GetUser(in *userrpc.GetUserReq) (*userrpc.GetUserResp, error) {
    // 业务逻辑
    // 使用 gRPC 标准错误码
    if in.Id <= 0 {
        return nil, status.Error(codes.InvalidArgument, "用户ID必须大于0")
    }
    // ...
}
```

#### 5. 在 API 服务中集成 RPC 客户端

**配置文件中添加 RPC 客户端配置**（`etc/user-api.yaml`）：

```yaml
UserRpc:
  Endpoints:
    - 127.0.0.1:8081
  Timeout: 3000
```

**在 ServiceContext 中创建客户端**：

```go
// internal/config/config.go
type Config struct {
    // ...
    UserRpc zrpc.RpcClientConf `json:"userRpc"`
}

// internal/svc/servicecontext.go
func NewServiceContext(c config.Config) *ServiceContext {
    // ...
    userRpcClient := zrpc.MustNewClient(c.UserRpc)
    userRpc := userrpcclient.NewUserRpc(userRpcClient)
    
    return &ServiceContext{
        // ...
        UserRpc: userRpc,
    }
}
```

#### 6. 在 API Logic 中调用 RPC

```go
func (l *GetUserDetailLogic) GetUserDetail(req *types.GetUserDetailReq) (*types.UserInfoResp, error) {
    // 调用 RPC
    rpcResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userrpc.GetUserReq{
        Id: req.ID,
    })
    if err != nil {
        // gRPC 错误码映射到业务错误
        if st, ok := status.FromError(err); ok {
            switch st.Code() {
            case codes.NotFound:
                return nil, errorx.ErrUserNotFound
            // ...
            }
        }
        return nil, errorx.ErrInternalError
    }
    
    // 转换为 API 响应格式
    return &types.UserInfoResp{
        ID:       rpcResp.Id,
        Username: rpcResp.Username,
        Email:    rpcResp.Email,
    }, nil
}
```

### 错误处理模式

**RPC 层**：使用 gRPC 标准错误码
- `codes.InvalidArgument`：参数错误
- `codes.NotFound`：资源不存在
- `codes.AlreadyExists`：资源已存在
- `codes.Internal`：内部错误

**API 层**：将 gRPC 错误码映射到业务错误码
- `codes.NotFound` → `errorx.ErrUserNotFound`
- `codes.AlreadyExists` → `errorx.ErrEmailExists`
- `codes.InvalidArgument` → `errorx.ErrInvalidParam`

### 职责划分原则

**RPC 层负责**：
- 领域业务逻辑（用户创建、更新、删除等）
- 数据验证（唯一性检查、格式验证等）
- 数据访问（数据库操作）
- 业务规则（密码加密、状态流转等）

**API 层负责**：
- HTTP 协议处理（请求解析、响应格式化）
- 认证和授权（JWT 验证、权限检查）
- 错误映射（gRPC 错误 → 业务错误）
- 数据组合（调用多个 RPC 组合数据）

### 迁移经验总结

1. **渐进式迁移**：先迁移查询接口，再迁移写操作接口
2. **保留旧代码**：迁移时保留旧逻辑作为备份，方便对比和回滚
3. **统一错误处理**：RPC 层使用 gRPC 标准错误码，API 层统一映射
4. **职责清晰**：RPC 只关注领域逻辑，API 只关注协议和权限
5. **复用 RPC 方法**：`UpdateUser` 既可以更新当前用户，也可以更新指定用户（通过传入不同的 ID）

### 运行方式

**启动 RPC 服务**：
```bash
cd backend/service/user/user-rpc
go run userrpc.go -f etc/userrpc.yaml
```

**启动 API 服务**：
```bash
cd backend/service/user/api
go run user-api.go -f etc/user-api.yaml
```

**注意**：API 服务依赖 RPC 服务，需要先启动 RPC 服务。

---

## 📚 最小必要知识

### 1. go-zero 项目结构

```text
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

```text
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

### 4. RPC 服务开发流程

```text
go-zero RPC：
1. 写 .proto 文件：定义 RPC 服务和方法
2. 运行 goctl rpc protoc：生成 server、logic 骨架
3. 在 logic 中实现业务逻辑
4. 在 API 服务的 ServiceContext 中创建 RPC 客户端
5. 在 API logic 中调用 RPC 方法
```

### 5. RPC 服务架构

```text
前端 → API 服务（HTTP） → RPC 服务（gRPC） → MySQL
       │                    │
       │                    └─ 领域业务逻辑
       └─ HTTP 协议、认证、权限、错误映射
```

**职责划分**：
- **API 层**：HTTP 协议处理、认证、权限、错误映射、组合数据
- **RPC 层**：领域业务逻辑、数据访问、业务规则验证

---

## 🔗 学习资源

1. **go-zero 官方文档**：<https://go-zero.dev/>
2. **GORM 文档**：<https://gorm.io/>
3. **go-zero-looklook**：<https://github.com/Mikaelemmmm/go-zero-looklook>
