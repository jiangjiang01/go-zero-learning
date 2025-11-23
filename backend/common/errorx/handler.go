package errorx

import (
	"go-zero-learning/common/response"
	"net/http"
	"time"

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

		// 记录业务错误日志（使用 Info 级别，因为这是预期的业务错误）
		logx.Infof("[业务错误] 路径：%s | 错误码：%d | 错误信息：%s",
			r.URL.Path,
			bizErr.Code,
			bizErr.Message,
		)

		httpx.WriteJson(w, statusCode, bizErr.Response)
		return
	}

	// 其他错误，记录日志并返回通用错误
	logx.Errorf("[系统错误] 路径：%s | 错误：%v", r.URL.Path, err)

	httpx.WriteJson(w, http.StatusInternalServerError, &BusinessError{
		Response: &response.Response{
			Code:      CodeInternalError,
			Message:   "服务器内部错误",
			Timestamp: time.Now().Unix(),
		},
	})
}
