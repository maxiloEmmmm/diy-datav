import util from 'pkg/util'

const AntVGeometryType = ['interval', 'point', 'line', 'area', 'path', 'polygon', 'edge', 'heatmap', 'schema']

const AntVConfigFilter = {
    type(t) {
        if(AntVGeometryType.includes(t)) {
            return t
        }

        return AntVConfigDefault.type()
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
    catColorSingle(t) {
        return util.isBoolean(t) ? t : AntVConfigDefault.catColorSingle()
    },
    catColorDefault(t) {
        return !!t ? t : AntVConfigDefault.catColorDefault()
    },
    catColor(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.catColor()
        }
        return {
            single: AntVConfigFilter.catColorSingle(t.single),
            enum: AntVConfigFilter.catColorEnum(t.enum),
            default: AntVConfigFilter.catColorDefault(t.default),
        }
    },
    catColorEnum(t) {
        if(util.isArray(t)) {
            return t.filter(o => util.isString(o) && !!o)
        }

        return AntVConfigDefault.catColorEnum()
    },
    catSize(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.catSize()
        }
        return {
            single: AntVConfigFilter.catSizeSingle(t.single),
            enum: AntVConfigFilter.catSizeEnum(t.enum),
            default: AntVConfigFilter.catSizeDefault(t.default),
        }
    },
    catSizeDefault(t) {
        return util.isNumber(t) ? t : AntVConfigDefault.catSizeDefault()
    },
    catSizeSingle(t) {
        return util.isBoolean(t) ? t : AntVConfigDefault.catSizeSingle()
    },
    catSizeEnum(t) {
        if(util.isArray(t)) {
            return t.filter(o => util.isNumber(o) && o !== 0)
        }

        return AntVConfigDefault.catSizeEnum()
    },
    catShape(typ, t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.catShape()
        }
        return {
            single: AntVConfigFilter.catShapeSingle(t.single),
            enum: AntVConfigFilter.catShapeEnum(typ, t.enum),
            default: AntVConfigFilter.catShapeDefault(typ, t.default),
        }
    },
    catShapeDefault(typ, t) {
        return util.isString(t) && !!t ? t : AntVConfigDefault.catShapeDefault(typ)
    },
    catShapeSingle(t) {
        return util.isBoolean(t) ? t : AntVConfigDefault.catShapeSingle()
    },
    catShapeEnum(typ, t) {
        if(util.isArray(t)) {
            return t.filter(o => util.string(o) && !!o)
        }

        return AntVConfigDefault.catShapeEnum(typ)
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
            single: AntVConfigDefault.catColorSingle(), enum: AntVConfigDefault.catColorEnum(), default: AntVConfigDefault.catColorDefault()
        }
    },
    catSize() {
        return {
            single: AntVConfigDefault.catSizeDefault(), enum: AntVConfigDefault.catSizeEnum(), default: AntVConfigDefault.catSizeDefault()
        }
    },
    catShape() {
        return {
            single: AntVConfigDefault.catShapeSingle(), enum: AntVConfigDefault.catShapeEnum(), default: AntVConfigDefault.catShapeDefault()
        }
    },
    catColorDefault() {
        return '#1890ff'
    },
    catSizeDefault() {
        return 5
    },
    catShapeDefault(typ) {
        return AntVConfigDefault.catShapeEnum(typ)[0]
    },
    catShapeEnum(typ) {
        switch (typ) {
            case 'interval':
                return ['rect']
            case 'point':
                return ['circle']
            case 'path':
            case 'line':
                return ['line']
            case 'area':
                return ['area']
            case 'heatmap':
            case 'polygon':
                return ['polygon']
            case 'schema':
                return ['schema']
            case 'edge':
                return ['line']
            default:
                throw `bug: AntVConfigDefault.catShapeEnum unknown type: ${typ}`
        }
    },
    catColorEnum() {
        return ['#1890ff', '#5AD8A6']
    },
    catSizeEnum() {
        return [1, 10]
    },
    catColorSingle() {
        return true
    },
    catSizeSingle() {
        return true
    },
    catShapeSingle() {
        return true
    },
    dataIndex() {
        return 0
    }
}

export const AntVConfigParse = function(config) {
    let cfg = AntVConfig()

    cfg.type = AntVConfigFilter.type(config?.type)

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
        cfg.cat.color = AntVConfigFilter.catColor(config.cat.color)
        cfg.cat.size = AntVConfigFilter.catSize(config.cat.size)
        cfg.cat.shape = AntVConfigFilter.catShape(cfg.type, config.cat.shape)
    }

    cfg.dataIndex = AntVConfigFilter.dataIndex(config?.dataIndex)

    return cfg
}

export const AntVConfig = function() {
    return {
        // TODO: 支持多类型
        type: '',
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
            color: {default: '#1890ff', enum: [], single: true},
            size: {default: 1, enum: [], single: true},
            shape: {default: 'circle', enum: [], single: true},
        },
        dataIndex: 0,
    }
}