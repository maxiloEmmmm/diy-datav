<script lang="tsx">
import {
    AntVConfig, AntVConfigParse, AntVConfigDefault, ViewBlockType,
    AntVAdjustType, AntVCoordinateAxis, AntVScaleTypeType, AntVFacetType,
    AntVIsPolarCoordinate, AntVLegendPosition, AntVLegendLayout
} from 'type/index.js'
import configMixin from '../config-mixin'
import easyExample from './antv-config-easy-example'
import util from 'pkg/util'
import inputChoose from './input-choose.vue'
export default {
    components: {inputChoose},
    mixins: [configMixin],
    data() {
        return {
            cfg: {
                ...ViewBlockType().config,
                type: AntVConfig()
            },
            store: {
                adjustOptions: AntVAdjustType.map(typ => ({label: typ, value: typ})),
                facetTypeOptions: AntVFacetType.map(typ => ({label: typ, value: typ})),
                scaleTypeTypeOptions: [
                    {label: '自动判断', value: AntVConfigDefault.scaleTypeType()},
                    ...AntVScaleTypeType,
                ],
                typeOptions: [
                    {label: 'interval', value: 'interval', desc: '用于绘制柱状图、直方图、南丁格尔玫瑰图、饼图、条形环图（玉缺图）、漏斗图等'},
                    {label: 'point', value: 'point', desc: '用于绘制点图、折线图中的点等'},
                    {label: 'line', value: 'line', desc: '用于绘制折线图、曲线图、阶梯线图等'},
                    {label: 'area', value: 'area', desc: '用于绘制区域图（面积图）、层叠区域图、区间区域图等'},
                    {label: 'path', value: 'path', desc: '用于绘制路径图，地图上的路径等'},
                    {label: 'polygon', value: 'polygon', desc: '用于绘制色块图（像素图）、热力图、地图等'},
                    {label: 'edge', value: 'edge', desc: '用于绘制流程图、树、弧长连接图、和弦图、桑基图等'},
                    {label: 'heatmap', value: 'heatmap', desc: '用于绘制热力图'},
                    {label: 'schema', value: 'schema', desc: '用于绘制 k 线图，箱型图'}
                ],
                coordinateOptions: [
                    {label: '笛卡尔坐标系', value: 'cartesian'},
                    {label: '极坐标系', value: 'polar'},
                    {label: '螺旋坐标系，基于阿基米德螺旋线', value: 'helix'},
                    {label: '特殊的极坐标系，半径长度固定，仅仅将数据映射到角度', value: 'theta'},
                ],
                shapeOptions: {
                    interval: [
                        {label: 'rect', value: 'rect'},
                        {label: 'hollow-rect', value: 'hollow-rect'},
                        {label: 'line', value: 'line'},
                        {label: 'tick', value: 'tick'},
                        {label: 'funnel', value: 'funnel'},
                        {label: 'pyramid', value: 'pyramid'},
                    ],
                    point: [
                        {label: 'circle', value: 'circle'},
                        {label: 'square', value: 'square'},
                        {label: 'bowtie', value: 'bowtie'},
                        {label: 'diamond', value: 'diamond'},
                        {label: 'hexagon', value: 'hexagon'},
                        {label: 'triangle', value: 'triangle'},
                        {label: 'triangle-down', value: 'triangle-down'},
                        {label: 'hollow-circle', value: 'hollow-circle'},
                        {label: 'hollow-square', value: 'hollow-square'},
                        {label: 'hollow-bowtie', value: 'hollow-bowtie'},
                        {label: 'hollow-diamond', value: 'hollow-diamond'},
                        {label: 'hollow-hexagon', value: 'hollow-hexagon'},
                        {label: 'hollow-triangle', value: 'hollow-triangle'},
                        {label: 'hollow-triangle-down', value: 'hollow-triangle-down'},
                        {label: 'cross', value: 'cross'},
                        {label: 'tick', value: 'tick'},
                        {label: 'plus', value: 'plus'},
                        {label: 'hyphen', value: 'hyphen'},
                        {label: 'line', value: 'line'},
                    ],
                    line: [
                        {label: 'line', value: 'line'},
                        {label: 'dot', value: 'dot'},
                        {label: 'dash', value: 'dash'},
                        {label: 'smooth', value: 'smooth'},
                        {label: 'hv', value: 'hv'},
                        {label: 'vh', value: 'vh'},
                        {label: 'hvh', value: 'hvh'},
                        {label: 'vhv', value: 'vhv'},
                    ],
                    area: [
                        {label: 'area', value: 'area'},
                        {label: 'smooth', value: 'smooth'},
                        {label: 'line', value: 'line'},
                        {label: 'smooth-line', value: 'smooth-line'},
                    ],
                    polygon: [
                        {label: 'polygon', value: 'polygon'}
                    ],
                    schema: [
                        {label: 'box', value: 'box'},
                        {label: 'candle', value: 'candle'}
                    ],
                    edge: [
                        {label: 'line', value: 'line'},
                        {label: 'vhv', value: 'vhv'},
                        {label: 'smooth', value: 'smooth'},
                        {label: 'arc', value: 'arc'},
                    ],
                },
                legendPosition: AntVLegendPosition,
                legendLayout: AntVLegendLayout
            },
            optionActiveKey: 'coordinate',
            easyModel: false
        }
    },
    render() {
        return <div class="antv-config">
            <ysz-list-item v-slots={{
                left: () => '简单模式'
            }}>
                <a-switch size="small" vModel={[this.easyModel, 'checked']} onChange={this.onModelChange} />
            </ysz-list-item>
            <a-tabs vModel={[this.optionActiveKey, 'activeKey']} tab-position="left" size="small">
                {this.easyModel
                    ? <a-tab-pane key="type" tab="种类">
                        <a-row style="height: 120px" gutter={[16,16]}>
                            {easyExample.map(Component => <a-col span={6}><Component onPatch={this.onMergeConfig}/></a-col>)}
                        </a-row>
                    </a-tab-pane>
                    : [
                        <a-tab-pane key="coordinate" tab="坐标系">
                            <ysz-list-item v-slots={{
                                left: () => '类型'
                            }}>
                                <a-radio-group size="small" options={this.store.coordinateOptions} vModel={[this.cfg.type.coordinate.type, 'value']} onChange={this.onChange}/>
                            </ysz-list-item>
                            {this.isPolarCoordinate
                                ? [
                                    <ysz-list-item v-slots={{
                                        left: () => '内径'
                                    }}>
                                        <a-input-number size="small" vModel={[this.cfg.type.coordinate.cfg.innerRadius, 'value']} onChange={this.onChange}/>
                                    </ysz-list-item>,
                                    <ysz-list-item v-slots={{
                                        left: () => '外径'
                                    }}>
                                        <a-input-number size="small" vModel={[this.cfg.type.coordinate.cfg.radius, 'value']} onChange={this.onChange}/>
                                    </ysz-list-item>
                                ]
                                : null
                            }
                            <ysz-list-item v-slots={{
                                left: () => '翻转'
                            }}>
                                <a-switch size="small" vModel={[this.cfg.type.coordinate.transpose, 'checked']} onChange={this.onChange}/>
                            </ysz-list-item>
                        </a-tab-pane>,
                    ]}
                <a-tab-pane key="scale" tab="度量">
                    <a-tabs tab-position="top" size="small">
                        {this.coordinateAxis.map(axis => <a-tab-pane key={axis} tab={axis}>
                            <ysz-list-item v-slots={{
                                left: () => '数据字段'
                            }}>
                                <a-input size="small" vModel={[this.cfg.type.scale[axis].field, 'value']} onChange={this.onChange}/>
                            </ysz-list-item>
                            <ysz-list-item v-slots={{
                                left: () => '自定义类型'
                            }}>
                                <a-select style="width:200px" options={this.store.scaleTypeTypeOptions} size="small" vModel={[this.cfg.type.scale[axis].type.type, 'value']} onChange={this.onChange}/>
                            </ysz-list-item>
                            <ysz-list-item v-slots={{
                                left: () => '标题'
                            }}>
                                <a-input size="small" vModel={[this.cfg.type.scale[axis].alias, 'value']} onChange={this.onChange}/>
                            </ysz-list-item>
                            <ysz-list-item v-slots={{
                                left: () => '前缀'
                            }}>
                                <a-input size="small" vModel={[this.cfg.type.scale[axis].format.prefix, 'value']} onChange={this.onChange}/>
                            </ysz-list-item>
                            <ysz-list-item v-slots={{
                                left: () => '后缀'
                            }}>
                                <a-input size="small" vModel={[this.cfg.type.scale[axis].format.suffix, 'value']} onChange={this.onChange}/>
                            </ysz-list-item>
                        </a-tab-pane>)}
                    </a-tabs>
                </a-tab-pane>
                <a-tab-pane key="facet" tab="分组">
                    <ysz-list-item v-slots={{
                        left: () => '启用'
                    }}>
                        <a-switch size="small" vModel={[this.cfg.type.facet.enable, 'checked']} onChange={this.onChange}/>
                    </ysz-list-item>
                    <ysz-list-item v-slots={{
                        left: () => '分割字段'
                    }}>
                        <a-input size="small" vModel={[this.cfg.type.facet.field, 'value']} onChange={this.onChange}/>
                    </ysz-list-item>
                    <ysz-list-item v-slots={{
                        left: () => '自定义类型'
                    }}>
                        <a-select style="width:200px" options={this.store.facetTypeOptions} size="small" vModel={[this.cfg.type.facet.type, 'value']} onChange={this.onChange}/>
                    </ysz-list-item>
                </a-tab-pane>
                <a-tab-pane key="legend" tab="图例">
                    <a-divider orientation="left">开启</a-divider>
                    <a-switch size="small" vModel={[this.cfg.type.legend.enable, 'checked']} onChange={this.onChange}/>
                    <a-divider orientation="left">位置</a-divider>
                    <a-select size="small" style="width: 200px" options={this.store.legendPosition} vModel={[this.cfg.type.legend.position, 'value']} onChange={this.onChange}/>
                    <a-divider orientation="left">布局方式</a-divider>
                    <a-select size="small" options={this.store.legendLayout} vModel={[this.cfg.type.legend.layout, 'value']} onChange={this.onChange}/>
                    <a-divider orientation="left">是否分页 | 分页触发数量</a-divider>
                    <a-switch size="small" vModel={[this.cfg.type.legend.flipPage, 'checked']} onChange={this.onChange}/>
                    <a-divider type="vertical"/>
                    <a-input-number size="small" vModel={[this.cfg.type.legend.maxRow, 'value']} onChange={this.onChange}/>
                </a-tab-pane>
                <a-tab-pane key="label" tab="标签">
                    <a-divider orientation="left">开启</a-divider>
                    <a-switch size="small" vModel={[this.cfg.type.label.enable, 'checked']} onChange={this.onChange}/>
                    <a-divider orientation="left">字段</a-divider>
                    <a-input size="small" vModel={[this.cfg.type.label.field, 'value']} onChange={this.onChange}/>
                </a-tab-pane>
                <a-tab-pane key="layers" tab="层">
                    <more initCount={this.cfg.type.layers.length} onAdd={payload => {
                        this.cfg.type.layers[payload.count] = AntVConfigDefault.layer()
                        this.onChange()
                        payload.done()
                    }} onRemove={payload => {
                        this.cfg.type.layers = this.cfg.type.layers.filter((v, i) => i !== payload.index)
                        this.onChange()
                        payload.done()
                    }} component={layerIndex => {
                        return <a-tabs tab-position="top" size="small">
                            <a-tab-pane key="type" tab="类型">
                                <a-radio-group size="small" options={this.store.typeOptions} vModel={[this.cfg.type.layers[layerIndex].type, 'value']} onChange={this.onChange}/>
                                <a-alert size="small" message={this.getTypeHelp(this.cfg.type.layers[layerIndex].type)}/>
                            </a-tab-pane>
                            <a-tab-pane key="adjust" tab="层叠">
                                <a-divider orientation="left">开启</a-divider>
                                <a-switch size="small" vModel={[this.cfg.type.layers[layerIndex].adjust.enable, 'checked']} onChange={this.onChange}/>
                                <a-divider orientation="left">类型</a-divider>
                                <a-select size="small" options={this.store.adjustOptions} vModel={[this.cfg.type.layers[layerIndex].adjust.type, 'value']} onChange={this.onChange}/>
                            </a-tab-pane>
                            <a-tab-pane key="cat" tab="分类">
                                <a-tabs tab-position="left" size="small">
                                    <a-tab-pane key="size" tab="大小">
                                        <ysz-list-item v-slots={{
                                            left: () => '启用'
                                        }}>
                                            <a-switch size="small" vModel={[this.cfg.type.layers[layerIndex].cat.size.enable, 'checked']} onChange={this.onChange}/>
                                        </ysz-list-item>
                                        {this.cfg.type.layers[layerIndex].cat.size.enable
                                            ? [
                                                <ysz-list-item v-slots={{
                                                    left: () => '是否统一'
                                                }}>
                                                    <a-switch size="small" vModel={[this.cfg.type.layers[layerIndex].cat.size.single, 'checked']} onChange={this.onChange}/>
                                                </ysz-list-item>,
                                                <ysz-list-item v-slots={{
                                                    left: () => '默认'
                                                }}>
                                                    <a-input size="small" vModel={[this.cfg.type.layers[layerIndex].cat.size.default, 'value']} onChange={this.onChange}/>
                                                </ysz-list-item>,
                                                ...(!this.cfg.type.layers[layerIndex].cat.size.single
                                                    ? [
                                                        <a-divider orientation="left">更多</a-divider>,
                                                        <ysz-list-item v-slots={{
                                                            left: () => '分类字段'
                                                        }}>
                                                            <a-input size="small" vModel={[this.cfg.type.layers[layerIndex].cat.size.field, 'value']} onChange={this.onChange}/>
                                                        </ysz-list-item>,
                                                        <ysz-list-item v-slots={{
                                                            left: () => '最小值'
                                                        }}>
                                                            <a-input size="small" vModel={[this.cfg.type.layers[layerIndex].cat.size.enum[0], 'value']} onChange={this.onChange}/>
                                                        </ysz-list-item>,
                                                        <ysz-list-item v-slots={{
                                                            left: () => '最大值'
                                                        }}>
                                                            <a-input size="small" vModel={[this.cfg.type.layers[layerIndex].cat.size.enum[1], 'value']} onChange={this.onChange}/>
                                                        </ysz-list-item>
                                                    ] : [])
                                            ] : null}
                                    </a-tab-pane>
                                    <a-tab-pane key="shape" tab="形状">
                                        <ysz-list-item v-slots={{
                                            left: () => '启用'
                                        }}>
                                            <a-switch size="small" vModel={[this.cfg.type.layers[layerIndex].cat.shape.enable, 'checked']} onChange={this.onChange}/>
                                        </ysz-list-item>
                                        {this.cfg.type.layers[layerIndex].cat.shape.enable
                                            ? [
                                                <ysz-list-item v-slots={{
                                                    left: () => '是否统一'
                                                }}>
                                                    <a-switch size="small" vModel={[this.cfg.type.layers[layerIndex].cat.shape.single, 'checked']} onChange={this.onChange}/>
                                                </ysz-list-item>,
                                                <ysz-list-item v-slots={{
                                                    left: () => '默认'
                                                }}>
                                                    <a-select size="small" options={this.getShapeOptions(this.cfg.type.layers[layerIndex].type, '')} vModel={[this.cfg.type.layers[layerIndex].cat.shape.default, 'value']} onChange={this.onChange}/>
                                                </ysz-list-item>,
                                                ...(!this.cfg.type.layers[layerIndex].cat.shape.single
                                                    ? [
                                                        <a-divider orientation="left">更多</a-divider>,
                                                        <ysz-list-item v-slots={{
                                                            left: () => '分类字段'
                                                        }}>
                                                            <a-input size="small" vModel={[this.cfg.type.layers[layerIndex].cat.shape.field, 'value']} onChange={this.onChange}/>
                                                        </ysz-list-item>,
                                                        <ysz-list-item v-slots={{
                                                            left: () => '额外'
                                                        }}>
                                                            <a-select size="small" style="width: 80%" mode="multiple" options={this.getShapeOptions(this.cfg.type.layers[layerIndex].type, this.cfg.type.layers[layerIndex].cat.shape.default)} vModel={[this.cfg.type.layers[layerIndex].cat.shape.enum, 'value']} onChange={this.onChange}/>
                                                        </ysz-list-item>
                                                    ]
                                                    : [])
                                            ] : null}
                                    </a-tab-pane>
                                    <a-tab-pane key="color" tab="颜色">
                                        <ysz-list-item v-slots={{
                                            left: () => '启用'
                                        }}>
                                            <a-switch size="small" vModel={[this.cfg.type.layers[layerIndex].cat.color.enable, 'checked']} onChange={this.onChange}/>
                                        </ysz-list-item>
                                        {this.cfg.type.layers[layerIndex].cat.color.enable
                                            ? [
                                                <ysz-list-item v-slots={{
                                                    left: () => '是否统一'
                                                }}>
                                                    <a-switch size="small" vModel={[this.cfg.type.layers[layerIndex].cat.color.single, 'checked']} onChange={this.onChange}/>
                                                </ysz-list-item>,
                                                <ysz-list-item v-slots={{
                                                    left: () => '默认'
                                                }}>
                                                    <color-pick vModel={[this.cfg.type.layers[layerIndex].cat.color.default, 'value']} onChange={this.onChange}/>
                                                </ysz-list-item>,
                                                ...(!this.cfg.type.layers[layerIndex].cat.color.single
                                                    ? [
                                                        <a-divider orientation="left">更多</a-divider>,
                                                        <ysz-list-item v-slots={{
                                                            left: () => '分类字段'
                                                        }}>
                                                            <a-input size="small" vModel={[this.cfg.type.layers[layerIndex].cat.color.field, 'value']} onChange={this.onChange}/>
                                                        </ysz-list-item>,
                                                        <ysz-list-item v-slots={{
                                                            left: () => '额外'
                                                        }}>
                                                            <more initCount={this.cfg.type.layers[layerIndex].cat.color.enum.length} onAdd={payload => {
                                                                this.cfg.type.layers[layerIndex].cat.color.enum[payload.count] = AntVConfigDefault.layerCatColorDefault()
                                                                this.onChange()
                                                                payload.done()
                                                            }} onRemove={payload => {
                                                                this.cfg.type.layers[layerIndex].cat.color.enum = this.cfg.type.layers[layerIndex].cat.color.enum.filter((v, i) => i !== payload.index)
                                                                this.onChange()
                                                                payload.done()
                                                            }} component={index => {
                                                                return <color-pick vModel={[this.cfg.type.layers[layerIndex].cat.color.enum[index], 'value']} onChange={this.onChange}/>
                                                            }} />
                                                        </ysz-list-item>
                                                    ] : [])
                                            ] : null}
                                    </a-tab-pane>
                                </a-tabs>
                            </a-tab-pane>
                        </a-tabs>
                    }} />
                </a-tab-pane>
                <a-tab-pane key="data" tab="数据">
                    <inputChoose inputs={this.cfg.common.input}
                         value={this.cfg.type.dataIndex}
                         onChange={value => {
                             this.cfg.type.dataIndex = value
                             this.onChange()
                        }} />
                </a-tab-pane>
            </a-tabs>
        </div>
    },
    computed: {
        coordinateAxis() {
            return AntVCoordinateAxis(this.cfg.type.coordinate.type, this.cfg.type.coordinate.transpose)
        },
        isPolarCoordinate() {
            return AntVIsPolarCoordinate(this.cfg.type.coordinate.type)
        }
    },
    methods: {
        transformConfig() {
            try {
                const blockCfg = JSON.parse(this.config)
                blockCfg.type = AntVConfigParse(blockCfg.type)
                this.cfg = blockCfg
            }catch(e) {
                console.log('AntV config parse failed in antv-config', e, this.config)
            }
        },
        onModelChange() {
            this.optionActiveKey = this.easyModel ? 'type' : 'coordinate'
        },
        onMergeConfig(payload) {
            this.cfg.type = AntVConfigParse(util.merge(this.cfg.type, payload))
            this.onChange()
        },
        getTypeHelp(typ) {
            let type = this.store.typeOptions.find(type => {
                return type.value === typ
            })

            return !!type ? type.desc : ''
        },
        getShapeOptions(typ, def) {
            const options = this.store.shapeOptions[typ]
            return (!options ? [] : options).filter(v => v !== def)
        }
    }
}
</script>

<style lang="scss" scoped>
.antv-config {
    &:deep .ant-tabs-tab {
        padding: 0; margin: 0; padding-right: 1rem; margin-bottom: 1rem
    }
    &:deep .ant-tabs-left-content ant-tabs-content {
        padding-left: 0
    }
}
</style>