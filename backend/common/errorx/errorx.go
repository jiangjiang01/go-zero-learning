package errorx

import (
	"fmt"
	"go-zero-learning/common/response"
	"net/http"
	"time"
)

// 错误码定义
const (
	// 通用错误码（1000-1999）
	CodeSuccess       = 0    // 成功
	CodeInvalidParam  = 1001 // 参数错误
	CodeUnauthorized  = 1002 // 未授权
	CodeForbidden     = 1003 // 禁止访问
	CodeNotFound      = 1004 // 资源不存在
	CodeInternalError = 1005 // 内部错误
	CodeInvalidToken  = 1006 // Token 无效或已过期

	// 用户相关错误码（2000-2999）
	CodeUserNotFound     = 2001 // 用户不存在
	CodeUserAlreadExists = 2002 // 用户已存在
	CodeUsernameExists   = 2003 // 用户名已存在
	CodeEmailExists      = 2004 // 邮箱已存在
	CodeInvalidPassword  = 2005 // 密码错误
	CodePasswordTooShort = 2006 // 密码长度不足
	CodeInvalidEmail     = 2007 // 邮箱格式不正确
	CodeCannotDeleteSelf = 2008 // 不能删除自己
	CodeNoUpdateFields   = 2009 // 没有需要更新的字段
	CodeNoUserInfo       = 2010 // 未找到用户信息

	// 角色相关错误码（3000-3999）
	CodeRoleNotFound       = 3001 // 角色不存在
	CodeRoleAlreadyExists  = 3002 // 角色已存在
	CodeRoleNameExists     = 3003 // 角色名称已存在
	CodeRoleCodeExists     = 3004 // 角色代码已存在
	CodeRoleNoUpdateFields = 3005 // 没有需要更新的字段

	// 权限相关错误码（4000-4999）
	CodePermissionNotFound       = 4001 // 权限不存在
	CodePermissionAlreadyExists  = 4002 // 权限已存在
	CodePermissionNameExists     = 4003 // 权限名称已存在
	CodePermissionCodeExists     = 4004 // 权限代码已存在
	CodePermissionNoUpdateFields = 4005 // 没有需要更新的字段

	// 菜单相关错误码（5000-5999）
	CodeMenuNotFound       = 5001 // 菜单不存在
	CodeMenuAlreadyExists  = 5002 // 菜单已存在
	CodeMenuCodeExists     = 5003 // 菜单代码已存在
	CodeMenuNoUpdateFields = 5004 // 没有需要更新的字段
	CodeMenuHasChildren    = 5005 // 菜单下有子菜单，不能删除
	CodeMenuCircularRef    = 5006 // 不能将父菜单设置为自己的子菜单

	// 商品相关错误码（6000-6999）
	CodeProductNotFound       = 6001 // 商品不存在
	CodeProductAlreadyExists  = 6002 // 商品已存在(id)
	CodeProductNameExists     = 6003 // 商品名称已存在
	CodeProductPriceTooLow    = 6004 // 商品价格太低
	CodeProductPriceTooHigh   = 6005 // 商品价格太高
	CodeProductNoUpdateFields = 6006 // 没有需要更新的字段
	CodeProductStatusInvalid  = 6007 // 商品状态无效

	// 订单相关错误码（7000-7999）
	CodeOrderNotFound        = 7001 // 订单不存在
	CodeOrderStatusInvalid   = 7002 // 订单状态无效
	CodeOrderCannotCancel    = 7003 // 订单不能取消
	CodeOrderCannotPay       = 7004 // 订单不能支付
	CodeOrderQuantityInvalid = 7005 // 订单数量无效
	CodeOrderPriceChange     = 7006 // 商品价格已变更
	CodeOrderStockNotEnough  = 7007 // 库存不足
	CodeOrderItemsEmpty      = 7008 // 订单商品列表不能为空
	CodeOrderAmountMismatch  = 7009 // 订单金额不匹配
	CodeOrderNotBelongToUser = 7010 // 订单不属于当前用户

	// 商品分类相关错误码（8000-8999）
	CodeCategoryNotFound       = 8001 // 商品分类不存在
	CodeCategoryNameExists     = 8003 // 商品分类名称已存在
	CodeCategoryHasChildren    = 8004 // 商品分类下有子分类，不能删除
	CodeCategoryHasProducts    = 8005 // 商品分类下有商品，不能删除
	CodeCategoryParentInvalid  = 8006 // 父分类无效
	CodeCategoryNoUpdateFields = 8007 // 没有需要更新的字段
)

// 业务错误（使用统一的 Response 结构）
type BusinessError struct {
	*response.Response
}

// 创建业务错误
func NewBusinessError(code int, message string) *BusinessError {
	return &BusinessError{
		Response: &response.Response{
			Code:      code,
			Message:   message,
			Timestamp: time.Now().Unix(),
		},
	}
}

// 创建业务错误（格式化消息）
func NewBusinessErrorf(code int, format string, args ...interface{}) *BusinessError {
	return &BusinessError{
		Response: &response.Response{
			Code:      code,
			Message:   fmt.Sprintf(format, args...),
			Timestamp: time.Now().Unix(),
		},
	}
}

// 实现 error 接口
func (e *BusinessError) Error() string {
	return e.Message
}

// 根据错误码获取 HTTP 状态码
func GetHTTPStatus(code int) int {
	switch code {
	case CodeSuccess:
		return http.StatusOK
	case CodeInvalidParam:
		return http.StatusBadRequest
	case CodeUnauthorized, CodeInvalidToken:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound,
		CodeUserNotFound,
		CodeRoleNotFound,
		CodePermissionNotFound,
		CodeMenuNotFound,
		CodeProductNotFound,
		CodeOrderNotFound,
		CodeCategoryNotFound:
		return http.StatusNotFound
	case CodeInternalError:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}

// 预定义的业务错误
var (
	ErrInvalidParam  = NewBusinessError(CodeInvalidParam, "参数错误")
	ErrUnauthorized  = NewBusinessError(CodeUnauthorized, "未授权")
	ErrInvalidToken  = NewBusinessError(CodeInvalidToken, "token 无效或已过期")
	ErrForbidden     = NewBusinessError(CodeForbidden, "禁止访问")
	ErrNotFound      = NewBusinessError(CodeNotFound, "资源不存在")
	ErrInternalError = NewBusinessError(CodeInternalError, "内部错误")

	ErrUserNotFound     = NewBusinessError(CodeUserNotFound, "用户不存在")
	ErrUserAlreadExists = NewBusinessError(CodeUserAlreadExists, "用户已存在")
	ErrUsernameExists   = NewBusinessError(CodeUsernameExists, "用户名已存在")
	ErrEmailExists      = NewBusinessError(CodeEmailExists, "邮箱已存在")
	ErrInvalidPassword  = NewBusinessError(CodeInvalidPassword, "用户名或密码错误")
	ErrPasswordTooShort = NewBusinessError(CodePasswordTooShort, "密码至少需要6位")
	ErrInvalidEmail     = NewBusinessError(CodeInvalidEmail, "邮箱格式不正确")
	ErrCannotDeleteSelf = NewBusinessError(CodeCannotDeleteSelf, "不能删除自己的账户")
	ErrNoUpdateFields   = NewBusinessError(CodeNoUpdateFields, "至少需要提供一个更新字段")
	ErrNoUserInfo       = NewBusinessError(CodeNoUserInfo, "未找到用户信息")

	// 角色相关错误
	ErrRoleNotFound       = NewBusinessError(CodeRoleNotFound, "角色不存在")
	ErrRoleAlreadExists   = NewBusinessError(CodeRoleAlreadyExists, "角色已存在")
	ErrRoleNameExists     = NewBusinessError(CodeRoleNameExists, "角色名称已存在")
	ErrRoleCodeExists     = NewBusinessError(CodeRoleCodeExists, "角色代码已存在")
	ErrRoleNoUpdateFields = NewBusinessError(CodeRoleNoUpdateFields, "至少需要提供一个更新字段")

	// 权限相关错误
	ErrPermissionNotFound       = NewBusinessError(CodePermissionNotFound, "权限不存在")
	ErrPermissionAlreadExists   = NewBusinessError(CodePermissionAlreadyExists, "权限已存在")
	ErrPermissionNameExists     = NewBusinessError(CodePermissionNameExists, "权限名称已存在")
	ErrPermissionCodeExists     = NewBusinessError(CodePermissionCodeExists, "权限代码已存在")
	ErrPermissionNoUpdateFields = NewBusinessError(CodePermissionNoUpdateFields, "至少需要提供一个更新字段")

	// 菜单相关错误
	ErrMenuNotFound       = NewBusinessError(CodeMenuNotFound, "菜单不存在")
	ErrMenuAlreadyExists  = NewBusinessError(CodeMenuAlreadyExists, "菜单已存在")
	ErrMenuCodeExists     = NewBusinessError(CodeMenuCodeExists, "菜单代码已存在")
	ErrMenuNoUpdateFields = NewBusinessError(CodeMenuNoUpdateFields, "至少需要提供一个更新字段")
	ErrMenuHasChildren    = NewBusinessError(CodeMenuHasChildren, "菜单下有子菜单，不能删除")
	ErrMenuCircularRef    = NewBusinessError(CodeMenuCircularRef, "不能将父菜单设置为自己的子菜单")

	// 商品相关错误
	ErrProductNotFound       = NewBusinessError(CodeProductNotFound, "商品不存在")
	ErrProductAlreadyExists  = NewBusinessError(CodeProductAlreadyExists, "商品已存在")
	ErrProductNameExists     = NewBusinessError(CodeProductNameExists, "商品名称已存在")
	ErrProductPriceTooLow    = NewBusinessError(CodeProductPriceTooLow, "商品价格太低")
	ErrProductPriceTooHigh   = NewBusinessError(CodeProductPriceTooHigh, "商品价格太高")
	ErrProductNoUpdateFields = NewBusinessError(CodeProductNoUpdateFields, "没有需要更新的字段")
	ErrProductStatusInvalid  = NewBusinessError(CodeProductStatusInvalid, "商品状态无效")

	// 订单相关错误
	ErrOrderNotFound        = NewBusinessError(CodeOrderNotFound, "订单不存在")
	ErrOrderStatusInvalid   = NewBusinessError(CodeOrderStatusInvalid, "订单状态无效")
	ErrOrderCannotCancel    = NewBusinessError(CodeOrderCannotCancel, "订单不能取消")
	ErrOrderCannotPay       = NewBusinessError(CodeOrderCannotPay, "订单不能支付")
	ErrOrderQuantityInvalid = NewBusinessError(CodeOrderQuantityInvalid, "订单数量无效")
	ErrOrderPriceChange     = NewBusinessError(CodeOrderPriceChange, "商品价格已变更")
	ErrOrderStockNotEnough  = NewBusinessError(CodeOrderStockNotEnough, "库存不足")
	ErrOrderItemsEmpty      = NewBusinessError(CodeOrderItemsEmpty, "订单商品列表不能为空")
	ErrOrderAmountMismatch  = NewBusinessError(CodeOrderAmountMismatch, "订单金额不匹配")
	ErrOrderNotBelongToUser = NewBusinessError(CodeOrderNotBelongToUser, "订单不属于当前用户")

	// 商品分类相关错误
	ErrCategoryNotFound       = NewBusinessError(CodeCategoryNotFound, "商品分类不存在")
	ErrCategoryNameExists     = NewBusinessError(CodeCategoryNameExists, "商品分类名称已存在")
	ErrCategoryHasChildren    = NewBusinessError(CodeCategoryHasChildren, "商品分类下有子分类，不能删除")
	ErrCategoryHasProducts    = NewBusinessError(CodeCategoryHasProducts, "商品分类下有商品，不能删除")
	ErrCategoryParentInvalid  = NewBusinessError(CodeCategoryParentInvalid, "父分类无效")
	ErrCategoryNoUpdateFields = NewBusinessError(CodeCategoryNoUpdateFields, "没有需要更新的字段")
)
