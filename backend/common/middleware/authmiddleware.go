package middleware

import (
	"context"
	"go-zero-learning/common/ctxdata"
	"go-zero-learning/common/errorx"
	"go-zero-learning/common/jwt"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// 权限中间件
type AuthMiddleware struct {
	logx.Logger
	jwtManager *jwt.JWTManager
}

// 创建权限中间件实例
func NewAuthMiddleware(jwtManager *jwt.JWTManager) *AuthMiddleware {
	return &AuthMiddleware{
		jwtManager: jwtManager,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 从请求头获取 Token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			errorx.HandleError(w, r, errorx.ErrUnauthorized)
			return
		}

		// 2. 解析 Token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			errorx.HandleError(w, r, errorx.NewBusinessError(errorx.CodeInvalidParam, "token 格式错误"))
			return
		}

		tokenStr := parts[1]

		// 3. 解析 Token
		claims, err := m.jwtManager.ParseToken(tokenStr)
		if err != nil {
			m.Errorf("token 验证失败：%v", err)
			errorx.HandleError(w, r, errorx.ErrInvalidToken)
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
