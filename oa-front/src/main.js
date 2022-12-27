import { createApp } from 'vue'

import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import App from './App.vue'
import router from './router'
import pinia from './pinia'

import directives from './directives/index.js'

// import './assets/main.css'

const app = createApp(App)

// app.use(router)
// app.use(ElementPlus, {
//   locale: zhCn,
// })

app
  .use(router)
  .use(pinia)
  .use(ElementPlus, {
    locale: zhCn,
  })

directives(app)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.mount('#app')
