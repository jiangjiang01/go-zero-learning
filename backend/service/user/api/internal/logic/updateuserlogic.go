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
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.InvalidArgument:
				// 包括“没有需要更新的字段”“ID非法”等
				return nil, errorx.ErrInvalidParam
			case codes.AlreadyExists:
				// 这里是邮箱已存在
				return nil, errorx.ErrEmailExists
			case codes.NotFound:
				return nil, errorx.ErrUserNotFound
			default:
				l.Errorf("调用 UserRpc.UpdateUser 失败：code=%v, msg=%s", st.Code(), st.Message())
				return nil, errorx.ErrInternalError
			}
		}
		l.Errorf("调用 UserRpc.UpdateUser 失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 8. 返回更新后的用户信息
	resp = &types.UserInfoResp{
		ID:       rpcResp.Id,
		Username: rpcResp.Username,
		Email:    rpcResp.Email,
	}

	return resp, nil
}
