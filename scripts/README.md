# 测试脚本目录

本目录用于存放各种测试脚本，包括 API 测试、功能测试等。

## 脚本列表

### `test-order-cancel-job.sh`
订单自动取消定时任务测试脚本

**用途：** 测试订单自动取消功能，验证定时任务是否正常取消超时未支付的订单

**使用方法：**
```bash
./scripts/test-order-cancel-job.sh
```

**测试场景：**
1. 登录获取 Token
2. 查询商品列表获取商品ID
3. 创建待支付订单
4. 等待定时任务执行（3.5分钟）
5. 查询订单状态是否已被自动取消
6. 检查库存是否已恢复

**前置条件：**
- 服务正在运行（默认端口 8888）
- 数据库中有测试用户 `admin`（密码：123456）
- 数据库中有可用商品
- 定时任务已启动（订单超时时间为1分钟，定时任务每分钟执行一次）

---

### `test-permission-middleware.sh`
权限验证中间件测试脚本

**用途：** 测试权限中间件的功能，验证权限验证是否正常工作

**使用方法：**
```bash
./scripts/test-permission-middleware.sh
```

**测试场景：**
1. 未登录用户访问受保护路由 → 应返回 401
2. 无角色用户访问受保护路由 → 应返回 403
3. 有权限用户访问 → 应返回 200
4. 无权限用户访问 → 应返回 403
5. 不同权限要求的路由验证
6. 不需要权限的路由（只需要登录）

**前置条件：**
- 服务正在运行（默认端口 8888）
- 数据库中有测试用户 `testuser`（管理员用户）
- 测试用户 `normaluser`（如果没有会自动创建）

---

### `init_test_data.sql`
数据库测试数据初始化脚本

**用途：** 初始化系统测试数据，包括用户、角色、权限、菜单等基础数据

**使用方法：**
```bash
# 方式1：使用 MySQL 命令行
mysql -u your_username -p your_database < scripts/init_test_data.sql

# 方式2：在 MySQL 客户端中执行
source scripts/init_test_data.sql;
```

**包含的测试数据：**
1. **用户数据：**
   - `admin` / `123456` - 管理员账号（拥有所有权限）
   - `testuser` / `123456` - 测试账号（拥有管理员权限，用于测试脚本）
   - `normaluser` / `Normal123` - 普通用户（无角色，用于测试脚本）

2. **角色数据：**
   - 管理员角色（拥有所有权限）
   - 普通用户角色（仅拥有查看权限）

3. **权限数据：**
   - 用户管理权限（user:list, user:create, user:update, user:delete）
   - 角色管理权限（role:list, role:create, role:update, role:delete）
   - 权限管理权限（permission:list, permission:create, permission:update, permission:delete）

4. **菜单数据：**
   - 仪表盘菜单
   - 系统管理菜单及其子菜单（用户管理、角色管理、权限管理、菜单管理、系统设置）

**注意事项：**
- 脚本会先清理旧数据，然后插入新数据
- 如果不想清理旧数据，请注释掉脚本开头的 DELETE 语句
- 所有密码已使用 bcrypt 加密存储
- 执行脚本前请确保数据库表结构已创建

---

## 命名规范

测试脚本统一使用以下命名规范：
- 使用 kebab-case（短横线分隔）
- 格式：`test-{功能名称}.sh`
- 示例：`test-permission-middleware.sh`、`test-user-api.sh`

## 添加新测试脚本

添加新的测试脚本时，请：
1. 使用统一的命名规范
2. 在脚本头部添加详细的注释说明
3. 更新本 README 文件，添加脚本说明

