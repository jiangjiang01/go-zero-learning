package consts

import "time"

// ==================== 分页常量 ====================
const (
	DefaultPageSize = 10
	MaxPageSize     = 100
)

// ==================== 购物车常量 ====================
const (
	MaxCartQuantity = 999
)

// ==================== 分类常量 ====================
const (
	MaxCategoryDepth      = 100
	MaxCategoryNameLength = 100
	TopLevelParentID      = 0
)

// ==================== 商品常量 ====================
const (
	MaxProductNameLength = 100
	MinProductPrice      = 1        // 0.01元
	MaxProductPrice      = 99999900 // 9999.99元
)

// ==================== 菜单常量 ====================
const (
	MenuTypeMenu          = 1 // 菜单
	MenuTypeButton        = 2 // 按钮
	TopLevelMenuParentID  = 0 // 顶级菜单的父ID
)

// ==================== 缓存常量 ====================
const (
	CacheEmptyResultTTL    = 60  // 秒，空结果缓存时间
	CacheEmptyRandomRange  = 30  // 秒，空结果随机增加范围
	CacheResultTTL         = 300 // 秒，有结果缓存时间（5分钟）
	CacheResultRandomRange = 60  // 秒，有结果随机增加范围
)

// ==================== 订单常量 ====================
const (
	OrderCancelTimeout = 30 * time.Minute // 订单待支付超时时间
)

// ==================== 文件操作常量 ====================
const (
	FilePermissionMode = 0755              // Unix 文件目录权限
	DatePathFormat     = "2006/01/02"      // 文件路径日期格式 YYYY/MM/DD
)
