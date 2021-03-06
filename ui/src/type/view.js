import * as common from './common'
import util from 'pkg/util'
import * as inputType from '@/components/input/type'
import * as typeType from '@/components/types/type'
import {httpInputConfig} from 'type/input'
import {TextConfig, TextConfigParse} from 'type/types'
import {PositionType, PositionTypeParse} from "./common";

export const NormalDesignMode = "normal"
export const LayoutDesignMode = "layout"

export const miniNormalBlockIndex = 20
export const miniLayoutBlockIndex = 1
export const maxiLayoutBlockIndex = 19

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
    return t
}

export const ViewBLockTypeCommonFilter = {
    input(t) {
        return util.isArray(t) ? t : ViewBLockTypeCommonDefault.input()
    },
    zIndex(typ, t) {
        if(!util.isNumber(t)) {
            return ViewBLockTypeCommonDefault.zIndex(typ)
        }
        switch (typ) {
        case typeType.Grid:
            if(t > maxiLayoutBlockIndex) {
                t = maxiLayoutBlockIndex
            }
            break
        default:
            if(t < miniLayoutBlockIndex) {
                t = miniNormalBlockIndex
            }
        }
        return t
    },
    refresh(t) {
        return util.isNumber(t) && t > 0 ? t : ViewBLockTypeCommonDefault.refresh()
    },
    bg(t) {
        return util.isString(t) ? t : ViewBLockTypeCommonDefault.bg()
    },
    desc(t) {
        let cfg = ViewBLockTypeCommonDefault.desc()
        if(!util.isObject(t)) {
            return cfg
        }

        cfg.textConfig = ViewBLockTypeCommonFilter.descTextConfig(t.textConfig)
        cfg.positionType = ViewBLockTypeCommonFilter.descPositionType(t.positionType)
        cfg.position = ViewBLockTypeCommonFilter.descPosition(t.position)
        return cfg
    },
    descTextConfig(t) {
        return TextConfigParse(t)
    },
    descPosition(t) {
        return common.PositionTypeParse(t)
    },
    descPositionType(t) {
        return descPosition.includes(t) ? t : ViewBLockTypeCommonDefault.descPositionType()
    },
    border(t) {
        let cfg = ViewBLockTypeCommonDefault.border()
        if(!util.isObject(t)) {
            return cfg
        }
        cfg.width = ViewBLockTypeCommonFilter.borderWidth(t.width)
        cfg.color = ViewBLockTypeCommonFilter.borderColor(t.color)
        cfg.style = ViewBLockTypeCommonFilter.borderStyle(t.style)
        return cfg
    },
    borderWidth(t) {
        return util.isNumber(t) && t >= 0 ? t : ViewBLockTypeCommonDefault.borderWidth()
    },
    borderColor(t) {
        return util.isString(t) && t ? t : ViewBLockTypeCommonDefault.borderColor()
    },
    borderStyle(t) {
        return borderStyle.includes(t) ? t : ViewBLockTypeCommonDefault.borderStyle()
    },
    grid(t) {
        return util.isString(t) ? t : ViewBLockTypeCommonDefault.grid()
    }
}

export const descPosition = ['top', 'bottom']
export const borderStyle = ['solid', 'dashed', 'dotted', 'double', 'groove', 'inset', 'outset', 'ridge']
export const ViewBLockTypeCommonDefault = {
    input() {
        return []
    },
    zIndex(typ) {
        switch (typ) {
        case typeType.Grid:
            return miniLayoutBlockIndex
        default:
            return miniNormalBlockIndex
        }
    },
    position() {
        return common.PositionType()
    },
    refresh() {
        return 10
    },
    bg() {
        // transparent
        return ''
    },
    desc() {
        return {
            textConfig: ViewBLockTypeCommonDefault.descTextConfig(),
            positionType: ViewBLockTypeCommonDefault.descPositionType(),
            position: ViewBLockTypeCommonDefault.descPosition()
        }
    },
    descPosition() {
        return common.PositionType()
    },
    descPositionType() {
        return descPosition[0]
    },
    descTextConfig() {
        return TextConfig()
    },
    border() {
        return {
            width: ViewBLockTypeCommonDefault.borderWidth(),
            color: ViewBLockTypeCommonDefault.borderColor(),
            style: ViewBLockTypeCommonDefault.borderStyle()
        }
    },
    borderWidth() {
        return 0.1
    },
    borderColor() {
        return '#000'
    },
    borderStyle() {
        return borderStyle[0]
    },
    grid() {
        return ''
    }
}

export const ViewBLockTypeCommonParse = (type, t) => {
    console.log('cg pre', t.grid)
    let cfg = {
        ...ViewBLockTypeCommon(),
        ...t
    }
    cfg.position = common.PositionTypeParse(t.position)
    cfg.input = ViewBLockTypeCommonFilter.input(t.input)
    cfg.zIndex = ViewBLockTypeCommonFilter.zIndex(type, t.zIndex)
    cfg.refresh = ViewBLockTypeCommonFilter.refresh(t.refresh)
    cfg.desc = ViewBLockTypeCommonFilter.desc(t.desc)
    cfg.bg = ViewBLockTypeCommonFilter.bg(t.bg)
    cfg.border= ViewBLockTypeCommonFilter.border(t.border)
    console.log('cg', t.grid)
    cfg.grid = ViewBLockTypeCommonFilter.grid(t.grid)
    return cfg
}

export const ViewBLockTypeCommon = () => {
    return {
        position: common.PositionType(),
        input: [],
        refresh: ViewBLockTypeCommonDefault.refresh(),
        zIndex: ViewBLockTypeCommonDefault.zIndex(),
        desc: ViewBLockTypeCommonDefault.desc(),
        border: ViewBLockTypeCommonDefault.border(),
        grid: ViewBLockTypeCommonDefault.grid()
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
        config: ViewBlockTypeConfig()
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
            let zIndex = miniNormalBlockIndex
            this.blocks.forEach(block => {
                try {
                    block.config.common.zIndex > zIndex && (zIndex = block.config.common.zIndex)
                }catch (e) {
                    console.log('parse block config err', e)
                }
            })

            block.config.common.zIndex = zIndex
            this.blocks.push(block)
            return block
        },
        pushBlock(block) {
            this.blocks.push(block)
        },
    }
}