// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"go-zero-learning/common/errorx"
	"go-zero-learning/common/response"
	"go-zero-learning/common/validator"
	"go-zero-learning/service/user/api/internal/logic"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddCartItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCartItemReq
		if err := httpx.Parse(r, &req); err != nil {
			errorx.HandleError(w, r, validator.ParseError(err))
			return
		}

		l := logic.NewAddCartItemLogic(r.Context(), svcCtx)
		resp, err := l.AddCartItem(&req)
		if err != nil {
			errorx.HandleError(w, r, err)
		} else {
			response.OkJson(w, r, resp)
		}
	}
}
