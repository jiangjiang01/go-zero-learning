package logic

import (
	"context"
	"go-zero-learning/common/ctxdata"
	"go-zero-learning/common/errorx"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"
	"go-zero-learning/service/user/user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.InvalidArgument:
				return nil, errorx.ErrInvalidParam
			case codes.NotFound:
				return nil, errorx.ErrUserNotFound
			default:
				l.Errorf("调用 UserRpc.DeleteUser 失败：code=%v, msg=%s", st.Code(), st.Message())
				return nil, errorx.ErrInternalError
			}
		}
		l.Errorf("调用 UserRpc.DeleteUser 失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 返回响应
	resp = &types.DeleteUserResp{
		Message: rpcResp.Message,
	}

	return resp, nil
}
