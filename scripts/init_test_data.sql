-- ============================================
-- 测试数据初始化脚本
-- 用途：初始化系统测试数据，包括用户、角色、权限、菜单等
-- 使用方法：在数据库中执行此脚本
-- 注意：此脚本会先清理旧数据，然后插入新数据
-- ============================================

-- 设置时区
SET time_zone = '+08:00';

-- ============================================
-- 清理旧数据（可选，如果不想清理请注释掉这部分）
-- ============================================
-- 注意：删除顺序很重要，需要先删除关联表的数据
-- 临时禁用安全更新模式以允许删除所有数据

SET SQL_SAFE_UPDATES = 0;

DELETE FROM `role_permissions`;
DELETE FROM `user_roles`;
DELETE FROM `menus`;
DELETE FROM `permissions`;
DELETE FROM `roles`;
DELETE FROM `users`;

-- 重新启用安全更新模式
SET SQL_SAFE_UPDATES = 1;

-- 重置自增ID（可选，如果需要从1开始）
ALTER TABLE `users` AUTO_INCREMENT = 1;
ALTER TABLE `roles` AUTO_INCREMENT = 1;
ALTER TABLE `permissions` AUTO_INCREMENT = 1;
ALTER TABLE `menus` AUTO_INCREMENT = 1;
ALTER TABLE `user_roles` AUTO_INCREMENT = 1;
ALTER TABLE `role_permissions` AUTO_INCREMENT = 1;

-- ============================================
-- 1. 用户数据
-- ============================================
-- 注意：密码使用 bcrypt 加密存储
-- admin 用户：密码 123456
-- testuser 用户：密码 123456（测试脚本需要）
-- normaluser 用户：密码 Normal123（测试脚本需要）

INSERT INTO `users` (`id`, `username`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, 'admin', 'admin@example.com', '$2a$10$DawBfZ2nRXPLJ7vx66p9oOB4HQzZ7USBijuK8QxV0T..xdePKB61.', NOW(), NOW()),
(2, 'testuser', 'testuser@example.com', '$2a$10$DawBfZ2nRXPLJ7vx66p9oOB4HQzZ7USBijuK8QxV0T..xdePKB61.', NOW(), NOW()),
(3, 'normaluser', 'normaluser@example.com', '$2a$10$/0OZVVaxD2H9g5zMSC.fduY4deDfs4kx8PLhuF3L4jB.FceXXBZ/W', NOW(), NOW());

-- ============================================
-- 2. 角色数据
-- ============================================

INSERT INTO `roles` (`id`, `name`, `code`, `desc`, `created_at`, `updated_at`) VALUES
(1, '管理员', 'admin', '系统管理员，拥有所有权限', NOW(), NOW()),
(2, '普通用户', 'user', '普通用户，拥有基本权限', NOW(), NOW());

-- ============================================
-- 3. 权限数据
-- ============================================

-- 用户管理权限
INSERT INTO `permissions` (`id`, `name`, `code`, `desc`, `created_at`, `updated_at`) VALUES
(1, '查看用户列表', 'user:list', '查看用户列表和用户详情的权限', NOW(), NOW()),
(2, '创建用户', 'user:create', '创建新用户的权限', NOW(), NOW()),
(3, '更新用户', 'user:update', '更新用户信息的权限', NOW(), NOW()),
(4, '删除用户', 'user:delete', '删除用户的权限', NOW(), NOW());

-- 角色管理权限
INSERT INTO `permissions` (`id`, `name`, `code`, `desc`, `created_at`, `updated_at`) VALUES
(5, '查看角色列表', 'role:list', '查看角色列表和角色详情的权限', NOW(), NOW()),
(6, '创建角色', 'role:create', '创建新角色的权限', NOW(), NOW()),
(7, '更新角色', 'role:update', '更新角色信息和分配用户角色的权限', NOW(), NOW()),
(8, '删除角色', 'role:delete', '删除角色的权限', NOW(), NOW());

-- 权限管理权限
INSERT INTO `permissions` (`id`, `name`, `code`, `desc`, `created_at`, `updated_at`) VALUES
(9, '查看权限列表', 'permission:list', '查看权限列表、权限详情、菜单列表和角色权限的权限', NOW(), NOW()),
(10, '创建权限', 'permission:create', '创建新权限和菜单的权限', NOW(), NOW()),
(11, '更新权限', 'permission:update', '更新权限信息、菜单信息和分配角色权限的权限', NOW(), NOW()),
(12, '删除权限', 'permission:delete', '删除权限和菜单的权限', NOW(), NOW());

-- ============================================
-- 4. 菜单数据
-- ============================================

-- 顶级菜单
INSERT INTO `menus` (`id`, `name`, `code`, `desc`, `parent_id`, `path`, `icon`, `type`, `sort`, `status`, `created_at`, `updated_at`) VALUES
(1, '系统管理', 'system', '系统管理模块', 0, '/system', 'Setting', 1, 1, 1, NOW(), NOW()),
(2, '仪表盘', 'dashboard', '系统仪表盘', 0, '/dashboard', 'Dashboard', 1, 0, 1, NOW(), NOW());

-- 系统管理子菜单
INSERT INTO `menus` (`id`, `name`, `code`, `desc`, `parent_id`, `path`, `icon`, `type`, `sort`, `status`, `created_at`, `updated_at`) VALUES
(3, '用户管理', 'system:user', '用户管理', 1, '/system/user', 'User', 1, 1, 1, NOW(), NOW()),
(4, '角色管理', 'system:role', '角色管理', 1, '/system/role', 'UserGroup', 1, 2, 1, NOW(), NOW()),
(5, '权限管理', 'system:permission', '权限管理', 1, '/system/permission', 'Lock', 1, 3, 1, NOW(), NOW()),
(6, '菜单管理', 'system:menu', '菜单管理', 1, '/system/menu', 'Menu', 1, 4, 1, NOW(), NOW()),
(7, '系统设置', 'system:settings', '系统设置', 1, '/system/settings', 'Setting', 1, 5, 1, NOW(), NOW());

-- ============================================
-- 5. 用户角色关联
-- ============================================

-- admin 用户分配管理员角色
INSERT INTO `user_roles` (`id`, `user_id`, `role_id`, `created_at`, `updated_at`) VALUES
(1, 1, 1, NOW(), NOW());

-- testuser 用户分配管理员角色（测试脚本需要）
INSERT INTO `user_roles` (`id`, `user_id`, `role_id`, `created_at`, `updated_at`) VALUES
(2, 2, 1, NOW(), NOW());

-- normaluser 用户不分配角色（测试脚本需要，用于测试无角色用户）

-- ============================================
-- 6. 角色权限关联
-- ============================================

-- 管理员角色拥有所有权限
INSERT INTO `role_permissions` (`id`, `role_id`, `permission_id`, `created_at`, `updated_at`) VALUES
(1, 1, 1, NOW(), NOW()),  -- user:list
(2, 1, 2, NOW(), NOW()),  -- user:create
(3, 1, 3, NOW(), NOW()),  -- user:update
(4, 1, 4, NOW(), NOW()),  -- user:delete
(5, 1, 5, NOW(), NOW()),  -- role:list
(6, 1, 6, NOW(), NOW()),  -- role:create
(7, 1, 7, NOW(), NOW()),  -- role:update
(8, 1, 8, NOW(), NOW()),  -- role:delete
(9, 1, 9, NOW(), NOW()),  -- permission:list
(10, 1, 10, NOW(), NOW()), -- permission:create
(11, 1, 11, NOW(), NOW()), -- permission:update
(12, 1, 12, NOW(), NOW()); -- permission:delete

-- 普通用户角色拥有基本权限（可根据需要调整）
INSERT INTO `role_permissions` (`id`, `role_id`, `permission_id`, `created_at`, `updated_at`) VALUES
(13, 2, 1, NOW(), NOW()); -- user:list（仅查看）

-- ============================================
-- 说明
-- ============================================
-- 1. 管理员账号：admin / 123456
-- 2. 测试账号：testuser / 123456（拥有管理员权限，测试脚本需要）
-- 3. 普通用户：normaluser / Normal123（无角色，测试脚本需要）
-- 4. 所有密码已使用 bcrypt 加密
-- 5. 测试脚本 test-permission-middleware.sh 会使用 testuser 和 normaluser 进行测试

