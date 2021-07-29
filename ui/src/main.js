import { createApp } from 'vue'
import App from './App.vue'

let app = createApp(App)

import mixin from './mixins'
app.use(mixin)

import util from './pkg/util'
app.use(util)

import api from './api'
app.use(api)

import components from './components'
app.use(components)

import store from './store/index.js'
app.use(store)

import router from './router.js'
app.use(router)

app.mount('#app')