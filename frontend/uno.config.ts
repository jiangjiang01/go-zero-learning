import { defineConfig } from 'unocss'
import presetUno from '@unocss/preset-uno'
import presetAttributify from '@unocss/preset-attributify'
import presetIcons from '@unocss/preset-icons'

export default defineConfig({
  presets: [
    presetUno(), // 默认预设
    presetAttributify(), // 属性模式支持
    presetIcons({
      // 图标预设
      scale: 1.2,
      warn: true
    })
  ],
  shortcuts: {
    // 自定义快捷方式
    'flex-center': 'flex items-center justify-center',
    'flex-between': 'flex items-center justify-between'
  },
  theme: {
    colors: {
      primary: {
        DEFAULT: '#409EFF',
        50: '#ecf5ff',
        100: '#d9ecff',
        200: '#b3d8ff',
        300: '#8cc5ff',
        400: '#66b1ff',
        500: '#409eff',
        600: '#337ecc',
        700: '#265e99',
        800: '#1a3f66',
        900: '#0d1f33'
      }
    }
  },
  rules: [
    // 自定义规则
  ]
})