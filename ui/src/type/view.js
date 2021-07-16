import * as common from './common'
import util from 'pkg/util'
import * as inputType from '@/components/input/type'
import * as typeType from '@/components/types/type'
import {httpInputConfig} from 'type/input'
import {StaticTextConfig} from 'type/types'

export const ViewBLockTypeConfigType = () => {
    try {
        return JSON.stringify(StaticTextConfig())
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
        refresh: 10 * 1000,
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