package response

import "time"

// Response 统一响应结构（成功和失败都使用这个结构）
type Response struct {
	Code      int         `json:"code"`                 // 状态码, 0 表示成功，非 0 表示失败
	Message   string      `json:"message"`              // 消息
	Data      interface{} `json:"data,omitempty"`       // 数据（成功时返回数据，失败时为空）
	Timestamp int64       `json:"timestamp,omitempty"`  // 时间戳（毫秒），可选
	RequestId string      `json:"request_id,omitempty"` // 请求ID（可选，后续扩展）
}

// SuccessResponse 创建成功响应
func SuccessResponse(data interface{}) *Response {
	return &Response{
		Code:      0,
		Message:   "success",
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
}

// SuccessResponse 创建成功响应（自定义消息）
func SuccessResponseWithMessage(data interface{}, message string) *Response {
	return &Response{
		Code:      0,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
}
