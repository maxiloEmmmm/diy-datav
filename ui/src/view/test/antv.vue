<script lang="tsx">
import {Chart} from '@antv/g2'

const data = [
    { year: '1991', value: 3 },
    { year: '1992', value: 4 },
    { year: '1993', value: 3.5 },
    { year: '1994', value: 5 },
    { year: '1995', value: 4.9 },
    { year: '1996', value: 6 },
    { year: '1997', value: 7 },
    { year: '1998', value: 9 },
    { year: '1999', value: 13 },
]

export default {
    render(){
        return <div ref="chart"></div>
    },
    mounted() {
        this.$nextTick(() => {
            const chart = new Chart({
                container: this.$refs.chart,
                autoFit: true,
                height: 500,
            })

            chart.data(data);
            chart.scale({
                year: {
                    range: [0, 0.5],
                },
                value: {
                    min: 0,
                    nice: true,
                },
            });

            chart.tooltip({
                showCrosshairs: true, // 展示 Tooltip 辅助线
                shared: true,
            });

            chart.line().position('year*value').label('value').shape('smooth');;
            chart.point().position('year*value');

            chart.render()
        })
    }
}

</script>