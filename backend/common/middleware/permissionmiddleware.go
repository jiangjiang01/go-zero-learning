package middleware

import (
	"go-zero-learning/common/ctxdata"
	"go-zero-learning/common/errorx"
	"go-zero-learning/model"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// PermissionMiddleware 权限验证中间件
type PermissionMiddleware struct {
	db *gorm.DB
}

// 创建权限验证中间件实例
func NewPermissionMiddleware(db *gorm.DB) *PermissionMiddleware {
	return &PermissionMiddleware{
		db: db,
	}
}

// Handle 权限验证处理函数
// re
func (m *PermissionMiddleware) Handle(requiredPermission string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger := logx.WithContext(r.Context())

			// 1. 从 context 获取用户 ID (由 authmiddleware 设置)
			userID, ok := ctxdata.GetUserID(r.Context())
			if !ok {
				errorx.HandleError(w, r, errorx.ErrUnauthorized)
				return
			}

			// 2. 查询用户角色
			var userRoles []model.UserRole
			err := m.db.Where("user_id = ?", userID).Find(&userRoles).Error
			if err != nil {
				logger.Errorf("查询用户角色失败：%v", err)
				errorx.HandleError(w, r, errorx.ErrInternalError)
				return
			}
			// 用户没有分配角色，拒绝访问
			if len(userRoles) == 0 {
				errorx.HandleError(w, r, errorx.NewBusinessError(errorx.CodeForbidden, "没有权限访问"))
				return
			}

			// 3. 收集所有角色 ID
			roleIDs := make([]int64, 0, len(userRoles))
			for _, role := range userRoles {
				roleIDs = append(roleIDs, role.RoleID)
			}

			// 4. 检查这些角色的所有权限
			var rolePermissions []model.RolePermission
			err = m.db.Where("role_id IN ?", roleIDs).Find(&rolePermissions).Error
			if err != nil {
				logger.Errorf("查询角色权限失败：%v", err)
				errorx.HandleError(w, r, errorx.ErrInternalError)
				return
			}

			// 5. 收集所有的权限 ID
			permissionIDs := make([]int64, 0, len(rolePermissions))
			for _, rolePermission := range rolePermissions {
				permissionIDs = append(permissionIDs, rolePermission.PermissionID)
			}

			// 6. 查询权限代码
			var permissions []model.Permission
			err = m.db.Where("id IN ?", permissionIDs).Find(&permissions).Error
			if err != nil {
				logger.Errorf("查询权限失败：%v", err)
				errorx.HandleError(w, r, errorx.ErrInternalError)
				return
			}

			// 7. 检查是否有需要的权限
			hasPermission := false
			for _, p := range permissions {
				if p.Code == requiredPermission {
					hasPermission = true
					break
				}
			}

			if !hasPermission {
				errorx.HandleError(w, r, errorx.NewBusinessError(errorx.CodeForbidden, "没有权限访问"))
				return
			}

			// 8. 有权限，继续处理请求
			next(w, r)
		}
	}
}
