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
    scrollPrecentOnce(t) {
        return util.isNumber(t) && t <= 100 && t >= 0  ? t : TableConfigDefault.scrollPrecentOnce()
    },
    scrollCycle(t) {
        return util.isNumber(t) && t <= 100 && t >= 1 ? t : TableConfigDefault.scrollCycle()
    },
    headerBGC(t) {
        return util.isString(t) && t ? t : TableConfigDefault.headerBGC()
    },
    headerC(t) {
        return util.isString(t) && t ? t : TableConfigDefault.headerC()
    }
}

export const TableConfigDefault = {
    dataIndex() {
        return -1
    },
    scrollPrecentOnce() {
        return 10
    },
    scrollCycle() {
        return 10
    },
    headerBGC() {
        return '#cecece'
    },
    headerC() {
        return '#000'
    }
}

export const TableConfigParse = (config) => {
    let cfg = TableConfig()
    cfg.dataIndex = TableConfigFilter.dataIndex(config.dataIndex)
    cfg.scrollPrecentOnce = TableConfigFilter.scrollPrecentOnce(config.scrollPrecentOnce)
    cfg.scrollCycle = TableConfigFilter.scrollCycle(config.scrollCycle)
    cfg.headerBGC = TableConfigFilter.headerBGC(config.headerBGC)
    cfg.headerC = TableConfigFilter.headerC(config.headerC)
    return cfg
}

export const TableConfig = function(){
    return {
        dataIndex: TableConfigDefault.dataIndex(),
        scrollPrecentOnce: TableConfigDefault.scrollPrecentOnce(),
        scrollCycle: TableConfigDefault.scrollCycle(),
        headerBGC: TableConfigDefault.headerBGC(),
        headerC: TableConfigDefault.headerC()
    }
}