import util from 'pkg/util'
import * as common from './common'
export const staticInputConfigFilter = {
    id(t) {
        return util.isNumber(t) ? t : staticInputConfigDefault.id()
    },
    ...common.commonInputConfigFilter
}

export const staticInputConfigDefault = {
    id() {
        return 0
    },
    ...common.commonInputConfigDefault
}

export const staticInputConfigParse = (t) => {
    let cfg = staticInputConfig()
    cfg.id = staticInputConfigFilter.id(t.id)
    return {
        ...cfg,
        ...common.commonInputConfigParse(t)
    }
}

export const staticInputConfig = () => {
    return {
        id: staticInputConfigDefault.id(),
        ...common.commonInputConfig()
    }
}