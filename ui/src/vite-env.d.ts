/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_API_ROOT?: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
