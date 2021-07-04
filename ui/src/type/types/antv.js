import util from 'pkg/util'

const AntVGeometryType = ['interval', 'point', 'line', 'area', 'path', 'polygon', 'edge', 'heatmap', 'schema']

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
    catColor(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.catColor()
        }
        return {
            single: AntVConfigDefault.catSingle(),
            enum: AntVConfigDefault.catEnum(),
            default: AntVConfigDefault.catColorDefault(),
        }
    },
    catSize(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.catSize()
        }
        return {
            single: AntVConfigDefault.catSingle(),
            enum: AntVConfigDefault.catEnum(),
            default: AntVConfigDefault.catSizeDefault(),
        }
    },
    catShape(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.catShape()
        }
        return {
            single: AntVConfigDefault.catSingle(),
            enum: AntVConfigDefault.catEnum(),
            default: AntVConfigDefault.catShapeDefault(),
        }
    },
    dataIndex(t) {
        let index = parseInt(t)
        if(index < 0) {
            return AntVConfigDefault.dataIndex()
        }
        return index
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
    catColor() {
        return {
            single: true, enum: [], default: AntVConfigDefault.catColorDefault()
        }
    },
    catColorDefault() {
        return 'blue'
    },
    catSize() {
        return {
            single: true, enum: [], default: AntVConfigDefault.catSizeDefault()
        }
    },
    catSizeDefault() {
        return 1
    },
    catShape() {
        return {
            single: true, enum: [], default: AntVConfigDefault.catShapeDefault()
        }
    },
    catShapeDefault() {
        return 'circle'
    },
    catSingle() {
        return true
    },
    catEnum() {
        return []
    },
    dataIndex() {
        return 0
    }
}

export const AntVConfigParse = function(config) {
    let cfg = AntVConfig()

    cfg.type = AntVConfigFilter.type(config?.type)
    cfg.color = AntVConfigFilter.color(config?.color)

    if (util.has(config, 'coordinate.type')) {
        cfg.coordinate.type = AntVConfigFilter.coordinateType(config.coordinate.type)
    }

    if (util.has(config, 'coordinate.transpose')) {
        cfg.coordinate.transpose = AntVConfigFilter.coordinateTranspose(config.coordinate.transpose)
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

    if (util.isObject(config?.cat)) {
        config.cat.color = AntVConfigFilter.catColor(config.cat.color)
        config.cat.size = AntVConfigFilter.catSize(config.cat.size)
        config.cat.shape = AntVConfigFilter.catShape(config.cat.shape)
    }

    cfg.dataIndex = AntVConfigFilter.dataIndex(config?.dataIndex)

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
        cat: {
            color: {default: 'blue', enum: [], single: true},
            size: {default: 1, enum: [], single: true},
            shape: {default: 'circle', enum: [], single: true},
        },
        dataIndex: 0,
    }
}