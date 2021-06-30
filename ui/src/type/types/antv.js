import util from 'pkg/util'

const AntVGeometryType = ['interval']

const AntVConfigFilter = {
    type(t) {
        if(AntVGeometryType.includes(t)) {
            return t
        }

        return AntVConfigDefault.type()
    },
    color(t) {
        return !!t ? t : AntVConfigDefault.color()
    },
    coordinateType(t) {
        return !!t ? t : AntVConfigDefault.coordinateType()
    },
    coordinateTranspose(t) {
        return !!t ? t : AntVConfigDefault.coordinateTranspose()
    },
    scaleField(t) {
        return !!t ? t : AntVConfigDefault.scaleField()
    },
    scaleAlias(t) {
        return !!t ? t : AntVConfigDefault.scaleAlias()
    },
    scaleFormatPrefix(t) {
        return !!t ? t : AntVConfigDefault.scaleFormatPrefix()
    },
    scaleFormatSuffix(t) {
        return !!t ? t : AntVConfigDefault.scaleFormatSuffix()
    },
    catField(t) {
        return !!t ? t : AntVConfigDefault.catField()
    },
    catEnum(t) {
        return util.isArray(t) ? t : AntVConfigDefault.catEnum()
    }
}

const AntVConfigDefault = {
    type() {
        return AntVGeometryType[0]
    },
    color() {
        return ''
    },
    coordinateType() {
        return 'cartesian'
    },
    coordinateTranspose() {
        return false
    },
    scaleField() {
        return ''
    },
    scaleAlias() {
        return ''
    },
    scaleFormatPrefix() {
        return ''
    },
    scaleFormatSuffix() {
        return ''
    },
    catField() {
        return ''
    },
    catEnum() {
        return []
    }
}

export const AntVConfigParse = function(config) {
    let cfg = AntVConfig()

    cfg.type = AntVConfigFilter.type(config?.type)
    cfg.color = AntVConfigFilter.color(config?.color)

    if (util.has(config, 'coordinate.type')) {
        cfg.coordinate.type = AntVConfigFilter.coordinateType(config.color)
    }

    if (util.has(config, 'coordinate.transpose')) {
        cfg.coordinate.transpose = AntVConfigFilter.coordinateTranspose(config.color)
    }

    ['x', 'y'].forEach(s => {
        if (util.has(config, `scale.${s}.field`)) {
            cfg.scale[s].field = AntVConfigFilter.scaleField(config.scale[s].field)
        }
        if (util.has(config, `scale.${s}.alias`)) {
            cfg.scale[s].alias = AntVConfigFilter.scaleAlias(config.scale[s].alias)
        }
        if (util.has(config, `scale.${s}.format.prefix`)) {
            cfg.scale[s].format.prefix = AntVConfigFilter.scaleFormatPrefix(config.scale[s].format.prefix)
        }
        if (util.has(config, `scale.${s}.format.suffix`)) {
            cfg.scale[s].format.suffix = AntVConfigFilter.scaleFormatSuffix(config.scale[s].format.suffix)
        }
    })

    if (util.isArray(config?.cats)) {
        config.cats.forEach((cat, catIndex) => {
            cfg.cats[catIndex].field = AntVConfigFilter.catField(cat.field)
            cfg.scale[catIndex].enum = AntVConfigFilter.catEnum(cat.enum)
        })
    }

    return cfg
}

export const AntVConfig = function() {
    return {
        // TODO: 支持多类型
        type: '',
        color: '',
        coordinate: {
            type: '',
            transpose: false
        },
        scale: {
            x: {
                field: "",
                alias: "",
                format: {
                    prefix: '',
                    suffix: ''
                }
            },
            y: {
                field: "",
                alias: "",
                format: {
                    prefix: '',
                    suffix: ''
                }
            }
        },
        cats: []
    }
}