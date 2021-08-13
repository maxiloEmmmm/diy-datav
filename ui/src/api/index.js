import mockUtil from './mock'

import http from 'pkg/http'
import * as apiType from './type'
import antdTool from 'pkg/antd-tool'
import * as viewStore from '@/store/view'
import store from '@/store'
import router from '@/router'
import {notification} from "ant-design-vue"

let rp = null
const refreshURL = "auth/refresh_token"
const loginURL = "auth/login"
http.interceptors.request.use(config => {
    if (config.url === refreshURL) {
        return config
    }
    if(rp) {
        return new Promise(ok => {
            rp.then(() => ok(config))
        })
    }else {
        let token = store.state.auth.token
        if(token && config.url !== loginURL) {
            try {
                let exp = JSON.parse(window.atob(token.split(".")[1])).exp
                let cur = Date.parse(new Date) / 1000
                let range = exp - cur
                if(range <= 0) {
                    store.commit("auth/SetToken", "")
                    router.push("/auth/login")
                    return config
                }else if(exp - cur <= (60 * 60 * 24 * 3)) {
                    rp = new Promise((ok) => {
                        api[type.AuthRefresh]().then(response => {
                            store.commit("auth/SetToken", response.data.token)
                            rp = null
                            ok()
                        })
                    })
                    return new Promise(ok => {
                        rp.then(() => ok(config))
                    })
                }
            } catch (error) {
                console.log("token.filter", error)
            }
        }
        return config
    }
})
http.interceptors.request.use(config => {
    let token = store.state.auth.token
    if(token) {
        config.headers.Authorization = `Bearer ${token}`
    }

    return config
})
http.interceptors.response.use(response => response, (error) => {
    if(!error.response) {
        notification.error({message: error.message, duration: 10})
        return
    }

    switch(error.response.status) {
        case 422:
            if('auth:token:parse' === error.response.data.code) {
                store.commit("auth/SetToken", "")
                router.push("/auth/login")
                return
            }
            notification.error({message: `存在错误 - 代号: ${error.response.data.code}`, description: error.response.data.message, duration: 10})
            break
        case 500:
            if(error.response.data.code) {
                notification.error({message: `系统错误 - 代号: ${error.response.data.code}`, description: error.response.data.message, duration: 10})
                break
            }
        default:
            notification.error({message: `未知错误 代号: ${error.response.status}`, description: error.message, duration: 10})
    }
})

export const type = apiType

http.defaults.baseURL = /^http/.test(import.meta.env.VITE_API)
    ? import.meta.env.VITE_API
    : `${window.location.protocol}//${window.location.host}${/^\//.test(import.meta.env.VITE_API) ? import.meta.env.VITE_API : `/${import.meta.env.VITE_API}`}`

export const api = {
    [apiType.AuthLogin](data) {
        return http.post(loginURL, data)
    },
    [apiType.AuthRefresh]() {
        return http.get(refreshURL)
    },
    [apiType.ViewInfo](id, type) {
        return http.get(`view/${id}?type=${type}`)
    },
    [apiType.ViewList]() {
        return http.get(`view`)
    },
    [apiType.ViewStore](view) {
        return http.put('view', view)
    },
    [apiType.ViewUploadBG](type, file) {
        return http.post(`view/bg/upload?type=${type}`, file, {'Content-Type': 'multipart/form-data'})
    },
    [apiType.ViewBGAssets]() {
        return http.get('view-bg-assets')
    },
    [apiType.AssetsTypeList]() {
        return http.get('assets-type-assets')
    },
    [apiType.Data](id) {
        return http.get(`data/${id}`)
    },
    [apiType.TmpEchoData](config) {
        return http.post(`data-tmp-echo`, config)
    },
    [apiType.StaticList]() {
        return http.get(`tc/kind/static`)
    },
    [apiType.HttpList]() {
        return http.get(`tc/kind/http`)
    },
    [apiType.SqlList]() {
        return http.get(`tc/kind/sql`)
    },
    [apiType.KindList]() {
        return http.get(`tc/kind`)
    },
    [apiType.MenuList]() {
        return http.get(`menu?page_size=10000`)
    },
    [apiType.MenuAdd](data) {
        return http.post(`menu`, data)
    },
    [apiType.MenuDelete](id) {
        return http.delete(`menu/${id}`)
    },
    [apiType.MenuUpdate](id, data) {
        return http.patch(`menu/${id}`, data)
    },
    [apiType.AuthMenu]() {
        return http.get(`auth/menu`)
    },
    [apiType.PermissionView](sub) {
        return http.get(`permission/view?sub=${sub}`)
    },
    [apiType.PermissionMenu](sub) {
        return http.get(`permission/menu?sub=${sub}`)
    },
    [apiType.PermissionApi](sub) {
        return http.get(`permission/api?sub=${sub}`)
    },
    [apiType.UserList]() {
        // TODO: all if page_size unset or zero
        return http.get(`user?page_size=10000`)
    },
    [apiType.RoleUser](role) {
        return http.get(`/casbin/role/${role}/user`)
    },
    [apiType.AddUserForRole](role, user) {
        return http.post(`/casbin/role/${role}/user/${user}`)
    },
    [apiType.RemoveUserForRole](role, user) {
        return http.delete(`/casbin/role/${role}/user/${user}`)
    },
    [apiType.Policy](rules) {
        return http.post(`/permission/change`, rules)
    },
}

export default {
    install(app) {
        antdTool.http.engine = http
        app.config.globalProperties.$api = api
        app.config.globalProperties.$api_url = http.defaults.baseURL
        app.config.globalProperties.$apiType = apiType
        if(import.meta.env.VITE_MOCK == 'on') {
            // TODO: repalce api generate
            mockUtil()
        }
    }
}

viewStore.SetFetchEngine(api[apiType.Data])
viewStore.SetFetchTmpEchoEngine(api[apiType.TmpEchoData])