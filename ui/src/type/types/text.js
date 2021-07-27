import util from 'pkg/util'

export const TextVerticalAlignOptions = [
    {label: '默认', value: 'stretch'},
    {label: '靠上', value: 'flex-start'},
    {label: '居中', value: 'center'},
    {label: '靠下️', value: 'flex-end'},
]
export const TextHorizontalAlignOptions = [
    {label: '默认', value: 'initial'},
    {label: '靠左', value: 'flex-start'},
    {label: '居中', value: 'center'},
    {label: '靠右️', value: 'flex-end'},
    {label: '均匀', value: 'space-between'},
    {label: '均匀(两边留白)', value: 'space-around'},
]

export const TextConfigFilter = {
    text: t => util.isString(t) ? t : TextConfigDefault.text(),
    align(t) {
        if(!util.isObject(t)) {
            return TextConfigDefault.align()
        }

        return {
            h: TextConfigFilter.alignH(t.h),
            v: TextConfigFilter.alignV(t.v)
        }
    },
    alignH(t) {
        return TextHorizontalAlignOptions.some(item => item.value === t) ? t : TextConfigDefault.alignH()
    },
    alignV(t) {
        return TextVerticalAlignOptions.some(item => item.value === t) ? t : TextConfigDefault.alignV()
    },
    bold(t) {
        return util.isBoolean(t) ? t : TextConfigDefault.bold()
    },
    italic(t) {
        return util.isBoolean(t) ? t : TextConfigDefault.italic()
    },
    underline(t) {
        return util.isBoolean(t) ? t : TextConfigDefault.underline()
    },
    sup(t) {
        return util.isString(t) ? t : TextConfigDefault.sup()
    },
    sub(t) {
        return util.isString(t) ? t : TextConfigDefault.sub()
    },
    color(t) {
        return util.isString(t) ? t : TextConfigDefault.color()
    },
    size(t) {
        return util.isNumber(t) ? t : TextConfigDefault.size()
    }
}

export const TextConfigDefault = {
    text: () => '',
    align() {
        return {
            h: TextConfigDefault.alignH(),
            v: TextConfigDefault.alignV()
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

export const TextConfigParse = (config) => {
    let cfg = TextConfig()
    cfg.text = TextConfigFilter.text(config.text)
    cfg.align = TextConfigFilter.align(config.align)
    cfg.bold = TextConfigFilter.bold(config.bold)
    cfg.italic = TextConfigFilter.italic(config.italic)
    cfg.underline = TextConfigFilter.underline(config.underline)
    cfg.sup = TextConfigFilter.sup(config.sup)
    cfg.sub = TextConfigFilter.sub(config.sub)
    cfg.color = TextConfigFilter.color(config.color)
    cfg.size = TextConfigFilter.size(config.size)
    return cfg
}

export const TextConfig = function(){
    return {
        text: TextConfigDefault.text(),
        align: TextConfigDefault.align(),
        bold: TextConfigDefault.bold(),
        italic: TextConfigDefault.italic(),
        underline: TextConfigDefault.underline(),
        sup: TextConfigDefault.sup(),
        sub: TextConfigDefault.sub(),
        color: TextConfigDefault.color(),
        size: TextConfigDefault.size(),
    }
}