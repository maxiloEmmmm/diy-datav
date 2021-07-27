import util from 'pkg/util'
import * as common from "./common";
import {httpInputConfig} from "./http";
export const sqlInputConfigFilter = {
    ...common.commonInputConfigFilter,
    sql(t) {
        return util.isString(t) ? t : sqlInputConfigDefault.sql()
    },
    engine(t) {
        return util.isNumber(t) ? t : sqlInputConfigDefault.engine()
    },
}

export const sqlInputConfigDefault = {
    ...common.commonInputConfigDefault,
    sql() {
        return ''
    },
    engine() {
        return 0
    },
}

export const sqlInputConfigParse = (t) => {
    let cfg = sqlInputConfig()
    cfg.sql = sqlInputConfigFilter.sql(t.sql)
    cfg.engine = sqlInputConfigFilter.engine(t.engine)
    return {
        ...common.commonInputConfigParse(t),
        ...cfg,
    }
}

export const sqlInputConfig = () => {
    return {
        ...common.commonInputConfig(),
        sql: sqlInputConfigDefault.sql(),
        engine: sqlInputConfigDefault.engine(),
    }
}