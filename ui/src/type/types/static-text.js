import util from 'pkg/util'

export const StaticTextConfigFilter = {
    text: t => util.isString(t) ? t : StaticTextConfigDefault.text(),
    center(t) {
        if(!util.isObject(t)) {
            return StaticTextConfigDefault.center()
        }

        return {
            h: StaticTextConfigFilter.centerH(t.h),
            v: StaticTextConfigFilter.centerV(t.v)
        }
    },
    centerH(t) {
        return util.isBoolean(t) ? t : StaticTextConfigDefault.centerH()
    },
    centerV(t) {
        return util.isBoolean(t) ? t : StaticTextConfigDefault.centerV()
    },
    bold(t) {
        return util.isBoolean(t) ? t : StaticTextConfigDefault.bold()
    },
    italic(t) {
        return util.isBoolean(t) ? t : StaticTextConfigDefault.italic()
    }
}

export const StaticTextConfigDefault = {
    text: () => '',
    center() {
        return {
            h: StaticTextConfigDefault.centerH(),
            v: StaticTextConfigDefault.centerV()
        }
    },
    centerH(){
        return false
    },
    centerV() {
        return false
    },
    bold() {
        return false
    },
    italic() {
        return false
    }
}

export const StaticTextConfigParse = (config) => {
    let cfg = StaticTextConfig()
    cfg.text = StaticTextConfigFilter.text(config.text)
    cfg.center = StaticTextConfigFilter.center(config.center)
    cfg.bold = StaticTextConfigFilter.bold(config.bold)
    cfg.italic = StaticTextConfigFilter.italic(config.italic)
    return cfg
}

export const StaticTextConfig = function(){
    return {
        text: StaticTextConfigDefault.text(),
        center: StaticTextConfigDefault.center(),
        bold: StaticTextConfigDefault.bold(),
        italic: StaticTextConfigDefault.italic()
    }
}