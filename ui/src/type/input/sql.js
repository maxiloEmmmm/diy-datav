import util from 'pkg/util'
import * as common from "./common";
import {httpInputConfig} from "./http";
export const sqlInputConfigFilter = {
    sql(t) {
        return util.isString(t) ? t : sqlInputConfigDefault.sql()
    },
    engine(t) {
        return util.isNumber(t) ? t : sqlInputConfigDefault.engine()
    },
    ...common.commonInputConfigFilter
}

export const sqlInputConfigDefault = {
    sql() {
        return ''
    },
    engine() {
        return 0
    },
    ...common.commonInputConfigDefault
}

export const sqlInputConfigParse = (t) => {
    let cfg = sqlInputConfig()
    cfg.sql = sqlInputConfigFilter.sql(t.sql)
    cfg.engine = sqlInputConfigFilter.engine(t.engine)
    return {
        ...cfg,
        ...common.commonInputConfigParse(t)
    }
}

export const sqlInputConfig = () => {
    return {
        sql: sqlInputConfigDefault.sql(),
        engine: sqlInputConfigDefault.engine(),
        ...common.commonInputConfig()
    }
}