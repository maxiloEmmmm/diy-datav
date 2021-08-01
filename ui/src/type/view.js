import * as common from './common'
import util from 'pkg/util'
import * as inputType from '@/components/input/type'
import * as typeType from '@/components/types/type'
import {httpInputConfig} from 'type/input'
import {TextConfig} from 'type/types'

// .config 要存数据库 为json字符串
// .config.common.input.*.config 要存数据库 为json字符串

export const ViewBLockTypeConfigType = () => {
    try {
        return TextConfig()
    }catch (e) {
        console.log("get ViewBLockTypeConfigType err", e)
    }
}

export const ViewBLockTypeType = () => {
    return typeType.Text
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
            key: ViewBLockTypeCommonInputItemDefault.key(),
            id: ViewBLockTypeCommonInputItemDefault.id()
        }
    }catch (e) {
        console.log('get ViewBLockTypeCommonInputItem err', e)
    }
}

export const ViewBLockTypeCommonInputItemFilter = {
    key(t) {
        return util.isString(t) && t ? t : ViewBLockTypeCommonInputItemDefault.key()
    },
    refresh(t) {
        return util.isNumber(t) && t > 0 ? t : ViewBLockTypeCommonInputItemDefault.refresh()
    }
}

export const ViewBLockTypeCommonInputItemDefault = {
    key() {
        return util.uuid('input')
    },
    refresh() {
        return 10
    },
    id() {
        return 0
    }
}

export const ViewBLockTypeCommonInputItemParse = (t) => {
    t.key = ViewBLockTypeCommonInputItemFilter.key(t.key)
    // t.refresh = ViewBLockTypeCommonInputItemFilter.refresh(t.refresh)
    return t
}

export const ViewBLockTypeCommonFilter = {
    input(t) {
        return util.isArray(t) ? t : ViewBLockTypeCommonDefault.input()
    },
    zIndex(t) {
        return util.isNumber(t) && t >= 1 ? t : ViewBLockTypeCommonDefault.zIndex()
    }
}

export const ViewBLockTypeCommonDefault = {
    input() {
        return []
    },
    zIndex() {
        return 1
    }
}

export const ViewBLockTypeCommonParse = (t) => {
    let cfg = {
        ...ViewBLockTypeCommon(),
        // TODO: format other field
        ...t
    }
    cfg.input = ViewBLockTypeCommonFilter.input(t.input)
    cfg.zIndex = ViewBLockTypeCommonFilter.zIndex(t.zIndex)
    return cfg
}

export const ViewBLockTypeCommon = () => {
    return {
        position: common.PositionType(),
        input: [],
        refresh: 10,
        zIndex: 1,
    }
}

export const ViewBlockType = function() {
    let blockKey = util.uuid()
    return {
        getKey() {
            return blockKey
        },
        id: 0,
        type: ViewBLockTypeType(),
        config: JSON.stringify(ViewBlockTypeConfig())
    }
}

export const ViewBlockTypeConfig = function() {
    return {
        common: ViewBLockTypeCommon(),
        type: ViewBLockTypeConfigType()
    }
}

export const ViewType = function() {
    return {
        id: 0,
        desc: '',
        config: '',
        bgAssetsID: 0,
        blocks: [],
        newBlock() {
            return ViewBlockType()
        },
        newBlockAndStore() {
            let block = this.newBlock()
            let zIndex = 1
            this.blocks.forEach(block => {
                try {
                    let b = JSON.parse(block.config)
                    b.common.zIndex > zIndex && (zIndex = b.common)
                }catch (e) {
                    console.log('parse block config err', e)
                }
            })

            try {
                let b = JSON.parse(block.config)
                b.common.zIndex = zIndex
                block.config = JSON.stringify(b)
            }catch (e) {
                console.log('parse block config err', e)
            }

            this.blocks.push(block)
            return block
        },
        pushBlock(block) {
            this.blocks.push(block)
        },
    }
}