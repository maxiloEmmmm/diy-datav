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

import * as designModel from './view/view/model'
app.use(createRouter({
    history: createWebHashHistory(),
    routes: [
        {path: '/view/design', component: () => import('./view/view/design.vue'), meta: {model: designModel.Design}},
        {path: '/view/design/:id', component: () => import('./view/view/design.vue'), meta: {model: designModel.Design}},
        {path: '/view/show/:id', component: () => import('./view/view/design.vue'), meta: {model: designModel.View}},
        {path: '/test/antv', component: () => import('./view/test/antv.vue')},
        {path: '/sys/tc', component: () => import('./view/sys/tc.vue')},
        {path: '/sys/view', component: () => import('./view/sys/view.vue')},
        {path: '/sys/assets', component: () => import('./view/sys/assets.vue')}
    ]
}))

app.mount('#app')