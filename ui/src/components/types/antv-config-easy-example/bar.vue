<script lang="tsx">
import {Chart} from '@antv/g2'
import data from './data.js'
export default {
    render() {
        return <div style="width:100%; height: 100%" ref="chart"/>
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
            chart.interval().position('x*y')
            chart.data(data).render()
        })
    }
}
</script>