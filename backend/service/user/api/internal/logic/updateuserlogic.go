package logic

import (
	"context"
	"go-zero-learning/common/ctxdata"
	"go-zero-learning/common/errorx"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"
	"go-zero-learning/service/user/user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UserInfoResp, err error) {
	// 1. 检查是否有需要更新的字段
	if req.Email == nil && req.Password == nil {
		return nil, errorx.ErrNoUpdateFields
	}

	// 2. 从上下文中获取用户 ID （由中间件设置）
	userID, ok := ctxdata.GetUserID(l.ctx)
	if !ok {
		return nil, errorx.ErrNoUserInfo
	}

	// 3. 调用 UserRpc.UpdateUser 更新用户信息
	rpcReq := &userrpc.UpdateUserReq{
		Id: userID,
	}
	if req.Email != nil {
		rpcReq.Email = *req.Email
	}
	if req.Password != nil {
		rpcReq.Password = *req.Password
	}

	rpcResp, err := l.svcCtx.UserRpc.UpdateUser(l.ctx, rpcReq)
	if err != nil {
		// 使用统一的错误映射函数
		if rpcErr := errorx.MapRpcError(err, l.Logger, "UserRpc.UpdateUser", errorx.RpcErrorMapper{
			NotFoundErr:      errorx.ErrUserNotFound,
			AlreadyExistsErr: errorx.ErrEmailExists,
		}); rpcErr != nil {
			return nil, rpcErr
		}
	}

	// 8. 返回更新后的用户信息
	resp = &types.UserInfoResp{
		ID:       rpcResp.Id,
		Username: rpcResp.Username,
		Email:    rpcResp.Email,
	}

	return resp, nil
}
