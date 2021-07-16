import * as type from './type'
import mockUtil from 'mockjs'
import http from 'pkg/http'
import { ViewType, ViewBlockType, AntVConfig, StaticTextConfig } from 'type'
import util from 'pkg/util'
import * as componentType from '@/components/types/type.js'
import * as inputType from '@/components/input/type.js'
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
            let block2 = ViewBlockType()
            view.blocks = [block, block2]

            block.type = componentType.AntV
            block2.type = componentType.StaticText

            let staticTextConfig = StaticTextConfig()
            staticTextConfig.text = "test"
            block2.config.type = staticTextConfig
            block2.config = JSON.stringify(block2.config)

            block.config.common.input = [
                {id: 1, type: inputType.Http, config: '{}'}
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
                { year: '1991', value: 0.19 },
                { year: '1992', value: 0.21 },
                { year: '1993', value: 0.27 },
                { year: '1994', value: 0.33 },
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