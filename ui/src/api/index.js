import mock from './mock'

import http from 'pkg/http'
import * as type from './type'

http.baseURL = import.meta.env.VITE_API

const api = {
    [type.ViewStore](view) {
        return http.put('view', view)
    },
    [type.ViewUploadBG](file) {
        return http.post('view/bg/upload')
    },
    [type.ViewDownloadBG](view) {
        return `${http.baseURL}view/${view.bgAssetsId}/bg`
    }
}

export default {
    install(app) {
        app.config.globalProperties.$api = api
        app.config.globalProperties.$apiType = type

        if(import.meta.env.VITE_MOCK == 'on') {
            mock()
        }
    }
}