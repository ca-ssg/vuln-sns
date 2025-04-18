import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { Quasar, Notify } from 'quasar'

// Import Quasar css
import '@quasar/extras/material-icons/material-icons.css'
import '@quasar/extras/fontawesome-v6/fontawesome-v6.css'
import 'quasar/dist/quasar.css'

// Import app css
import './style.css'

import App from './App.vue'
import router from './router'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.use(Quasar, {
  plugins: { Notify },
  config: {
    dark: true,
    brand: {
      primary: '#1DA1F2',
      secondary: '#657786',
      accent: '#F91880',
      dark: '#000000'
    },
    notify: {
      position: 'top',
      timeout: 2500,
      textColor: 'white'
    }
  }
})
app.mount('#app')
