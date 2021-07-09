<script lang="tsx">
import { Chart } from '@antv/g2'
import {AntVConfig, AntVConfigParse, AntVCoordinateAxis} from 'type/types/index.js'
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

                ['x', 'y', 'z'].forEach(scale => {
                    chart.scale(this.cfg.scale[scale].field, {
                        alias: this.cfg.scale[scale].alias,
                        formatter: value => `${this.cfg.scale[scale].format.prefix}${value}${this.cfg.scale[scale].format.suffix}`,
                        nice: true,
                    })

                    if(this.cfg.scale[scale].alias !== "") {
                        chart.axis(this.cfg.scale[scale].field, {title: {}})
                    }
                })
                // TODO: axis

                let fields = []
                Object.keys(this.cfg.scale).forEach(scale => {
                    if (AntVCoordinateAxis(this.cfg.coordinate.type, this.cfg.coordinate.transpose).includes(scale)) {
                        fields.push(this.cfg.scale[scale].field)
                    }
                })

                // if facet support
                if(this.cfg.facet.enable) {
                    chart.facet(this.cfg.facet.type, {
                        fields: [this.cfg.facet.field],
                        eachView: (view) => {
                            this.viewSet(view, fields)
                        }
                    })
                }else {
                    this.viewSet(chart, fields)
                }

                chart.data(this.getData())

                chart.render()
                this.renderAfter()
            })
        },
        viewSet(view, fields) {
            this.cfg.layers.forEach(layer => {
                const geometry = view[layer.type]()
                    .position({
                        fields,
                    })

                if(layer.adjust.enable) {
                    geometry.adjust(layer.adjust.type)
                }

                layer.cat.color.single
                    ? geometry.color(layer.cat.color.default)
                    : geometry.color({
                        fields: [layer.cat.color.field],
                        values: layer.cat.color.enum
                    })

                layer.cat.size.single
                    ? geometry.size(layer.cat.size.default)
                    : geometry.size({
                        fields: [layer.cat.size.field],
                        values: layer.cat.size.enum
                    })

                layer.cat.shape.single
                    ? geometry.shape(layer.cat.shape.default)
                    : geometry.shape({
                        fields: [layer.cat.shape.field],
                        values: layer.cat.shape.enum
                    })
            })
        },
        renderAfter() {
            this.mixinDispatchWindowResize()
        }
    }
}
</script>