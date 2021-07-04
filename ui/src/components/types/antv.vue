<script lang="tsx">
import { Chart } from '@antv/g2'
import {AntVConfig, AntVConfigParse} from 'type/types/index.js'
import util from 'pkg/util'
export default {
    props: {
        data: {
            type: Array,
            default() {
                return []
            }
        },
        config: {
            type: Object,
            default() {
                return AntVConfig()
            }
        }
    },
    render() {
        return <div ref="chart" style="width: 100%; height: 100%"></div>
    },
    data() {
        return {
            cfg: AntVConfig()
        }
    },
    chart: null,
    watch: {
        data: {
            deep: true,
            handler(val) {
                if(this.chart) {
                    this.chart.data(this.getData())
                    this.chart.render()
                    this.renderAfter()
                }
            }
        },
        config: {
            deep: true,
            handler() {
                // config maybe change after chart init than in dom
                this.$nextTick(() => {
                    if(this.chart) {
                        this.chart.destroy()
                    }
                    this.render()
                })
            }
        }
    },
    mounted() {
        this.render()
    },
    methods: {
        parse() {
            try {
                this.cfg = AntVConfigParse(this.config)
            }catch(e) {
                console.log('AntV config parse failed in antv-config', e, this.config)
            }
        },
        getData() {
            let data = []
            if(util.isArray(this.data)) {
                data = this.data[this.cfg.dataIndex]
            }
            return data === undefined ? [] : data
        },
        render() {
            this.parse()
            this.$nextTick(() => {
                const chart = this.chart = new Chart({
                    container: this.$refs.chart,
                    autoFit: true,
                })

                // coordinate
                const coordinate = chart.coordinate(this.cfg.coordinate.type)

                if(this.cfg.coordinate.transpose) {
                    coordinate.transpose()
                }

                //    scale
                chart.scale({
                    [this.cfg.scale.x.field]: {
                        alias: this.cfg.scale.x.alias,
                        formatter: value => `${this.cfg.scale.x.format.prefix}${value}${this.cfg.scale.x.format.suffix}`
                    },
                    [this.cfg.scale.y.field]: {
                        alias: this.cfg.scale.y.alias,
                        formatter: value => `${this.cfg.scale.y.format.prefix}${value}${this.cfg.scale.y.format.suffix}`
                    }
                })

                // TODO: axis


                // TODO: y maybe not define
                // TODO: more type
                const geometry = chart[this.cfg.type]()
                    .position({
                        fields: [this.cfg.scale.x.field, this.cfg.scale.y.field]
                    })

                // TODO: 暂时只支持单字段
                this.cfg.cat.color.single
                    ? geometry.color(this.cfg.cat.color.default)
                    : geometry.color({
                        fields: [this.cfg.scale.x.field],
                        values: this.cfg.cat.color.enum
                    })

                this.cfg.cat.size.single
                    ? geometry.size(this.cfg.cat.size.default)
                    : geometry.size({
                        fields: [this.cfg.scale.x.field],
                        values: this.cfg.cat.size.enum
                    })

                this.cfg.cat.shape.single
                    ? geometry.shape(this.cfg.cat.shape.default)
                    : geometry.shape({
                        fields: [this.cfg.scale.x.field],
                        values: this.cfg.cat.shape.enum
                    })

                chart.data(this.getData())

                // TODO: facet view support
                chart.render()
                this.renderAfter()
            })
        },
        renderAfter() {
            this.mixinDispatchWindowResize()
        }
    }
}
</script>