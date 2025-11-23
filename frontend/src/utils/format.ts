/**
 * 格式化日期时间
 * @param date - 日期
 * @param format - 格式，默认 'YYYY-MM-DD HH:mm:ss'
 * @returns 格式化后的日期字符串
 */
export function formatDateTime(date: string | Date | null | undefined, format = 'YYYY-MM-DD HH:mm:ss'): string {
  if (!date) return ''
  
  const d = new Date(date)
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

