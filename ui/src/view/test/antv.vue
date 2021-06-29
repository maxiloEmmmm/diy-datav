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

        // shape 图形的形状
        // label 直接映射在图形上的文字
        // scale.*?.range 输出域、值域，表示在绘图范围内可用于绘制的范围
        // scale.*?.nice 自动调整定义域的最大最小值
        // 视觉通道设计 https://g2.antv.vision/zh/docs/manual/concepts/grammar-of-graphics#g2-%E8%A7%86%E8%A7%89%E9%80%9A%E9%81%93%E7%9A%84%E8%AE%BE%E8%AE%A1
        // 分类/连续 https://g2.antv.vision/zh/docs/manual/tutorial/scale#%E5%BA%A6%E9%87%8F%E7%B1%BB%E5%9E%8B
        this.$nextTick(() => {
            const chart = new Chart({
                container: this.$refs.chart,
                autoFit: true,
            })

            chart.data(data);
            chart.scale({
                year: {
                    nice: true
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

            chart.line().position('year*value').label('value').shape('smooth');
            chart.point().position('year*value');

            chart.render()
        })
    }
}

</script>