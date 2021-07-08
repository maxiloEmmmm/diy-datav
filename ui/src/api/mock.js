import * as type from './type'
import mockUtil from 'mockjs'
import http from 'pkg/http'
import { ViewType, ViewBlockType, AntVConfig } from 'type'
import util from 'pkg/util'
import * as componentType from '@/components/types/type.js'
import * as apiType from "./type";

function mock() {
    arguments[0] = `${http.defaults.baseURL.replace(/\/$/, '')}/${arguments[0].replace(/^\//, '')}`
    return mockUtil.mock(...arguments)
}

function mockReg() {
    arguments[0] = new RegExp(`^${http.defaults.baseURL.replace(/\/$/, '')}/${arguments[0].replace(/^\//, '')}`)
    return mockUtil.mock(...arguments)
}

export default function() {
    return {
        [type.ViewInfo]: mockReg('view/[^/]+$', 'get', function(request) {
            let view = ViewType()
            let block = ViewBlockType()
            view.blocks = [block]

            block.type = componentType.AntV
            block.config.common.input = [
                {id: 1}
            ]

            let antVConfig = AntVConfig()
            antVConfig.layers = [{type: 'line'}]
            antVConfig.scale.x.field = 'year'
            antVConfig.scale.y.field = 'value'
            block.config.type = antVConfig
            block.config = JSON.stringify(block.config)
            return {code: 'ok', msg: '', data: view}
        }),
        [apiType.Data]: mockReg('data/[^/]+$', 'get', function(request) {
            return {code: 'ok', msg: '', data: [
                { year: '1991', value: 3 },
                { year: '1992', value: 4 },
                { year: '1993', value: 3.5 },
                { year: '1994', value: 5 },
                { year: '1995', value: 4.9 },
                { year: '1996', value: 6 },
                { year: '1997', value: 7 },
                { year: '1998', value: 9 },
                { year: '1999', value: 13 },
            ]}
        }),
        [type.ViewStore]: mock('view', 'put', function(request) {
            return {code: 'ok', msg: '', data: request.body}
        }),
        [type.ViewUploadBG]: mock('view/bg/upload', 'post', function(request) {
            return {code: 'ok', msg: '', data: request.body}
        })
    }
}