import util from 'pkg/util'
import * as common from './common'
import {commonInputConfigParse} from "./common";

export const httpInputConfigFilter = {
    ...common.commonInputConfigFilter,
    url(t) {
        return util.isString(t) ? t : httpInputConfigDefault.url()
    },
    ref(t) {
        return util.isNumber(t) ? t : httpInputConfigDefault.ref()
    },
}

export const httpInputConfigDefault = {
    ...common.commonInputConfigDefault,
    url() {
        return ''
    },
    ref() {
        return 0
    },
}

export const httpInputConfigParse = (t) => {
    let cfg = httpInputConfig()
    cfg.url = httpInputConfigFilter.url(t.url)
    cfg.ref = httpInputConfigFilter.ref(t.ref)
    return {
        ...common.commonInputConfigParse(t),
        ...cfg,
    }
}

export const httpInputConfig = () => {
    return {
            ...common.commonInputConfig(),
        url: httpInputConfigDefault.url(),
        ref: httpInputConfigDefault.ref(),
    }
}