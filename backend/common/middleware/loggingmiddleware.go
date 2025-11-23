package middleware

import (
	"go-zero-learning/common/ctxdata"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoggingMiddleware 请求日志中间件
type LoggingMiddleware struct {
}

// 创建请求日志中间件实例
func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (m *LoggingMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 从请求上下文中获取 logger (包含 trace 信息)
		logger := logx.WithContext(r.Context())

		// 记录请求开始
		logger.Infof(
			"[请求开始] %s %s | IP: %s | User-Agent: %s",
			r.Method,
			r.URL.Path,
			getClientIP(r),
			r.UserAgent(),
		)

		// 如果有用户信息，记录用户ID
		if userID, ok := ctxdata.GetUserID(r.Context()); ok {
			logger.Infof("[用户信息] UserID: %d", userID)
		}

		// 执行下一个处理器
		next(w, r)

		// 记录请求结束
		duration := time.Since(start)
		logger.Infof(
			"[请求结束] %s %s | 耗时：%v",
			r.Method,
			r.URL.Path,
			duration,
		)
	}
}

// getClientIP 获取客户端真实IP
func getClientIP(r *http.Request) string {
	// 1. 尝试从 X-Forwarded-For 头获取（代理/负载均衡）
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		return ip
	}

	// 2. 尝试从 X-Real-IP 获取
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	// 3. 使用 RemoteAddr
	return r.RemoteAddr
}
