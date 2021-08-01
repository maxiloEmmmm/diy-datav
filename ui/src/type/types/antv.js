import util from 'pkg/util'

export const AntVScaleFields = ['x', 'y', 'z']
export const AntVGeometryType = ['interval', 'point', 'line', 'area', 'path', 'polygon', 'edge', 'heatmap', 'schema']
export const AntVAdjustType = ['stack', 'jitter', 'dodge', 'symmetric']
export const AntVFacetType = ['rect', 'mirror', 'list', 'matrix', 'circle', 'tree']
export const AntVScaleTypeType = [
    {label: '分类度量', value: 'cat'},
    {label: '时间分类度量', value: 'timeCat'},
    {label: '线性度量', value: 'linear'},
    {label: '连续的时间度量', value: 'time'},
    {label: 'log 度量', value: 'log'},
    {label: 'pow 度量', value: 'pow'},
    {label: '分段度量，用户可以指定不均匀的分段', value: 'quantize'},
    {label: '等分度量，根据数据的分布自动计算分段', value: 'quantile'},
    {label: '常量度量', value: 'identity'},
]
export const AntVIsPolarCoordinate = (t) => {
    return ['polar', 'theta'].includes(t)
}
export const AntVLegendLayout = [
    {label: '垂直', value: 'vertical'},
    {label: '水平', value: 'horizontal'},
]
export const AntVLegendPosition = [
    {label: 'left', value: 'left'},
    {label: 'left-top', value: 'left-top'},
    {label: 'left-bottom', value: 'left-bottom'},
    {label: 'right', value: 'right'},
    {label: 'right-top', value: 'right-top'},
    {label: 'right-bottom', value: 'right-bottom'},
    {label: 'top', value: 'top'},
    {label: 'top-left', value: 'top-left'},
    {label: 'top-right', value: 'top-right'},
    {label: 'bottom', value: 'bottom'},
    {label: 'bottom-left', value: 'bottom-left'},
    {label: 'bottom-right', value: 'bottom-right'}
]
export const AntVCoordinateAxis = (coordinate, transpose) => {
    coordinate = AntVConfigFilter.coordinateType(coordinate)
    let base = [AntVScaleFields[0]]
    if ({
        polar: 2,
        theta: 1,
        rect: 2,
        cartesian: 2,
        helix: 2
    }[coordinate] > 1) {
        if(coordinate !== 'polar' || !transpose) {
            base.push(AntVScaleFields[1])
        }
    }

    return base
}
export const AntVConfigFilter = {
    coordinate(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.coordinate()
        }

        return {
            type: AntVConfigFilter.coordinateType(t.type),
            cfg: AntVConfigFilter.coordinateCfg(t.cfg),
            transpose: AntVConfigFilter.coordinateTranspose(t.transpose)
        }
    },
    coordinateCfg(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.coordinateCfg()
        }

        return {
            radius: AntVConfigFilter.coordinateCfgRadius(t.radius),
            innerRadius: AntVConfigFilter.coordinateCfgInnerRadius(t.innerRadius)
        }
    },
    coordinateCfgRadius(t) {
        return util.isNumber(t) ? t : AntVConfigDefault.coordinateCfgRadius()
    },
    coordinateCfgInnerRadius(t) {
        return util.isNumber(t) ? t : AntVConfigDefault.coordinateCfgInnerRadius()
    },
    coordinateType(t) {
        return !!t ? t : AntVConfigDefault.coordinateType()
    },
    coordinateTranspose(t) {
        return !!t ? t : AntVConfigDefault.coordinateTranspose()
    },
    scaleType(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.scaleType()
        }

        return {
            type: AntVConfigFilter.scaleTypeType(t.type)
        }
    },
    scaleTypeType(t) {
        return !!AntVScaleTypeType.filter(typ => typ.value === t) ? t : AntVConfigDefault.scaleTypeType()
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
    layerCatItemField(t) {
        return !!t ? t : AntVConfigDefault.layerCatItemField()
    },
    layerCatColor(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.layerCatColor()
        }
        return {
            enable: AntVConfigFilter.layerCatItemEnable(t.enable),
            field: AntVConfigFilter.layerCatItemField(t.field),
            single: AntVConfigFilter.layerCatColorSingle(t.single),
            enum: AntVConfigFilter.layerCatColorEnum(t.enum),
            default: AntVConfigFilter.layerCatColorDefault(t.default),
        }
    },
    layerCatColorEnum(t) {
        // at least one
        if(util.isArray(t) && t.length >= 1) {
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
            enable: AntVConfigFilter.layerCatItemEnable(t.enable),
            field: AntVConfigFilter.layerCatItemField(t.field),
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
            enable: AntVConfigFilter.layerCatItemEnable(t.enable),
            field: AntVConfigFilter.layerCatItemField(t.field),
            single: AntVConfigFilter.layerCatShapeSingle(t.single),
            enum: AntVConfigFilter.layerCatShapeEnum(typ, t.enum),
            default: AntVConfigFilter.layerCatShapeDefault(typ, t.default),
        }
    },
    layerCatItemEnable(t) {
        return util.isBoolean(t) ? t : AntVConfigDefault.layerCatItemEnable()
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
        if(!util.isNumber(index) || index < -1) {
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
    legend(legend) {
        let cfg = AntVConfigDefault.legend()
        if (util.isObject(legend)) {
            cfg.enable = AntVConfigFilter.legendEnable(legend.enable)
            cfg.position = AntVConfigFilter.legendPosition(legend.position)
            cfg.layout = AntVConfigFilter.legendLayout(legend.layout)
            cfg.flipPage = AntVConfigFilter.legendFlipPage(legend.flipPage)
            cfg.maxRow = AntVConfigFilter.legendMaxRow(legend.maxRow)
        }
        return cfg
    },
    legendEnable(t) {
        return util.isBoolean(t) ? t : AntVConfigDefault.legendEnable()
    },
    legendPosition(t) {
        return AntVLegendPosition.some(p => p.value === t) ? t : AntVConfigDefault.legendPosition()
    },
    legendLayout(t) {
        return AntVLegendLayout.some(p => p.value === t) ? t : AntVConfigDefault.legendLayout()
    },
    legendFlipPage(t) {
        return util.isBoolean(t) ? t : AntVConfigDefault.legendFlipPage()
    },
    legendMaxRow(t) {
        return util.isNumber(t) && t > 0 ? t : AntVConfigDefault.legendMaxRow()
    },
    layers(t) {
        if (!util.isArray(t)) {
            return AntVConfigDefault.layers()
        }
        return t.map(AntVConfigFilter.layer)
    },
    facet(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.facet()
        }

        return {
            enable: AntVConfigFilter.facetEnable(t.enable),
            field: AntVConfigFilter.facetField(t.field),
            type: AntVConfigFilter.facetType(t.type)
        }
    },
    facetEnable(t) {
        return util.isBoolean(t) ? t : AntVConfigDefault.facetEnable()
    },
    facetField(t) {
        return !!t ? t : AntVConfigDefault.facetField()
    },
    facetType(t) {
        return AntVFacetType.includes(t) ? t : AntVConfigDefault.facetType()
    },
    scale(t) {
        if(!util.isObject(t)) {
            return AntVConfigDefault.scale()
        }

        let s = {}
        AntVScaleFields.forEach(scale => {
            s[scale] = AntVConfigFilter.scaleItem(t[scale])
        })
        return s
    },
    scaleItem(t) {
        let cfg = AntVConfigDefault.scaleItem()
        if(util.isObject(t)) {
            cfg.type = AntVConfigFilter.scaleType(t.type)
            cfg.field = AntVConfigFilter.scaleField(t.field)
            cfg.alias = AntVConfigFilter.scaleAlias(t.alias)
            cfg.format.prefix = AntVConfigFilter.scaleFormatPrefix(t.format.prefix)
            cfg.format.suffix = AntVConfigFilter.scaleFormatSuffix(t.format.suffix)
        }
        return cfg
    },
}

export const AntVConfigDefault = {
    coordinate() {
        return {
            type: AntVConfigDefault.coordinateType(),
            cfg: AntVConfigDefault.coordinateCfg(),
            transpose: AntVConfigDefault.coordinateTranspose()
        }
    },
    coordinateCfg(){
        return {
            radius: AntVConfigDefault.coordinateCfgRadius(),
            innerRadius: AntVConfigDefault.coordinateCfgInnerRadius()
        }
    },
    coordinateCfgRadius() {
        return null
    },
    coordinateCfgInnerRadius() {
        return null
    },
    coordinateType() {
        return 'cartesian'
    },
    coordinateTranspose() {
        return false
    },
    scaleType() {
        return {
            type: AntVConfigDefault.scaleTypeType()
        }
    },
    scaleTypeType() {
        return ''
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
            enable: AntVConfigDefault.layerCatItemEnable(), field: AntVConfigDefault.layerCatItemField(), single: AntVConfigDefault.layerCatColorSingle(), enum: AntVConfigDefault.layerCatColorEnum(), default: AntVConfigDefault.layerCatColorDefault()
        }
    },
    layerCatSize() {
        return {
            enable: AntVConfigDefault.layerCatItemEnable(), field: AntVConfigDefault.layerCatItemField(), single: AntVConfigDefault.layerCatSizeSingle(), enum: AntVConfigDefault.layerCatSizeEnum(), default: AntVConfigDefault.layerCatSizeDefault()
        }
    },
    layerCatShape(typ) {
        return {
            enable: AntVConfigDefault.layerCatItemEnable(), field: AntVConfigDefault.layerCatItemField(), single: AntVConfigDefault.layerCatShapeSingle(), enum: AntVConfigDefault.layerCatShapeEnum(), default: AntVConfigDefault.layerCatShapeDefault(typ)
        }
    },
    layerCatItemField() {
        return ""
    },
    layerCatItemEnable() {
        return false
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
        // at least one
        return [AntVConfigDefault.layerCatColorDefault()]
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
        return -1
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
    facet() {
        return {
            enable: AntVConfigDefault.facetEnable(),
            field: AntVConfigDefault.facetField(),
            type: AntVConfigDefault.facetType()
        }
    },
    facetEnable() {
        return false
    },
    facetField() {
        return ''
    },
    facetType() {
        return AntVFacetType[0]
    },
    scale() {
        let s = {}
        AntVScaleFields.forEach(scale => {
            s[scale] = AntVConfigDefault.scaleItem()
        })
        return s
    },
    scaleItem() {
        return {
            field: AntVConfigDefault.scaleField(),
            alias: AntVConfigDefault.scaleAlias(),
            format: {
                prefix: AntVConfigDefault.scaleFormatPrefix(),
                suffix: AntVConfigDefault.scaleFormatSuffix()
            },
            type: AntVConfigDefault.scaleType()
        }
    },
    legend() {
        return {
            enable: AntVConfigDefault.legendEnable(),
            position: AntVConfigDefault.legendPosition(),
            layout: AntVConfigDefault.legendLayout(),
            flipPage: AntVConfigDefault.legendFlipPage(),
            maxRow: AntVConfigDefault.legendMaxRow(),
        }
    },
    legendEnable() {
        return false
    },
    legendPosition() {
        return AntVLegendPosition[0].value
    },
    legendLayout() {
        return AntVLegendLayout[0].value
    },
    legendMaxRow() {
        return 5
    },
    legendFlipPage() {
        return true
    }
}

export const AntVConfigParse = function(config) {
    let cfg = AntVConfig()

    cfg.coordinate = AntVConfigFilter.coordinate(config?.coordinate)
    cfg.scale = AntVConfigFilter.scale(config?.scale)
    cfg.layers = AntVConfigFilter.layers(config?.layers)
    cfg.legend = AntVConfigFilter.legend(config?.legend)
    cfg.facet = AntVConfigFilter.facet(config?.facet)
    cfg.dataIndex = AntVConfigFilter.dataIndex(config?.dataIndex)

    return cfg
}

export const AntVConfig = function() {
    return {
        coordinate: AntVConfigDefault.coordinate(),
        scale: AntVConfigDefault.scale(),
        layers: AntVConfigDefault.layers(),
        facet: AntVConfigDefault.facet(),
        dataIndex: AntVConfigDefault.dataIndex(),
    }
}