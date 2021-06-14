import mockUtil from './mock'

import http from 'pkg/http'
import * as type from './type'

http.defaults.baseURL = import.meta.env.VITE_API

export const api = {
    [type.ViewStore](view) {
        return http.put('view', view)
    },
    [type.ViewUploadBG](file) {
        return http.post('view/bg/upload')
    }
}

export default {
    install(app) {
        app.config.globalProperties.$api = api
        app.config.globalProperties.$apiType = type
        if(import.meta.env.VITE_MOCK == 'on') {
            // TODO: repalce api generate
            mockUtil()
        }
    }
}