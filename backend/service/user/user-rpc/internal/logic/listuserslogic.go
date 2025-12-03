package logic

import (
	"context"
	"strings"

	"go-zero-learning/model"
	"go-zero-learning/service/user/user-rpc/internal/svc"
	"go-zero-learning/service/user/user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ListUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUsersLogic {
	return &ListUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户列表（分页， 无搜索条件）
func (l *ListUsersLogic) ListUsers(in *userrpc.ListUsersReq) (*userrpc.ListUsersResp, error) {
	// 1. 参数校验
	if in.Page <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Page 必须大于 0")
	}
	if in.PageSize <= 0 || in.PageSize > 100 {
		return nil, status.Error(codes.InvalidArgument, "PageSize 必须在1~100之间")
	}

	offset := (in.Page - 1) * in.PageSize
	limit := in.PageSize

	var users []model.User
	var total int64

	db := l.svcCtx.DB.WithContext(l.ctx).Model(&model.User{})

	// 2. 关键词搜索（用户名或邮箱）
	if in.Keyword != "" {
		keyword := strings.TrimSpace(in.Keyword)
		likeStr := "%" + keyword + "%"
		db = db.Where("username LIKE ? OR email LIKE ?", likeStr, likeStr)
	}

	// 3. 查询总数
	if err := db.Count(&total).Error; err != nil {
		l.Errorf("查询用户总数失败：%v", err)
		return nil, status.Error(codes.Internal, "内部错误")
	}

	// 4. 查询当前页数据
	if err := db.Order("id DESC").Offset(int(offset)).Limit(int(limit)).Find(&users).Error; err != nil {
		l.Errorf("查询用户列表失败：%v", err)
		return nil, status.Error(codes.Internal, "内部错误")
	}

	// 5. 构建响应
	resp := &userrpc.ListUsersResp{
		Users: make([]*userrpc.UserItem, 0, len(users)),
		Total: total,
	}

	for _, user := range users {
		resp.Users = append(resp.Users, &userrpc.UserItem{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return resp, nil
}
