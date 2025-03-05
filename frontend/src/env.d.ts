/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<Record<string, unknown>, Record<string, unknown>, unknown>
  export default component
}

// Add interface for import.meta.env
interface ImportMetaEnv {
  readonly VITE_API_URL: string
  // more env variables...
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}

// ESLint configuration modules
declare module '@vue/eslint-config-typescript' {
  export const defineConfigWithVueTs: any;
  export const vueTsConfigs: any;
}

declare module '@vue/eslint-config-prettier/skip-formatting' {
  const skipFormatting: any;
  export default skipFormatting;
}
