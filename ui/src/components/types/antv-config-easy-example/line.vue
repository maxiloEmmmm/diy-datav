<script lang="tsx">
import {Chart} from '@antv/g2'
import data from './data.js'
export default {
    render() {
        return <div style="width:100%; height: 100%" ref="chart" onClick={this.onPatch}/>
    },
    mounted() {
        this.$nextTick(() => {
            const chart = new Chart({
                container: this.$refs.chart,
                autoFit: true,
            })
            chart.tooltip(false).legend(false)
            chart.axis('x', {label: null})
            chart.axis('y', {label: null})
            chart.line().position('x*y')
            chart.data(data).render()
        })
    },
    emits: ['patch'],
    methods: {
        onPatch() {
            this.$emit('patch', {
                coordinate: {
                    type: 'cartesian'
                },
                layers: [
                    {
                        type: 'line',
                    }
                ]
            })
        }
    }
}
</script>