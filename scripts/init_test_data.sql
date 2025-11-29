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

-- 【扩展】删除顺序：先删除关联表，再删除主表
DELETE FROM `order_items`;
DELETE FROM `orders`;
DELETE FROM `products`;
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
ALTER TABLE `products` AUTO_INCREMENT = 1;
ALTER TABLE `orders` AUTO_INCREMENT = 1;
ALTER TABLE `order_items` AUTO_INCREMENT = 1;

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

-- 【扩展】商品管理权限
INSERT INTO `permissions` (`id`, `name`, `code`, `desc`, `created_at`, `updated_at`) VALUES
(13, '查看商品列表', 'product:list', '查看商品列表和商品详情的权限', NOW(), NOW()),
(14, '创建商品', 'product:create', '创建新商品的权限', NOW(), NOW()),
(15, '更新商品', 'product:update', '更新商品信息的权限', NOW(), NOW()),
(16, '删除商品', 'product:delete', '删除商品的权限', NOW(), NOW());

-- 【扩展】订单管理权限
INSERT INTO `permissions` (`id`, `name`, `code`, `desc`, `created_at`, `updated_at`) VALUES
(17, '查看订单列表', 'order:list', '查看订单列表和订单详情的权限', NOW(), NOW()),
(18, '创建订单', 'order:create', '创建新订单的权限', NOW(), NOW()),
(19, '更新订单', 'order:update', '更新订单状态的权限', NOW(), NOW());

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
(7, '商品管理', 'system:product', '商品管理', 1, '/system/product', 'Goods', 1, 6, 1, NOW(), NOW()),
(8, '订单管理', 'system:order', '订单管理', 1, '/system/order', 'ShoppingCart', 1, 7, 1, NOW(), NOW()),
(9, '系统设置', 'system:settings', '系统设置', 1, '/system/settings', 'Setting', 1, 8, 1, NOW(), NOW());

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
(12, 1, 12, NOW(), NOW()), -- permission:delete
-- 【扩展】商品管理权限
(13, 1, 13, NOW(), NOW()), -- product:list
(14, 1, 14, NOW(), NOW()), -- product:create
(15, 1, 15, NOW(), NOW()), -- product:update
(16, 1, 16, NOW(), NOW()), -- product:delete
-- 【扩展】订单管理权限
(17, 1, 17, NOW(), NOW()), -- order:list
(18, 1, 18, NOW(), NOW()), -- order:create
(19, 1, 19, NOW(), NOW()); -- order:update

-- 普通用户角色拥有基本权限（可根据需要调整）
INSERT INTO `role_permissions` (`id`, `role_id`, `permission_id`, `created_at`, `updated_at`) VALUES
(20, 2, 1, NOW(), NOW()); -- user:list（仅查看）

-- ============================================
-- 【扩展】7. 商品测试数据
-- ============================================
-- 注意：价格以分为单位存储（1元 = 100分）

INSERT INTO `products` (`id`, `name`, `description`, `price`, `status`, `created_at`, `updated_at`) VALUES
(1, 'iPhone 15 Pro', '苹果最新款手机，配备A17 Pro芯片，支持5G网络', 799900, 1, NOW(), NOW()),
(2, 'MacBook Pro 14寸', 'M3芯片，14英寸 Liquid Retina XDR 显示屏', 1499900, 1, NOW(), NOW()),
(3, 'AirPods Pro', '主动降噪无线耳机，支持空间音频', 189900, 1, NOW(), NOW()),
(4, 'iPad Air', 'M2芯片，10.9英寸 Liquid Retina 显示屏', 439900, 1, NOW(), NOW()),
(5, 'Apple Watch Series 9', '智能手表，支持健康监测和运动追踪', 299900, 1, NOW(), NOW()),
(6, 'Magic Keyboard', '适用于 iPad 的键盘，带触控板', 239900, 1, NOW(), NOW()),
(7, 'AirTag 4件装', '物品追踪器，帮助找回丢失的物品', 99000, 1, NOW(), NOW()),
(8, '下架商品示例', '这是一个已下架的商品示例', 99900, 0, NOW(), NOW());

-- ============================================
-- 【扩展】8. 订单测试数据
-- ============================================
-- 注意：订单金额以分为单位存储

-- 订单1：admin用户的订单（待支付）
INSERT INTO `orders` (`id`, `order_no`, `user_id`, `total_amount`, `status`, `remark`, `created_at`, `updated_at`) VALUES
(1, 'ORD20241129000001', 1, 799900, 1, '测试订单1', NOW(), NOW());

-- 订单1的订单项
INSERT INTO `order_items` (`id`, `order_id`, `product_id`, `product_name`, `product_desc`, `price`, `quantity`, `amount`, `created_at`) VALUES
(1, 1, 1, 'iPhone 15 Pro', '苹果最新款手机，配备A17 Pro芯片，支持5G网络', 799900, 1, 799900, NOW());

-- 订单2：admin用户的订单（已支付）
INSERT INTO `orders` (`id`, `order_no`, `user_id`, `total_amount`, `status`, `remark`, `created_at`, `updated_at`) VALUES
(2, 'ORD20241129000002', 1, 1689800, 2, '测试订单2', NOW(), NOW());

-- 订单2的订单项
INSERT INTO `order_items` (`id`, `order_id`, `product_id`, `product_name`, `product_desc`, `price`, `quantity`, `amount`, `created_at`) VALUES
(2, 2, 2, 'MacBook Pro 14寸', 'M3芯片，14英寸 Liquid Retina XDR 显示屏', 1499900, 1, 1499900, NOW()),
(3, 2, 3, 'AirPods Pro', '主动降噪无线耳机，支持空间音频', 189900, 1, 189900, NOW());

-- 订单3：admin用户的订单（已发货）
INSERT INTO `orders` (`id`, `order_no`, `user_id`, `total_amount`, `status`, `remark`, `created_at`, `updated_at`) VALUES
(3, 'ORD20241129000003', 1, 439900, 3, '测试订单3', NOW(), NOW());

-- 订单3的订单项
INSERT INTO `order_items` (`id`, `order_id`, `product_id`, `product_name`, `product_desc`, `price`, `quantity`, `amount`, `created_at`) VALUES
(4, 3, 4, 'iPad Air', 'M2芯片，10.9英寸 Liquid Retina 显示屏', 439900, 1, 439900, NOW());

-- 订单4：admin用户的订单（已完成）
INSERT INTO `orders` (`id`, `order_no`, `user_id`, `total_amount`, `status`, `remark`, `created_at`, `updated_at`) VALUES
(4, 'ORD20241129000004', 1, 299900, 4, '测试订单4', NOW(), NOW());

-- 订单4的订单项
INSERT INTO `order_items` (`id`, `order_id`, `product_id`, `product_name`, `product_desc`, `price`, `quantity`, `amount`, `created_at`) VALUES
(5, 4, 5, 'Apple Watch Series 9', '智能手表，支持健康监测和运动追踪', 299900, 1, 299900, NOW());

-- 订单5：admin用户的订单（已取消）
INSERT INTO `orders` (`id`, `order_no`, `user_id`, `total_amount`, `status`, `remark`, `created_at`, `updated_at`) VALUES
(5, 'ORD20241129000005', 1, 239900, 5, '测试订单5（已取消）', NOW(), NOW());

-- 订单5的订单项
INSERT INTO `order_items` (`id`, `order_id`, `product_id`, `product_name`, `product_desc`, `price`, `quantity`, `amount`, `created_at`) VALUES
(6, 5, 6, 'Magic Keyboard', '适用于 iPad 的键盘，带触控板', 239900, 1, 239900, NOW());

-- 订单6：testuser用户的订单（多商品订单）
-- 【修复】总金额 = 99000 + 189900 = 288900
INSERT INTO `orders` (`id`, `order_no`, `user_id`, `total_amount`, `status`, `remark`, `created_at`, `updated_at`) VALUES
(6, 'ORD20241129000006', 2, 288900, 1, 'testuser的测试订单', NOW(), NOW());

-- 订单6的订单项（多商品）
INSERT INTO `order_items` (`id`, `order_id`, `product_id`, `product_name`, `product_desc`, `price`, `quantity`, `amount`, `created_at`) VALUES
(7, 6, 7, 'AirTag 4件装', '物品追踪器，帮助找回丢失的物品', 99000, 1, 99000, NOW()),
(8, 6, 3, 'AirPods Pro', '主动降噪无线耳机，支持空间音频', 189900, 1, 189900, NOW());

-- ============================================
-- 说明
-- ============================================
-- 1. 管理员账号：admin / 123456
-- 2. 测试账号：testuser / 123456（拥有管理员权限，测试脚本需要）
-- 3. 普通用户：normaluser / Normal123（无角色，测试脚本需要）
-- 4. 所有密码已使用 bcrypt 加密
-- 5. 测试脚本 test-permission-middleware.sh 会使用 testuser 和 normaluser 进行测试
-- 【扩展】6. 商品测试数据：8个商品（7个上架，1个下架）
-- 【扩展】7. 订单测试数据：6个订单，覆盖所有订单状态（待支付、已支付、已发货、已完成、已取消）
-- 【扩展】8. 订单项测试数据：8个订单项，包含单商品和多商品订单
-- 【扩展】9. 商品和订单权限已分配给管理员角色
-- 【扩展】10. 商品管理和订单管理菜单已添加到系统管理菜单下

