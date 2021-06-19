import * as common from './common'
import util from 'pkg/util'
export const ViewBlockType = function() {
    let blockKey = util.uuid()
    return {
        getKey() {
            return blockKey
        },
        id: '',
        type: '',
        config: {
            common: {
                postion: common.PositionType(),
                input: [],
                refresh: 10 * 1000,
            },
            type: ''
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