import util from 'pkg/util'

export const sqlInputConfigFilter = {
    host(t) {
        return util.isString(t) ? t : sqlInputConfigDefault.host()
    },
    port(t) {
        return util.isNumber(t) ? t : sqlInputConfigDefault.port()
    },
    user(t) {
        return util.isString(t) ? t : sqlInputConfigDefault.user()
    },
    pass(t) {
        return util.isString(t) ? t : sqlInputConfigDefault.pass()
    },
    db(t) {
        return util.isString(t) ? t : sqlInputConfigDefault.db()
    }
}

export const sqlInputConfigDefault = {
    host() {
        return 'localhost'
    },
    port() {
        return 3306
    },
    user() {
        return 'ReadOnlyUser'
    },
    pass() {
        return ''
    },
    db() {
        return ''
    }
}

export const sqlInputConfigParse = (t) => {
    let cfg = sqlInputConfig()
    cfg.host = sqlInputConfigFilter.host(t.host)
    cfg.port = sqlInputConfigFilter.port(t.port)
    cfg.user = sqlInputConfigFilter.user(t.user)
    cfg.pass = sqlInputConfigFilter.pass(t.pass)
    cfg.db = sqlInputConfigFilter.db(t.db)
    return cfg
}

export const sqlInputConfig = () => {
    return {
        host: sqlInputConfigDefault.host(),
        port: sqlInputConfigDefault.port(),
        user: sqlInputConfigDefault.user(),
        pass: sqlInputConfigDefault.pass(),
        db: sqlInputConfigDefault.db(),
    }
}