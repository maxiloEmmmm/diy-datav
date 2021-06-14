import { createApp } from 'vue'
import App from './App.vue'

let app = createApp(App)

import api from './api'
app.use(api)

import components from './components'
app.use(components)

import store from './store'
app.use(store)

import 'ant-design-vue/dist/antd.css'
import antd from 'ant-design-vue'
app.use(antd)

import { createRouter, createWebHashHistory } from 'vue-router'

app.use(createRouter({
    history: createWebHashHistory(),
    routes: [
        {path: '/view/design', component: () => import('./view/view/design.vue')}
    ]
}))

app.mount('#app')