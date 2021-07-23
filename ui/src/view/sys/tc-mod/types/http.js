import util from 'pkg/util'

export const httpTypeConfigFilter = {
    url(t) {
        return util.isString(t) ? t : httpTypeConfigDefault.url()
    },
}

export const httpTypeConfigDefault = {
    url() {
        return ''
    },
}

export const httpTypeConfigParse = (t) => {
    let cfg = httpTypeConfig()
    cfg.url = httpTypeConfigFilter.url(t.url)
    return cfg
}

export const httpTypeConfig = () => {
    return {
        url: httpTypeConfigDefault.url(),
    }
}