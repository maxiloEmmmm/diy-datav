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

import { createRouter, createWebHashHistory } from 'vue-router'

app.use(createRouter({
    history: createWebHashHistory(),
    routes: [
        {path: '/view/design', component: () => import('./view/view/design.vue')},
        {path: '/test/antv', component: () => import('./view/test/antv.vue')},
        {path: '/sys/tc', component: () => import('./view/sys/tc.vue')}
    ]
}))

app.mount('#app')