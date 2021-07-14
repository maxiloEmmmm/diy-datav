import util from 'pkg/util'

export const StaticTextVerticalAlignOptions = [
    {label: '默认', value: 'stretch'},
    {label: '靠上', value: 'flex-start'},
    {label: '居中', value: 'center'},
    {label: '靠下️', value: 'flex-end'},
]
export const StaticTextHorizontalAlignOptions = [
    {label: '默认', value: 'initial'},
    {label: '靠左', value: 'flex-start'},
    {label: '居中', value: 'center'},
    {label: '靠右️', value: 'flex-end'},
    {label: '均匀', value: 'space-between'},
    {label: '均匀(两边留白)', value: 'space-around'},
]

export const StaticTextConfigFilter = {
    text: t => util.isString(t) ? t : StaticTextConfigDefault.text(),
    align(t) {
        if(!util.isObject(t)) {
            return StaticTextConfigDefault.align()
        }

        return {
            h: StaticTextConfigFilter.alignH(t.h),
            v: StaticTextConfigFilter.alignV(t.v)
        }
    },
    alignH(t) {
        return StaticTextHorizontalAlignOptions.some(item => item.value === t) ? t : StaticTextConfigDefault.alignH()
    },
    alignV(t) {
        return StaticTextVerticalAlignOptions.some(item => item.value === t) ? t : StaticTextConfigDefault.alignV()
    },
    bold(t) {
        return util.isBoolean(t) ? t : StaticTextConfigDefault.bold()
    },
    italic(t) {
        return util.isBoolean(t) ? t : StaticTextConfigDefault.italic()
    },
    underline(t) {
        return util.isBoolean(t) ? t : StaticTextConfigDefault.underline()
    },
    sup(t) {
        return util.isString(t) ? t : StaticTextConfigDefault.sup()
    },
    sub(t) {
        return util.isString(t) ? t : StaticTextConfigDefault.sub()
    },
    color(t) {
        return util.isString(t) ? t : StaticTextConfigDefault.color()
    },
    size(t) {
        return util.isNumber(t) ? t : StaticTextConfigDefault.size()
    }
}

export const StaticTextConfigDefault = {
    text: () => '',
    align() {
        return {
            h: StaticTextConfigDefault.alignH(),
            v: StaticTextConfigDefault.alignV()
        }
    },
    alignH(){
        return 'initial'
    },
    alignV() {
        return 'stretch'
    },
    bold() {
        return false
    },
    italic() {
        return false
    },
    underline() {
        return false
    },
    sup() {
        return ''
    },
    sub() {
        return ''
    },
    color() {
        return '#000'
    },
    size() {
        return 1
    }
}

export const StaticTextConfigParse = (config) => {
    let cfg = StaticTextConfig()
    cfg.text = StaticTextConfigFilter.text(config.text)
    cfg.align = StaticTextConfigFilter.align(config.align)
    cfg.bold = StaticTextConfigFilter.bold(config.bold)
    cfg.italic = StaticTextConfigFilter.italic(config.italic)
    cfg.underline = StaticTextConfigFilter.underline(config.underline)
    cfg.sup = StaticTextConfigFilter.sup(config.sup)
    cfg.sub = StaticTextConfigFilter.sub(config.sub)
    cfg.color = StaticTextConfigFilter.color(config.color)
    cfg.size = StaticTextConfigFilter.size(config.size)
    return cfg
}

export const StaticTextConfig = function(){
    return {
        text: StaticTextConfigDefault.text(),
        align: StaticTextConfigDefault.align(),
        bold: StaticTextConfigDefault.bold(),
        italic: StaticTextConfigDefault.italic(),
        underline: StaticTextConfigDefault.underline(),
        sup: StaticTextConfigDefault.sup(),
        sub: StaticTextConfigDefault.sub(),
        color: StaticTextConfigDefault.color(),
        size: StaticTextConfigDefault.size(),
    }
}