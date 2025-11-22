package errorx

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// HandleError 统一处理错误
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}

	// 判断是否为业务错误
	if bizErr, ok := err.(*BusinessError); ok {
		statusCode := GetHTTPStatus(bizErr.Code)
		httpx.WriteJson(w, statusCode, bizErr)
		return
	}

	// 其他错误，记录日志并返回通用错误
	logx.Errorf("处理请求失败：%v", err)
	httpx.WriteJson(w, http.StatusInternalServerError, &BusinessError{
		Code:    CodeInternalError,
		Message: "服务器内部错误",
	})
}
