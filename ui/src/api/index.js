import mockUtil from './mock'

import http from 'pkg/http'
import * as apiType from './type'

export const type = apiType

http.defaults.baseURL = import.meta.env.VITE_API

export const api = {
    [apiType.ViewInfo](id) {
        return http.get(`view/${id}`)
    },
    [apiType.ViewStore](view) {
        return http.put('view', view)
    },
    [apiType.ViewUploadBG](file) {
        return http.post('view/bg/upload')
    },
    [apiType.Data](id) {
        return http.get(`data/${id}`)
    }
}

export default {
    install(app) {
        app.config.globalProperties.$api = api
        app.config.globalProperties.$apiType = apiType
        if(import.meta.env.VITE_MOCK == 'on') {
            // TODO: repalce api generate
            mockUtil()
        }
    }
}