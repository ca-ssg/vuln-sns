declare module 'pinia';
declare module 'vue';
declare module 'axios';

interface ImportMeta {
  env: {
    VITE_API_URL: string;
    [key: string]: any;
  };
}
