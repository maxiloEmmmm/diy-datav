import mockUtil from './mock'

import http from 'pkg/http'
import * as apiType from './type'
import antdTool from 'pkg/antd-tool'

export const type = apiType

http.defaults.baseURL = /^http/.test(import.meta.env.VITE_API)
    ? import.meta.env.VITE_API
    : `${window.location.protocol}//${window.location.host}${/^\//.test(import.meta.env.VITE_API) ? import.meta.env.VITE_API : `/${import.meta.env.VITE_API}`}`

export const api = {
    [apiType.ViewInfo](id) {
        return http.get(`view/${id}`)
    },
    [apiType.ViewList]() {
        return http.get(`view`)
    },
    [apiType.ViewStore](view) {
        return http.put('view', view)
    },
    [apiType.ViewUploadBG](file) {
        return http.post('view/bg/upload')
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
    }
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