import type { FormRules } from 'element-plus'

/**
 * 验证用户名
 */
export const validateUsername = (rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error('请输入用户名'))
  } else if (value.length < 3 || value.length > 20) {
    callback(new Error('用户名长度在 3 到 20 个字符'))
  } else {
    callback()
  }
}

/**
 * 验证密码
 */
export const validatePassword = (rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error('请输入密码'))
  } else if (value.length < 6 || value.length > 20) {
    callback(new Error('密码长度在 6 到 20 个字符'))
  } else {
    callback()
  }
}

/**
 * 验证邮箱
 */
export const validateEmail = (rule: any, value: string, callback: any) => {
  const emailReg = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
  if (!value) {
    callback(new Error('请输入邮箱'))
  } else if (!emailReg.test(value)) {
    callback(new Error('请输入正确的邮箱格式'))
  } else {
    callback()
  }
}

/**
 * 验证手机号
 */
export const validatePhone = (rule: any, value: string, callback: any) => {
  const phoneReg = /^1[3-9]\d{9}$/
  if (!value) {
    callback(new Error('请输入手机号'))
  } else if (!phoneReg.test(value)) {
    callback(new Error('请输入正确的手机号格式'))
  } else {
    callback()
  }
}

