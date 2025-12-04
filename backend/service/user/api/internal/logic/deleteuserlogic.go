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

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	// 1. 参数校验
	if req.ID <= 0 {
		return nil, errorx.ErrInvalidParam
	}

	// 2. 获取当前登录用户ID(防止用户删除自己)
	userID, ok := ctxdata.GetUserID(l.ctx)
	if ok && userID == req.ID {
		return nil, errorx.ErrCannotDeleteSelf
	}

	// 3. 调用 RPC 删除用户
	rpcResp, err := l.svcCtx.UserRpc.DeleteUser(l.ctx, &userrpc.DeleteUserReq{
		Id: req.ID,
	})
	if err != nil {
		// 使用统一的错误映射函数
		if rpcErr := errorx.MapRpcError(err, l.Logger, "UserRpc.DeleteUser", errorx.RpcErrorMapper{
			NotFoundErr: errorx.ErrUserNotFound,
		}); rpcErr != nil {
			return nil, rpcErr
		}
	}

	// 4. 返回响应
	resp = &types.DeleteUserResp{
		Message: rpcResp.Message,
	}

	return resp, nil
}
