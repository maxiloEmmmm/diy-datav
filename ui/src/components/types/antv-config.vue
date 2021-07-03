<script lang="tsx">
import {AntVConfig, AntVConfigParse, ViewBlockType} from 'type/index.js'
import configMixin from './config-mixin'
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
                ]
            },
            optionActiveKey: 'type',
        }
    },
    render() {
        return <div class="antv-config">
            <a-tabs vModel={[this.optionActiveKey, 'activeKey']} tab-position="left" size="small">
                <a-tab-pane key="type" tab="类型">
                    <a-radio-group options={this.store.typeOptions} vModel={[this.cfg.type.type, 'value']} onChange={this.onChange}/>
                    <a-alert message={this.typeHelp}/>
                </a-tab-pane>
                <a-tab-pane key="coordinate.type" tab="坐标系">
                    <a-divider orientation="left">类型</a-divider>
                    <a-radio-group options={this.store.coordinateOptions} vModel={[this.cfg.type.coordinate.type, 'value']} onChange={this.onChange}/>
                    <a-divider orientation="left">翻转</a-divider>
                    <a-switch vModel={[this.cfg.type.coordinate.transpose, 'checked']} onChange={this.onChange}/>
                </a-tab-pane>
                <a-tab-pane key="cats" tab="分类">
                    <a-tabs tab-position="top" size="small">
                        <a-tab-pane key="size" tab="大小">
                        </a-tab-pane>
                        <a-tab-pane key="shape" tab="形状">
                        </a-tab-pane>
                        <a-tab-pane key="color" tab="颜色">
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