declare module 'pinia' {
  import { Component, ComputedRef, Ref } from 'vue'

  export interface StoreOptions {
    state?: () => any
    getters?: Record<string, any>
    actions?: Record<string, any>
  }

  export function defineStore(
    id: string,
    options: StoreOptions | (() => any)
  ): any

  export function createPinia(): any
  
  export function mapStores(...stores: any[]): any
  export function mapState(store: any, map: string[] | Record<string, string | (() => any)>): any
  export function mapGetters(store: any, map: string[] | Record<string, string>): any
  export function mapActions(store: any, map: string[] | Record<string, string>): any
}
