/**
 * 判断是否为外部链接
 * @param path - 路径
 * @returns 是否为外部链接
 */
export function isExternal(path: string): boolean {
  return /^(https?:|mailto:|tel:)/.test(path)
}

