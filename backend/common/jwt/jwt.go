package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWT 管理器
type JWTManager struct {
	secret               []byte
	tokenExpireDuaration time.Duration
}

// 创建 JWT 管理器实例
func NewJWTManager(secret string, expireDays int) *JWTManager {
	return &JWTManager{
		secret:               []byte(secret),
		tokenExpireDuaration: time.Duration(expireDays) * 24 * time.Hour,
	}
}

// Claims JWT 载荷结构(包含用户信息)
type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT Token
func (m *JWTManager) GenerateToken(userID int64, username string) (string, error) {
	// 创建 claims
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.tokenExpireDuaration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                             // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                             // 生效时间
			Issuer:    "jwt-basic",                                                // 签发者
		},
	}

	// 使用 HS256 算法签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名字符串
	tokenStr, err := token.SignedString(m.secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// 解析 JWT Token
func (m *JWTManager) ParseToken(tokenStr string) (*Claims, error) {
	// 解析 Token
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return m.secret, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证 token 是否有效
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
