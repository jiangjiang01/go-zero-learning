package logic

import (
	"context"
	"fmt"
	"strings"

	"user-rpc/internal/svc"
	"user-rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *userrpc.CreateUserReq) (*userrpc.CreateUserResp, error) {
	// 1. 参数验证
	if strings.TrimSpace(in.Username) == "" {
		return nil, status.Errorf(codes.InvalidArgument, "用户名不能为空")
	}

	if strings.TrimSpace(in.Email) == "" {
		return nil, status.Errorf(codes.InvalidArgument, "邮箱不能为空")
	}

	if !strings.Contains(in.Email, "@") {
		return nil, status.Errorf(codes.InvalidArgument, "邮箱格式不正确")
	}

	// 2. 唯一性检查（检查用户名和邮箱是否已存在）
	l.svcCtx.UserMutex.RLock()
	for _, user := range l.svcCtx.UserStore {
		if user.Username == in.Username {
			l.svcCtx.UserMutex.RUnlock()
			return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("用户名已存在：%s", in.Username))
		}
		if user.Email == in.Email {
			l.svcCtx.UserMutex.RUnlock()
			return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("邮箱已存在：%s", in.Email))
		}
	}
	l.svcCtx.UserMutex.RUnlock()

	// 3. 生成 ID
	l.svcCtx.IdMutex.Lock()
	id := l.svcCtx.NextID
	l.svcCtx.NextID++
	l.svcCtx.IdMutex.Unlock()

	// 4. 创建用户对象
	user := &userrpc.User{
		Id:       id,
		Username: in.Username,
		Email:    in.Email,
	}

	// 5. 存储到内存
	l.svcCtx.UserMutex.Lock()
	l.svcCtx.UserStore[id] = user
	l.svcCtx.UserMutex.Unlock()

	l.Infof("创建用户成功：ID=%d, Username=%s, Email=%s", id, in.Username, in.Email)

	return &userrpc.CreateUserResp{
		User: user,
	}, nil
}
