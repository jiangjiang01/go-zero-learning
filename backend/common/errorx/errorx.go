package errorx

import (
	"fmt"
	"net/http"
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
)

// 业务错误
type BusinessError struct {
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误消息
}

// 创建业务错误（格式化消息）
func NewBusinessError(code int, message string, args ...interface{}) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: fmt.Sprintf(message, args...),
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
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound:
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
)
