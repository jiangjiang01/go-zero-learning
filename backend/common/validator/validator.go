package validator

import (
	"fmt"
	"go-zero-learning/common/errorx"
	"strings"
)

// ParseError 将 httpx.Parse 的错误转换为友好的参数错误
func ParseError(err error) *errorx.BusinessError {
	if err == nil {
		return nil
	}

	errMsg := err.Error()

	// 处理常见的参数解析错误
	// 1. 字段未设置错误：field "password" is not set
	if strings.Contains(errMsg, "is not set") {
		// 提取字段名
		fieldName := extractFieldName(errMsg)
		if fieldName != "" {
			return errorx.NewBusinessError(errorx.CodeInvalidParam, fmt.Sprintf("参数 %s 不能为空", translateFieldName(fieldName)))
		}
		return errorx.NewBusinessError(errorx.CodeInvalidParam, "缺少必填参数")

	}

	// 2. 字段格式错误：invalid character 'x' looking for beginning of value
	if strings.Contains(errMsg, "invalid character") {
		return errorx.NewBusinessErrorf(errorx.CodeInvalidParam, "参数格式错误，请检查 JSON 格式")
	}

	// 3. 路径参数错误：strconv.ParseInt: parsing "abc": invalid syntax
	if strings.Contains(errMsg, "parsing") && strings.Contains(errMsg, "invalid syntax") {
		return errorx.NewBusinessError(errorx.CodeInvalidParam, "路径参数格式错误")
	}

	// 4. 其他错误，返回通用参数错误
	return errorx.NewBusinessError(errorx.CodeInvalidParam, fmt.Sprintf("参数错误：%s", errMsg))
}

// extractFieldName 从错误信息中提取字段名
func extractFieldName(errMsg string) string {
	// 匹配：field "password" is not set，返回 "password"
	prefix := `field "`
	if idx := strings.Index(errMsg, prefix); idx != -1 {
		start := idx + len(prefix)
		if end := strings.Index(errMsg[start:], `"`); end != -1 {
			return errMsg[start : start+end]
		}
	}

	return ""
}

// translateFieldName 将字段名翻译成中文
func translateFieldName(fieldName string) string {
	fieldMap := map[string]string{
		"username":      "用户名",
		"password":      "密码",
		"email":         "邮箱",
		"id":            "ID",
		"page":          "页码",
		"page_size":     "每页数量",
		"keyword":       "搜索关键词",
		"token":         "Token",
		"authorization": "Authorization",
	}

	if chinese, ok := fieldMap[strings.ToLower(fieldName)]; ok {
		return chinese
	}

	return fieldName
}
