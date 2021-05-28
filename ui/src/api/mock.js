import * as type from './type'
import mockUtil from 'mockjs'
import http from 'pkg/http'

function mock() {
    arguments[0] = `${http.defaults.baseURL.replace(/\/$/, '')}/${arguments[0].replace(/^\//, '')}`
    return mockUtil.mock(...arguments)
}

export default function() {
    return {
        [type.ViewStore]: mock('view', 'put', function(request) {
            return {code: 'ok', msg: '', body: request.body}
        }),
        [type.ViewUploadBG]: mock('view/bg/upload', 'post', function(request) {
            return {code: 'ok', msg: '', body: request.body}
        })
    }
}