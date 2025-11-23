/**
 * 验证用户是否有某个权限
 * @param permissions - 用户权限列表
 * @param permission - 需要验证的权限
 * @returns 是否有权限
 */
export function hasPermission(permissions: string[] | undefined, permission: string): boolean {
  if (!permissions || permissions.length === 0) {
    return false
  }
  return permissions.includes(permission)
}

/**
 * 验证用户是否有某个角色
 * @param roles - 用户角色列表
 * @param role - 需要验证的角色
 * @returns 是否有角色
 */
export function hasRole(roles: string[] | undefined, role: string): boolean {
  if (!roles || roles.length === 0) {
    return false
  }
  return roles.includes(role)
}

/**
 * 过滤有权限的菜单
 * @param menus - 菜单列表
 * @param permissions - 用户权限列表
 * @returns 过滤后的菜单列表
 */
export function filterMenus(menus: any[], permissions: string[]): any[] {
  if (!menus || menus.length === 0) {
    return []
  }

  return menus
    .filter((menu) => {
      // 如果没有权限要求，则显示
      if (!menu.permission) {
        return true
      }
      // 如果有权限要求，则验证权限
      return hasPermission(permissions, menu.permission)
    })
    .map((menu) => {
      // 递归处理子菜单
      if (menu.children && menu.children.length > 0) {
        menu.children = filterMenus(menu.children, permissions)
      }
      return menu
    })
}

