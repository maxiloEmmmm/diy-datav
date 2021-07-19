import * as common from './common'
import util from 'pkg/util'
import * as inputType from '@/components/input/type'
import * as typeType from '@/components/types/type'
import {httpInputConfig} from 'type/input'
import {StaticTextConfig} from 'type/types'

// .config 要存数据库 为json字符串
// .config.common.input.*.config 要存数据库 为json字符串

export const ViewBLockTypeConfigType = () => {
    try {
        return StaticTextConfig()
    }catch (e) {
        console.log("get ViewBLockTypeConfigType err", e)
    }
}

export const ViewBLockTypeType = () => {
    return typeType.StaticText
}

export const ViewBLockTypeCommonInputItemType = () => {
    return inputType.Http
}

export const ViewBLockTypeCommonInputItemConfig = () => {
    return JSON.stringify(httpInputConfig())
}

export const ViewBLockTypeCommonInputItem = () => {
    try {
        return {
            config: ViewBLockTypeCommonInputItemConfig(),
            type: ViewBLockTypeCommonInputItemType(),
            title: '',
            id: ''
        }
    }catch (e) {
        console.log('get ViewBLockTypeCommonInputItem err', e)
    }
}

export const ViewBLockTypeCommon = () => {
    return {
        position: common.PositionType(),
        input: [],
        refresh: 10,
    }
}

export const ViewBlockType = function() {
    let blockKey = util.uuid()
    return {
        getKey() {
            return blockKey
        },
        id: '',
        type: ViewBLockTypeType(),
        config: {
            common: ViewBLockTypeCommon(),
            type: ViewBLockTypeConfigType()
        },
    }
}

export const ViewType = function() {
    return {
        id: '',
        desc: '',
        config: '',
        bgAssetsId: '',
        blocks: [],
        newBlock() {
            return ViewBlockType()
        },
        newBlockAndStore() {
            let block = this.newBlock()
            this.blocks.push(block)
            return block
        },
        pushBlock(block) {
            this.blocks.push(block)
        },
    }
}