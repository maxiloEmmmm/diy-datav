import util from 'pkg/util'

export const mysqlType = 'mysql'
export const sqlite3Type = 'sqlite3'
export const postgresType = 'postgres'
export const gremlinType = 'gremlin'

export const sqlTypeOption = [mysqlType, sqlite3Type, postgresType, gremlinType]

export const sqlTypeConfigFilter = {
    host(t) {
        return util.isString(t) ? t : sqlTypeConfigDefault.host()
    },
    port(type, t) {
        return util.isNumber(t) ? t : sqlTypeConfigDefault.port(type)
    },
    user(t) {
        return util.isString(t) ? t : sqlTypeConfigDefault.user()
    },
    pass(t) {
        return util.isString(t) ? t : sqlTypeConfigDefault.pass()
    },
    db(t) {
        return util.isString(t) ? t : sqlTypeConfigDefault.db()
    },
    type(t) {
        return [mysqlType, sqlite3Type, postgresType, gremlinType].includes(t) ? t : sqlTypeConfigDefault.type()
    }
}

export const sqlTypeConfigDefault = {
    host() {
        return 'localhost'
    },
    port(type) {
        switch (type) {
        case sqlite3Type:
            return null
        case postgresType:
            return 5432
        case gremlinType:
            return null
        case mysqlType:
        default:
            return 3306
        }
    },
    user() {
        return 'readOnlyUser'
    },
    pass() {
        return ''
    },
    db() {
        return ''
    },
    type() {
        return mysqlType
    }
}

export const sqlTypeConfigParse = (t) => {
    let cfg = sqlTypeConfig()
    cfg.host = sqlTypeConfigFilter.host(t.host)
    cfg.port = sqlTypeConfigFilter.port(t.type, t.port)
    cfg.user = sqlTypeConfigFilter.user(t.user)
    cfg.pass = sqlTypeConfigFilter.pass(t.pass)
    cfg.db = sqlTypeConfigFilter.db(t.db)
    cfg.type = sqlTypeConfigFilter.type(t.type)
    return cfg
}

export const sqlTypeConfig = () => {
    return {
        host: sqlTypeConfigDefault.host(),
        port: sqlTypeConfigDefault.port(sqlTypeConfigDefault.type()),
        user: sqlTypeConfigDefault.user(),
        pass: sqlTypeConfigDefault.pass(),
        db: sqlTypeConfigDefault.db(),
        type: sqlTypeConfigDefault.type(),
    }
}