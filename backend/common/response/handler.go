package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// OkJson 统一成功响应
func OkJson(w http.ResponseWriter, r *http.Request, data interface{}) {
	httpx.OkJsonCtx(r.Context(), w, SuccessResponse(data))
}

// OkJsonWithMessage 统一成功响应（自定义消息）
func OkJsonWithMessage(w http.ResponseWriter, r *http.Request, data interface{}, message string) {
	httpx.OkJsonCtx(r.Context(), w, SuccessResponseWithMessage(data, message))
}
