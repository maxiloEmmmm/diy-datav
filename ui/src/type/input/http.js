import util from 'pkg/util'
import * as common from './common'
import {commonInputConfigParse} from "./common";

export const httpInputConfigFilter = {
    url(t) {
        return util.isString(t) ? t : httpInputConfigDefault.url()
    },
    ref(t) {
        return util.isNumber(t) ? t : httpInputConfigDefault.ref()
    },
    ...common.commonInputConfigFilter
}

export const httpInputConfigDefault = {
    url() {
        return ''
    },
    ref() {
        return 0
    },
    ...common.commonInputConfigDefault
}

export const httpInputConfigParse = (t) => {
    let cfg = httpInputConfig()
    cfg.url = httpInputConfigFilter.url(t.url)
    cfg.ref = httpInputConfigFilter.ref(t.ref)
    return {
        ...cfg,
        ...common.commonInputConfigParse(t)
    }
}

export const httpInputConfig = () => {
    return {
        url: httpInputConfigDefault.url(),
        ref: httpInputConfigDefault.ref(),
        ...common.commonInputConfig()
    }
}