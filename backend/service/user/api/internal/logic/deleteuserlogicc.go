package logic

import (
	"context"
	"errors"
	"go-zero-learning/common/ctxdata"
	"go-zero-learning/model"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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
		return nil, errors.New("无效的用户ID")
	}

	// 2. 获取当前登录用户ID(防止用户删除自己)
	userID, ok := ctxdata.GetUserID(l.ctx)
	if ok && userID == req.ID {
		return nil, errors.New("不能删除自己的账户")
	}

	// 3. 查询用户是否存在
	var user model.User
	if err = l.svcCtx.DB.First(&user, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		l.Errorf("查询用户失败：%v", err)
		return nil, errors.New("删除失败")
	}
	// 4. 删除用户
	if err = l.svcCtx.DB.Delete(&user).Error; err != nil {
		l.Errorf("删除用户失败：%v", err)
		return nil, errors.New("删除失败")
	}

	// 5. 返回响应
	resp = &types.DeleteUserResp{
		Message: "用户删除成功",
	}

	return resp, nil
}
