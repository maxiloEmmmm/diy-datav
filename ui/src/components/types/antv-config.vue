<script lang="tsx">
import {AntVConfig, AntVConfigParse, AntVConfigDefault, ViewBlockType} from 'type/index.js'
import configMixin from './config-mixin'
import easyExample from './antv-config-easy-example'
export default {
    mixins: [configMixin],
    data() {
        return {
            cfg: {
                ...ViewBlockType().config,
                type: AntVConfig()
            },
            store: {
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
                    ]
                }
            },
            optionActiveKey: 'type',
            easyModel: false
        }
    },
    render() {
        return <div class="antv-config">
            <ysz-list-item v-slots={{
                left: () => '简单模式'
            }}>
                <a-switch vModel={[this.easyModel, 'checked']} />
            </ysz-list-item>
            <a-tabs vModel={[this.optionActiveKey, 'activeKey']} tab-position="left" size="small">
                {this.easyModel
                    ? <a-tab-pane key="type" tab="种类">
                        <a-row style="height: 120px" gutter={[16,16]}>
                            {easyExample.map(Component => <a-col span={6}><Component/></a-col>)}
                        </a-row>
                    </a-tab-pane>
                    : [
                        <a-tab-pane key="type" tab="类型">
                            <a-radio-group options={this.store.typeOptions} vModel={[this.cfg.type.type, 'value']} onChange={this.onChange}/>
                            <a-alert message={this.typeHelp}/>
                        </a-tab-pane>,
                        <a-tab-pane key="coordinate.type" tab="坐标系">
                            <a-divider orientation="left">类型</a-divider>
                            <a-radio-group options={this.store.coordinateOptions} vModel={[this.cfg.type.coordinate.type, 'value']} onChange={this.onChange}/>
                            <a-divider orientation="left">翻转</a-divider>
                            <a-switch vModel={[this.cfg.type.coordinate.transpose, 'checked']} onChange={this.onChange}/>
                        </a-tab-pane>
                    ]}
                <a-tab-pane key="cats" tab="分类">
                    <a-tabs tab-position="top" size="small">
                        <a-tab-pane key="size" tab="大小">
                            <ysz-list-item v-slots={{
                                left: () => '是否统一'
                            }}>
                                <a-switch vModel={[this.cfg.type.cat.size.single, 'checked']} onChange={this.onChange}/>
                            </ysz-list-item>
                            <ysz-list-item v-slots={{
                                left: () => '默认'
                            }}>
                                <a-input vModel={[this.cfg.type.cat.size.default, 'value']} onChange={this.onChange}/>
                            </ysz-list-item>
                            {!this.cfg.type.cat.size.single
                                ? [
                                    <a-divider orientation="left">更多</a-divider>,
                                    <ysz-list-item v-slots={{
                                        left: () => '最小值'
                                    }}>
                                        <a-input vModel={[this.cfg.type.cat.size.enum[0], 'value']} onChange={this.onChange}/>
                                    </ysz-list-item>,
                                    <ysz-list-item v-slots={{
                                        left: () => '最大值'
                                    }}>
                                        <a-input vModel={[this.cfg.type.cat.size.enum[1], 'value']} onChange={this.onChange}/>
                                    </ysz-list-item>
                                ] : null}
                        </a-tab-pane>
                        <a-tab-pane key="shape" tab="形状">
                            <ysz-list-item v-slots={{
                                left: () => '是否统一'
                            }}>
                                <a-switch vModel={[this.cfg.type.cat.shape.single, 'checked']} onChange={this.onChange}/>
                            </ysz-list-item>
                            <ysz-list-item v-slots={{
                                left: () => '默认'
                            }}>
                                <a-select options={this.shapeOptions} vModel={[this.cfg.type.cat.shape.default, 'value']} onChange={this.onChange}/>
                            </ysz-list-item>
                            {!this.cfg.type.cat.shape.single
                                ? <ysz-list-item v-slots={{
                                    left: () => '更多'
                                }}>
                                    <a-select style="width: 80%" mode="multiple" options={this.shapeOptions} vModel={[this.cfg.type.cat.shape.enum, 'value']} onChange={this.onChange}/>
                                </ysz-list-item>
                                : null}
                        </a-tab-pane>
                        <a-tab-pane key="color" tab="颜色">
                            <ysz-list-item v-slots={{
                                left: () => '是否统一'
                            }}>
                                <a-switch vModel={[this.cfg.type.cat.color.single, 'checked']} onChange={this.onChange}/>
                            </ysz-list-item>
                            <ysz-list-item v-slots={{
                                left: () => '默认'
                            }}>
                                <color-pick vModel={[this.cfg.type.cat.color.default, 'value']} onChange={this.onChange}/>
                            </ysz-list-item>
                            {!this.cfg.type.cat.color.single
                                ? <ysz-list-item v-slots={{
                                    left: () => '更多'
                                }}>
                                    <more initCount={this.cfg.type.cat.color.enum.length} onAdd={payload => {
                                        this.cfg.type.cat.color.enum[payload.count] = AntVConfigDefault.catColorDefault()
                                        payload.done()
                                    }} onRemove={payload => {
                                        this.cfg.type.cat.color.enum = this.cfg.type.cat.color.enum.filter((v, i) => i !== payload.index)
                                        payload.done()
                                    }} component={index => {
                                        return <color-pick vModel={[this.cfg.type.cat.color.enum[index], 'value']} onChange={this.onChange}/>
                                    }} />
                                </ysz-list-item> : null}
                        </a-tab-pane>
                    </a-tabs>
                </a-tab-pane>
            </a-tabs>
        </div>
    },
    computed: {
        typeHelp() {
            let type = this.store.typeOptions.find(type => {
               return type.value === this.cfg.type.type
            })

            return !!type ? type.desc : ''
        },
        shapeOptions() {
            const options = this.store.shapeOptions[this.cfg.type.type]
            return (!options ? [] : options).filter(v => v !== this.cfg.type.cat.shape.default)
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
    }
}
</script>

<style lang="scss" scoped>
.antv-config {
    &::v-deep .ant-tabs-tab {
        padding: 0; margin: 0; padding-right: 1rem; margin-bottom: 1rem
    }
}


</style>