package handler

import (
	"go-zero-learning/common/errorx"
	"go-zero-learning/common/response"
	"go-zero-learning/common/validator"
	"go-zero-learning/service/user/api/internal/logic"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			errorx.HandleError(w, r, validator.ParseError(err))
			return
		}

		l := logic.NewGetUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetUserDetail(&req)
		if err != nil {
			errorx.HandleError(w, r, err)
		} else {
			response.OkJson(w, r, resp)
		}
	}
}
