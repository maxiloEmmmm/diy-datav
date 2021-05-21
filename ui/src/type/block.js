import * as common from './common'

export const ViewBlockType = function() {
    return {
        postion: common.PositionType(),
        config: '',
        type: '',
        dataset: [],
        refresh: 10 * 1000,
    }
}