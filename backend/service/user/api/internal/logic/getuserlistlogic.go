package logic

import (
	"context"
	"go-zero-learning/common/consts"
	"go-zero-learning/common/errorx"
	"go-zero-learning/service/user/api/internal/svc"
	"go-zero-learning/service/user/api/internal/types"
	"go-zero-learning/service/user/user-rpc/userrpc"

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
	// 1. 参数校验和默认值设置
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = consts.DefaultPageSize
	}
	// 限制每页最大数量，防止过大查询
	if req.PageSize > consts.MaxPageSize {
		req.PageSize = consts.MaxPageSize
	}

	// 2. 调用 UserRpc.ListUsers（替代直接访问DB）
	rpcResp, err := l.svcCtx.UserRpc.ListUsers(l.ctx, &userrpc.ListUsersReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
	})
	if err != nil {
		// 使用统一的错误映射函数
		if rpcErr := errorx.MapRpcError(err, l.Logger, "UserRpc.ListUsers", errorx.RpcErrorMapper{}); rpcErr != nil {
			return nil, rpcErr
		}
	}

	// 3. 转换为 API 的响应类型
	userList := make([]types.UserInfoResp, 0, len(rpcResp.Users))
	for _, user := range rpcResp.Users {
		userList = append(userList, types.UserInfoResp{
			ID:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	// 4. 返回响应
	resp = &types.GetUserListResp{
		Users:    userList,
		Total:    rpcResp.Total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	return resp, nil
}

// 备份旧的写法
/*
func (l *GetUserListLogic) GetUserListOld(req *types.GetUserListReq) (resp *types.GetUserListResp, err error) {
	// 1. 参数校验和默认值设置
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}
	// 限制每页最大数量，防止过大查询
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	// 2. 构建查询条件
	db := l.svcCtx.DB.Model(&model.User{})

	// 搜索关键词（用户名或邮箱）
	if req.Keyword != "" {
		keyword := strings.TrimSpace(req.Keyword)
		likeStr := "%" + keyword + "%"
		db = db.Where("username LIKE ? OR email LIKE ?", likeStr, likeStr)
	}

	// 3. 查询总数
	var total int64
	if err = db.Count(&total).Error; err != nil {
		l.Errorf("查询用户总数失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 4. 分页查询用户列表
	var users []model.User
	offset := (req.Page - 1) * req.PageSize
	if err = db.Offset(int(offset)).Limit(int(req.PageSize)).Find(&users).Error; err != nil {
		l.Errorf("查询用户列表失败：%v", err)
		return nil, errorx.ErrInternalError
	}

	// 5. 转换为响应格式
	userList := make([]types.UserInfoResp, 0, len(users))
	for _, user := range users {
		userList = append(userList, types.UserInfoResp{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	// 6. 返回响应
	resp = &types.GetUserListResp{
		Users:    userList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	return resp, nil
}
*/
