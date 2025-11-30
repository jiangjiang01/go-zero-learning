// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"go-zero-learning/common/errorx"
	"go-zero-learning/common/response"
	"go-zero-learning/service/user/api/internal/logic"
	"go-zero-learning/service/user/api/internal/svc"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 文件上传需要手动处理 multipart/form-data， 不能使用 httpx.Parse 解析
		l := logic.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile(r)
		if err != nil {
			errorx.HandleError(w, r, err)
		} else {
			response.OkJson(w, r, resp)
		}
	}
}
