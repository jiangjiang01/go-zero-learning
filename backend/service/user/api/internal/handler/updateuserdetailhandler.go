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

func UpdateUserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			errorx.HandleError(w, r, validator.ParseError(err))
			return
		}

		l := logic.NewUpdateUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserDetail(&req)
		if err != nil {
			errorx.HandleError(w, r, err)
		} else {
			response.OkJson(w, r, resp)
		}
	}
}
