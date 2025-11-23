/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

declare module 'nprogress' {
  interface NProgress {
    start(): void
    done(): void
    configure(options: { showSpinner: boolean }): void
  }
  const nprogress: NProgress
  export default nprogress
}

interface ImportMetaEnv {
  readonly VITE_APP_TITLE: string
  readonly VITE_APP_BASE_API: string
  readonly VITE_APP_BASE_URL: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}