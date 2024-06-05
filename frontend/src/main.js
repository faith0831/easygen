import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import hljs from 'highlight.js/lib/core'
import javascript from 'highlight.js/lib/languages/javascript'
import csharp from 'highlight.js/lib/languages/csharp'
import go from 'highlight.js/lib/languages/go'
import hljsVuePlugin from '@highlightjs/vue-plugin'
import 'highlight.js/styles/vs.css'

hljs.registerLanguage('javascript', javascript)
hljs.registerLanguage('csharp', csharp)
hljs.registerLanguage('go', go)

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.use(ElementPlus)
app.use(hljsVuePlugin)
app.mount('#app')
