package logic

import (
	"context"
	"go-zero-learning/common/errorx"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"
	"go-zero-learning/service/user/user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailLogic) GetUserDetail(req *types.GetUserDetailReq) (resp *types.UserInfoResp, err error) {
	// 1. 参数校验
	if req.ID <= 0 {
		return nil, errorx.ErrInvalidParam
	}

	// 2. 调用 UserRpc.GetUser （替换直接访问 DB）
	rpcResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userrpc.GetUserReq{
		Id: req.ID,
	})
	if err != nil {
		// 使用统一的错误映射函数
		if rpcErr := errorx.MapRpcError(err, l.Logger, "UserRpc.GetUser", errorx.RpcErrorMapper{
			NotFoundErr: errorx.ErrUserNotFound,
		}); rpcErr != nil {
			return nil, rpcErr
		}
	}

	// 3. 构建响应结果
	resp = &types.UserInfoResp{
		ID:       rpcResp.Id,
		Username: rpcResp.Username,
		Email:    rpcResp.Email,
	}

	// 4. 返回响应
	return resp, nil
}

// 备份旧的写法
/*
func (l *GetUserDetailLogic) GetUserDetailOld(req *types.GetUserDetailReq) (resp *types.UserInfoResp, err error) {
	// 1. 参数校验
	if req.ID <= 0 {
		return nil, errorx.ErrInvalidParam
	}

	// 2. 查询用户信息
	var user model.User
	if err = l.svcCtx.DB.First(&user, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrUserNotFound
		}
		l.Errorf("查询用户失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 3. 构建响应结果
	resp = &types.UserInfoResp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	// 4. 返回响应
	return resp, nil
}
*/
