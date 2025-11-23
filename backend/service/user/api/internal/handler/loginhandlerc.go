// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"go-zero-learning/common/errorx"
	"go-zero-learning/common/response"
	"go-zero-learning/service/user/api/internal/logic"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			errorx.HandleError(w, r, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			errorx.HandleError(w, r, err)
		} else {
			response.OkJson(w, r, resp)
		}
	}
}
