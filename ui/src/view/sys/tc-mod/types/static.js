import util from 'pkg/util'
export const staticTypeConfigFilter = {
    data(t) {
        return util.isString(t) ? t : staticTypeConfigDefault.data()
    },
}

export const staticTypeConfigDefault = {
    data() {
        return '[]'
    },
}

export const staticTypeConfigParse = (t) => {
    let cfg = staticTypeConfig()
    cfg.data = staticTypeConfigFilter.data(t.data)
    return cfg
}

export const staticTypeConfig = () => {
    return {
        data: staticTypeConfigDefault.data(),
    }
}