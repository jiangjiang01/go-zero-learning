/**
 * 格式化日期时间
 * @param date - 日期（支持 Date、字符串、Unix 时间戳（秒））
 * @param format - 格式，默认 'YYYY-MM-DD HH:mm:ss'
 * @returns 格式化后的日期字符串
 */
export function formatDateTime(date: string | Date | number | null | undefined, format = 'YYYY-MM-DD HH:mm:ss'): string {
  if (!date) return ''
  
  // 如果是数字（Unix 时间戳，秒级），转换为毫秒
  const timestamp = typeof date === 'number' ? date * 1000 : date
  const d = new Date(timestamp)
  const year = d.getFullYear()
  const month = String(d.getMonth() + 1).padStart(2, '0')
  const day = String(d.getDate()).padStart(2, '0')
  const hours = String(d.getHours()).padStart(2, '0')
  const minutes = String(d.getMinutes()).padStart(2, '0')
  const seconds = String(d.getSeconds()).padStart(2, '0')

  return format
    .replace('YYYY', String(year))
    .replace('MM', month)
    .replace('DD', day)
    .replace('HH', hours)
    .replace('mm', minutes)
    .replace('ss', seconds)
}

/**
 * 格式化状态
 * @param status - 状态值
 * @returns 状态对象 {text, type}
 */
export function formatStatus(status: number | boolean): { text: string; type: 'success' | 'danger' | 'warning' | 'info' } {
  if (status === 1 || status === true) {
    return { text: '启用', type: 'success' }
  } else {
    return { text: '禁用', type: 'danger' }
  }
}

