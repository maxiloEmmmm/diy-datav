import * as common from './common'

export const ViewBlockType = function() {
    return {
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
        }
    }
}