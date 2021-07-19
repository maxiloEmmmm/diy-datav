import util from 'pkg/util'

export const sqlInputConfigFilter = {
    sql(t) {
        return util.isString(t) ? t : sqlInputConfigDefault.sql()
    },
    engine(t) {
        return util.isNumber(t) ? t : sqlInputConfigDefault.engine()
    },
}

export const sqlInputConfigDefault = {
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
    return cfg
}

export const sqlInputConfig = () => {
    return {
        sql: sqlInputConfigDefault.sql(),
        engine: sqlInputConfigDefault.engine(),
    }
}