/// <reference types="vite/client" />

// Fix for import.meta.env
interface ImportMeta {
  readonly env: {
    [key: string]: string;
    VITE_API_URL: string;
  };
}

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<Record<string, unknown>, Record<string, unknown>, unknown>
  export default component
}

// Declare modules for which TypeScript can't find type declarations
declare module 'pinia' {
  export function defineStore(
    id: string, 
    setup: () => {
      [key: string]: any;
    }
  ): any;
}

declare module 'vue' {
  export function ref<T>(value: T): { value: T };
  export function computed<T>(getter: () => T): { value: T };
}

declare module 'axios' {
  interface AxiosRequestConfig {
    baseURL?: string;
    withCredentials?: boolean;
    headers?: Record<string, string>;
    [key: string]: any;
  }

  interface AxiosInstance {
    post: (url: string, data?: any) => Promise<any>;
    put: (url: string, data?: any) => Promise<any>;
    interceptors: {
      request: {
        use: (fn: (config: any) => any) => void;
      };
    };
  }

  const axios: {
    create: (config: AxiosRequestConfig) => AxiosInstance;
  };

  export default axios;
}
