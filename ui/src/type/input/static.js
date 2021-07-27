import util from 'pkg/util'
import * as common from './common'
export const staticInputConfigFilter = {
    ...common.commonInputConfigFilter,
    id(t) {
        return util.isNumber(t) ? t : staticInputConfigDefault.id()
    },
}

export const staticInputConfigDefault = {
    ...common.commonInputConfigDefault,
    id() {
        return 0
    },
}

export const staticInputConfigParse = (t) => {
    let cfg = staticInputConfig()
    cfg.id = staticInputConfigFilter.id(t.id)
    return {
        ...common.commonInputConfigParse(t),
        ...cfg,
    }
}

export const staticInputConfig = () => {
    return {
        ...common.commonInputConfig(),
        id: staticInputConfigDefault.id(),
    }
}