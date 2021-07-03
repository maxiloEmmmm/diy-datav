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
                ]
            }
        }
    },
    render() {
        return <div>
            <a-radio-group options={this.store.typeOptions} value={this.cfg.type.type} onChange={this.onTypeChange}/>
            <a-alert message={this.typeHelp}></a-alert>
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
        onTypeChange(val) {
            this.cfg.type.type = val.target.value
            this.onChange()
        },
    }
}
</script>