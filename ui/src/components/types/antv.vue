<script lang="tsx">
import { Chart } from '@antv/g2'
import {AntVConfig} from 'type'
export default {
    props: {
        data: {
            type: Array,
            default() {
                return []
            }
        },
        config: {
            type: String,
            default: ""
        }
    },
    render() {
        return <div ref="chart" style="width: 100%;"></div>
    },
    data() {
        return {
            cfg: AntVConfig()
        }
    },
    chart: null,
    created() {
        this.parse()
    },
    watch: {
        data(val) {
            if(this.chart) {
                this.chart.data(val)
                this.chart.render()
                this.renderAfter()
            }
        }
    },
    mounted() {
        this.$nextTick(() => {
            const chart = this.chart = new Chart({
                container: this.$refs.chart,
                autoFit: true,
            })

            // coordinate
            if(this.cfg.coordinate.type) {
                chart.coordinate(this.cfg.coordinate.type)
            }

            if(this.cfg.coordinate.transpose) {
                chart.getCoordinate().transpose()
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
                .position(`${this.cfg.scale.x.field}*${this.cfg.scale.y.field}`)

            let hasColorCat = false
            this.cfg.cats.forEach(cat => {
                switch (cat.type) {
                    case 'size':
                        // string, array
                        geometry.size(cat.field, cat.enum)
                        break
                    case 'shape':
                        geometry.shape(cat.field, cat.enum)
                        break
                    case 'color':
                        hasColorCat = true
                        geometry.color(cat.field, cat.enum)
                        break
                }
            })

            if(!hasColorCat && this.cfg.color) {
                geometry.color(this.cfg.color)
            }

            chart.data(this.data)

            // TODO: facet view support
            chart.render()
            this.renderAfter()
        })
    },
    methods: {
        parse() {

        },
        renderAfter() {
            const e = document.createEvent('Event')
            e.initEvent('resize', true, true)
            window.dispatchEvent(e)
        }
    }
}
</script>