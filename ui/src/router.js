import {createRouter, createWebHashHistory} from "vue-router";
import * as designModel from "./view/view/model";
import store from '@/store'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {path: '/view/design', component: () => import('./view/view/design.vue'), meta: {model: designModel.Design}},
        {path: '/view/design/:id', component: () => import('./view/view/design.vue'), meta: {model: designModel.Design}},
        {path: '/view/show/:id', component: () => import('./view/view/design.vue'), meta: {model: designModel.View}},
        {path: '/test/antv', component: () => import('./view/test/antv.vue')},
        {path: '/sys/tc', component: () => import('./view/sys/tc.vue')},
        {path: '/sys/user', component: () => import('./view/sys/user.vue')},
        {path: '/sys/view', component: () => import('./view/sys/view.vue')},
        {path: '/', component: () => import('./view/sys/index.vue')},
        {path: '/sys/assets', component: () => import('./view/sys/assets.vue')},
        {path: '/sys/role', component: () => import('./view/sys/role.vue')},
        {path: '/sys/menu', component: () => import('./view/sys/menu.vue')},
        {path: '/auth/login', component: () => import('./view/sys/login.vue'), meta: {auth: false}}
    ]
})

router.beforeEach(function(to, from, next) {
    if(to.meta.auth === undefined || to.meta.auth) {
        let hasToken = store.state.auth.token != ""
        if(!hasToken) {
            store.commit("auth/SetForward", window.location.hash)
            next("/auth/login")
            return
        }
    }

    next()
})

export default router