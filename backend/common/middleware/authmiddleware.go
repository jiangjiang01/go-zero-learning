package middleware

import (
	"context"
	"errors"
	"go-zero-learning/common/ctxdata"
	"go-zero-learning/common/jwt"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 权限中间件
type AuthMiddleware struct {
	logx.LessLogger
}

// 创建权限中间件实例
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 从请求头获取 Token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("未提取认证 token"))
			return
		}

		// 2. 解析 Token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			httpx.ErrorCtx(r.Context(), w, errors.New("token 格式错误"))
			return
		}

		tokenStr := parts[1]

		// 3. 解析 Token
		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			m.Errorf("token 验证失败：%v", err)
			httpx.ErrorCtx(r.Context(), w, errors.New("token 无效或已过期"))
			return
		}

		// 4. 将用户信息存储到请求上下文中（供后续使用）
		// go-zero 使用 context 传递值
		ctx := r.Context()
		// 这里直接使用字符串会有警告，可能会导致多个中间件使用相同的 key 导致被覆盖
		// 使用自定义的 key 类型可以避免这个问题
		ctx = context.WithValue(ctx, ctxdata.UserIDKey, claims.UserID)     // 使用自定义的 key
		ctx = context.WithValue(ctx, ctxdata.UsernameKey, claims.Username) // 使用自定义的 key

		// 5. 继续处理请求（下一个中间件）
		next(w, r.WithContext(ctx))
	}
}
