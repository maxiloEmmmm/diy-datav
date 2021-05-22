import mock from './mock'

import http from 'pkg/http'

http.interceptors.request.use(function(config) {
    // TODO: hook mock
    return config
})

const api = {
    view: {
        store(view) {
            return http.put('view', view)
        },
        uploadBG(file) {
            return http.post('view/bg/upload')
        },
        downloadBG(view) {
            return `${http.baseURL}view/${view.bgAssetsId}/bg`
        }
    }
}

export default {
    install(app) {
        app.config.globalProperties.$api = api
    }
}