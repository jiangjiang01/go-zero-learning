package logic

import (
	"context"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.GetUserListResp, err error) {
	// 查询所有用户
	var users []model.User
	if err = l.svcCtx.DB.Find(&users).Error; err != nil {
		l.Errorf("查询用户列表失败：%v", err)
		return nil, err
	}

	// 转换为响应格式
	userList := make([]types.UserInfoResp, 0, len(users))
	for _, user := range users {
		userList = append(userList, types.UserInfoResp{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	resp = &types.GetUserListResp{
		Users: userList,
	}
	return resp, nil
}
