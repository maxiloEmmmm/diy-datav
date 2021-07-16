import util from 'pkg/util'

export const httpInputConfigFilter = {
    url(t) {
        return util.isString(t) ? t : httpInputConfigDefault.url()
    }
}

export const httpInputConfigDefault = {
    url() {
        return ''
    }
}

export const httpInputConfigParse = (t) => {
    let cfg = httpInputConfig()
    cfg.url = httpInputConfigFilter.url(t.url)
    return cfg
}

export const httpInputConfig = () => {
    return {
        url: httpInputConfigDefault.url()
    }
}