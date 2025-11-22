package ctxdata

import "context"

type ctxKey string

const (
	UserIDKey   ctxKey = "user_id"
	UsernameKey ctxKey = "username"
)

// 从上下文中获取用户 ID
func GetUserID(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey).(int64)
	return userID, ok
}

// 从上下文中获取用户名
func GetUsername(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(UsernameKey).(string)
	return username, ok
}
