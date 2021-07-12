import util from 'pkg/util'

export const StaticTextFilter = t => util.isString(t) ? t : StaticTextDefault()

export const StaticTextDefault = () => ''

export const StaticTextParse = (config) => {
    let cfg = StaticText()
    cfg.text = StaticTextFilter(config.text)
    return cfg
}

export const StaticText = function(){
    return {
        text: StaticTextDefault()
    }
}