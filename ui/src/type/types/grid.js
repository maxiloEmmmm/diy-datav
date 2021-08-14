import util from 'pkg/util'
import {PaddingParse, Padding} from '../common.js'
export const GridConfigFilter = {
    rows(t) {
        if(!util.isArray(t)) {
            return GridConfigDefault.rows()
        }

        return t.map(GridConfigFilter.row)
    },
    row(t) {
        if(!util.isObject(t)) {
            return GridConfigDefault.row()
        }

        return {
            height: GridConfigFilter.rowHeight(t.height),
            rowCols: GridConfigFilter.rowCols(t.rowCols),
            padding: GridConfigFilter.rowPadding(t.padding)
        }
    },
    rowPadding: PaddingParse,
    rowColPadding: PaddingParse,
    rowHeight(t) {
        return !util.isNumber(t) || t > 1 || t < 0 ? GridConfigDefault.rowHeight() : t
    },
    rowCols(t) {
        if(!util.isArray(t)) {
            return GridConfigDefault.rowCols()
        }

        return t.map(GridConfigFilter.rowCol)
    },
    rowCol(t) {
        if(!util.isObject(t)) {
            return GridConfigDefault.rowCol()
        }

        return {
            width: GridConfigFilter.rowColWidth(t.width),
            padding: GridConfigFilter.rowColPadding(t.padding)
        }
    },
    rowColWidth(t) {
        return !util.isNumber(t) || t > 1 || t < 0 ? GridConfigDefault.rowColWidth() : t
    }
}

export const GridConfigDefault = {
    rows() {
        return []
    },
    row() {
        return {
            height: GridConfigDefault.rowHeight(),
            padding: GridConfigDefault.rowPadding(),
            rowCols: GridConfigDefault.rowCols()
        }
    },
    rowPadding: Padding,
    rowColPadding: Padding,
    rowHeight() {
        return 1
    },
    rowCols() {
        return []
    },
    rowCol() {
        return {
            width: GridConfigDefault.rowColWidth(),
            padding: GridConfigDefault.rowColPadding(),
        }
    },
    rowColWidth() {
        return 1
    }
}

export const GridConfigParse = (config) => {
    let cfg = GridConfig()
    cfg.rows = GridConfigFilter.rows(config.rows)
    return cfg
}

export const GridConfig = function(){
    return {
        rows: GridConfigDefault.rows()
    }
}