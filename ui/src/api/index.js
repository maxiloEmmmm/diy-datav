import mockUtil from 'mockjs'

import http from 'pkg/http'
import * as type from './type'

http.baseURL = import.meta.env.VITE_API

import bgMockUrl from '@/assets/bg_design.png'

const MockOn = import.meta.env.VITE_MOCK == 'on'

const api = {
    [type.ViewStore]: {
        handler(view) {
            return http.put('view', view)
        },
        mock() {
            mockUtil.mock('view', 'put', function(request) {
                return {code: 'ok', msg: '', body: request.body}
            })
        }
    },
    [type.ViewUploadBG](file) {
        return http.post('view/bg/upload')
    },
    [type.ViewDownloadBG]: {
        handler(view) {
            return `${http.baseURL}view/${view.bgAssetsId}/bg`
        },
    }
}

export default {
    install(app) {
        app.config.globalProperties.$api = api
        app.config.globalProperties.$apiType = type

        if(MockOn) {
            for (let k in api) {
                let item = api[k]
                item.mock && item.mock()
            }
        }
    }
}