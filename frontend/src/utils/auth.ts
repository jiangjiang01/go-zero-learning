const TOKEN_KEY = 'admin_token'
const USER_INFO_KEY = 'user_info'

// 获取 Token
export function getToken(): string | null {
  return localStorage.getItem(TOKEN_KEY)
}

// 设置 Token
export function setToken(token: string): void {
  localStorage.setItem(TOKEN_KEY, token)
}

// 删除 Token
export function removeToken(): void {
  localStorage.removeItem(TOKEN_KEY)
}

// 获取用户信息
export function getUserInfo(): any {
  const userInfo = localStorage.getItem(USER_INFO_KEY)
  return userInfo ? JSON.parse(userInfo) : null
}

// 设置用户信息
export function setUserInfo(userInfo: any): void {
  localStorage.setItem(USER_INFO_KEY, JSON.stringify(userInfo))
}

// 删除用户信息
export function removeUserInfo(): void {
  localStorage.removeItem(USER_INFO_KEY)
}

