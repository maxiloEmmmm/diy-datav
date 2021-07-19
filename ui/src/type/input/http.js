import util from 'pkg/util'

export const httpInputConfigFilter = {
    url(t) {
        return util.isString(t) ? t : httpInputConfigDefault.url()
    },
    ref(t) {
        return util.isNumber(t) ? t : httpInputConfigDefault.ref()
    }
}

export const httpInputConfigDefault = {
    url() {
        return ''
    },
    ref() {
        return 0
    }
}

export const httpInputConfigParse = (t) => {
    let cfg = httpInputConfig()
    cfg.url = httpInputConfigFilter.url(t.url)
    cfg.ref = httpInputConfigFilter.url(t.ref)
    return cfg
}

export const httpInputConfig = () => {
    return {
        url: httpInputConfigDefault.url(),
        ref: httpInputConfigDefault.ref()
    }
}