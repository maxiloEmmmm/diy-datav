import util from 'pkg/util'
import {AntVConfigDefault} from "./antv";

export const TableConfigFilter = {
    dataIndex(t) {
        let index = parseInt(t)
        if(!util.isNumber(index) || index < -1) {
            return TableConfigDefault.dataIndex()
        }
        return index
    },
}

export const TableConfigDefault = {
    dataIndex() {
        return -1
    },
}

export const TableConfigParse = (config) => {
    let cfg = TableConfig()
    cfg.dataIndex = TableConfigFilter.dataIndex(config.dataIndex)
    return cfg
}

export const TableConfig = function(){
    return {
        dataIndex: TableConfigDefault.dataIndex()
    }
}