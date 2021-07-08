import util from 'pkg/util'

export const AntVGeometryType = ['interval', 'point', 'line', 'area', 'path', 'polygon', 'edge', 'heatmap', 'schema']
export const AntVAdjustType = ['stack', 'jitter', 'dodge', 'symmetric']
export const AntVConfigFilter = {
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
    layerCatColorSingle(t) {
        return util.isBoolean(t) ? t : AntVConfigDefault.layerCatColorSingle()
    },
    layerCatColorDefault(t) {
        return !!t ? t : AntVConfigDefault.layerCatColorDefault()
    },
    layerCatField(t) {
        return !!t ? t : AntVConfigDefault.layerCatField()
    },
    layerCatColor(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.layerCatColor()
        }
        return {
            field: AntVConfigFilter.layerCatField(t.field),
            single: AntVConfigFilter.layerCatColorSingle(t.single),
            enum: AntVConfigFilter.layerCatColorEnum(t.enum),
            default: AntVConfigFilter.layerCatColorDefault(t.default),
        }
    },
    layerCatColorEnum(t) {
        if(util.isArray(t)) {
            return t.filter(AntVConfigFilter.layerCatColorEnumItem)
        }

        return AntVConfigDefault.layerCatColorEnum()
    },
    layerCatColorEnumItem(t) {
        return util.isString(t) && !!t ? t : AntVConfigDefault.layerCatColorDefault()
    },
    layerCatSize(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.layerCatSize()
        }
        return {
            field: AntVConfigFilter.layerCatField(t.field),
            single: AntVConfigFilter.layerCatSizeSingle(t.single),
            enum: AntVConfigFilter.layerCatSizeEnum(t.enum),
            default: AntVConfigFilter.layerCatSizeDefault(t.default),
        }
    },
    layerCatSizeDefault(t) {
        return util.isNumber(t) ? t : AntVConfigDefault.layerCatSizeDefault()
    },
    layerCatSizeSingle(t) {
        return util.isBoolean(t) ? t : AntVConfigDefault.layerCatSizeSingle()
    },
    layerCatSizeEnum(t) {
        if(util.isArray(t)) {
            return t.filter(AntVConfigFilter.layerCatSizeEnumItem)
        }

        return AntVConfigDefault.layerCatSizeEnum()
    },
    layerCatSizeEnumItem(t) {
        return util.isNumber(t) && t !== 0 ? t : AntVConfigDefault.layerCatSizeDefault()
    },
    layerCatShape(typ, t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.layerCatShape(typ)
        }
        return {
            field: AntVConfigFilter.layerCatField(t.field),
            single: AntVConfigFilter.layerCatShapeSingle(t.single),
            enum: AntVConfigFilter.layerCatShapeEnum(typ, t.enum),
            default: AntVConfigFilter.layerCatShapeDefault(typ, t.default),
        }
    },
    layerCatShapeDefault(typ, t) {
        return util.isString(t) && !!t ? t : AntVConfigDefault.layerCatShapeDefault(typ)
    },
    layerCatShapeSingle(t) {
        return util.isBoolean(t) ? t : AntVConfigDefault.layerCatShapeSingle()
    },
    layerCatShapeEnum(typ, t) {
        if(util.isArray(t)) {
            return t.filter(o => AntVConfigFilter.layerCatShapeEnumItem(typ, o))
        }

        return AntVConfigDefault.layerCatShapeEnum()
    },
    layerCatShapeEnumItem(typ, t) {
        return util.isString(t) && !!t ? t : AntVConfigDefault.layerCatShapeDefault(typ)
    },
    dataIndex(t) {
        let index = parseInt(t)
        if(index < 0) {
            return AntVConfigDefault.dataIndex()
        }
        return index
    },
    layerAdjustType(t) {
        return AntVAdjustType.includes(t) ? t : AntVConfigDefault.layerAdjustType()
    },
    layerAdjustEnable(t) {
        return util.isBoolean(t) ? t : AntVConfigDefault.layerAdjustEnable()
    },
    layerType(t) {
        if(AntVGeometryType.includes(t)) {
            return t
        }

        return AntVConfigDefault.layerType()
    },
    layer(layer) {
        let lt = AntVConfigDefault.layer()
        if (!util.isObject(layer)) {
            return lt
        }

        lt.type = AntVConfigFilter.layerType(layer.type)

        if (util.isObject(layer.cat)) {
            lt.cat.color = AntVConfigFilter.layerCatColor(layer.cat.color)
            lt.cat.size = AntVConfigFilter.layerCatSize(layer.cat.size)
            lt.cat.shape = AntVConfigFilter.layerCatShape(lt.type, layer.cat.shape)
        }else {
            lt.cat = AntVConfigDefault.layerCat(layer.type)
        }

        if (util.isObject(layer?.adjust)) {
            lt.adjust.type = AntVConfigFilter.layerAdjustType(layer.adjust.type)
            lt.adjust.enable = AntVConfigFilter.layerAdjustEnable(layer.adjust.enable)
        }else {
            lt.adjust = AntVConfigDefault.layerAdjust()
        }

        return lt
    },
}

export const AntVConfigDefault = {
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
    layerCatColor() {
        return {
            field: AntVConfigDefault.layerCatField(), single: AntVConfigDefault.layerCatColorSingle(), enum: AntVConfigDefault.layerCatColorEnum(), default: AntVConfigDefault.layerCatColorDefault()
        }
    },
    layerCatSize() {
        return {
            field: AntVConfigDefault.layerCatField(), single: AntVConfigDefault.layerCatSizeSingle(), enum: AntVConfigDefault.layerCatSizeEnum(), default: AntVConfigDefault.layerCatSizeDefault()
        }
    },
    layerCatShape(typ) {
        return {
            field: AntVConfigDefault.layerCatField(), single: AntVConfigDefault.layerCatShapeSingle(), enum: AntVConfigDefault.layerCatShapeEnum(), default: AntVConfigDefault.layerCatShapeDefault(typ)
        }
    },
    layerCatField() {
        return ""
    },
    layerCatColorDefault() {
        return '#1890ff'
    },
    layerCatSizeDefault() {
        return 5
    },
    layerCatShapeDefault(typ) {
        switch (typ) {
            case 'interval':
                return 'rect'
            case 'point':
                return 'circle'
            case 'path':
            case 'line':
                return 'line'
            case 'area':
                return 'area'
            case 'heatmap':
            case 'polygon':
                return 'polygon'
            case 'schema':
                return 'schema'
            case 'edge':
                return 'line'
            default:
                throw `bug: AntVConfigDefault.layerCatShapeDefault unknown type: ${typ}`
        }
    },
    layerCatShapeEnum() {
        return []
    },
    layerCatColorEnum() {
        return []
    },
    layerCatSizeEnum() {
        return [1, 10]
    },
    layerCatColorSingle() {
        return true
    },
    layerCatSizeSingle() {
        return true
    },
    layerCatShapeSingle() {
        return true
    },
    dataIndex() {
        return 0
    },
    layerAdjustEnable() {
        return false
    },
    layerAdjustType() {
        return 'stack'
    },
    layerCat(typ) {
        return {
            color: AntVConfigDefault.layerCatColor(),
            size: AntVConfigDefault.layerCatSize(),
            shape: AntVConfigDefault.layerCatShape(typ)
        }
    },
    layerType() {
        return AntVGeometryType[0]
    },
    layerAdjust() {
        return {
            enable: AntVConfigDefault.layerAdjustEnable(),
            type: AntVConfigDefault.layerAdjustType()
        }
    },
    layer() {
        return {
            type: AntVConfigDefault.layerType(),
            adjust: AntVConfigDefault.layerAdjust(),
            cat: AntVConfigDefault.layerCat(AntVConfigDefault.layerType())
        }
    },
    layers() {
        return []
    },
}

export const AntVConfigParse = function(config) {
    let cfg = AntVConfig()

    if (util.has(config, 'coordinate.type')) {
        cfg.coordinate.type = AntVConfigFilter.coordinateType(config.coordinate.type)
    }

    if (util.has(config, 'coordinate.transpose')) {
        cfg.coordinate.transpose = AntVConfigFilter.coordinateTranspose(config.coordinate.transpose)
    }

    ['x', 'y', 'z'].forEach(s => {
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

    if (util.isArray(config?.layers)) {
        config.layers.forEach(layer => cfg.layers.push(AntVConfigFilter.layer(layer)))
    }

    cfg.dataIndex = AntVConfigFilter.dataIndex(config?.dataIndex)

    return cfg
}

export const AntVConfig = function() {
    return {
        // TODO: 支持多类型
        coordinate: {
            type: AntVConfigDefault.coordinateType(),
            transpose: AntVConfigDefault.coordinateTranspose()
        },
        scale: {
            x: {
                field: AntVConfigDefault.scaleField(),
                alias: AntVConfigDefault.scaleAlias(),
                format: {
                    prefix: AntVConfigDefault.scaleFormatPrefix(),
                    suffix: AntVConfigDefault.scaleFormatSuffix()
                }
            },
            y: {
                field: AntVConfigDefault.scaleField(),
                alias: AntVConfigDefault.scaleAlias(),
                format: {
                    prefix: AntVConfigDefault.scaleFormatPrefix(),
                    suffix: AntVConfigDefault.scaleFormatSuffix()
                }
            },
            z: {
                field: AntVConfigDefault.scaleField(),
                alias: AntVConfigDefault.scaleAlias(),
                format: {
                    prefix: AntVConfigDefault.scaleFormatPrefix(),
                    suffix: AntVConfigDefault.scaleFormatSuffix()
                }
            }
        },
        layers: AntVConfigDefault.layers(),
        dataIndex: AntVConfigDefault.dataIndex(),
    }
}