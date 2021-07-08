<script lang="tsx">
import {Chart} from '@antv/g2'

const data = [
    { year: '1991', value: 8, v2: 1, v3: 2, v4: 1 },
    { year: '1991', value: 3, v2: 1, v3: 2, v4: 1 },
    { year: '1992', value: 4, v2: 12, v3: 7, v4: 1.2 },
    { year: '1993', value: 3.5, v2: 14, v3: 7, v4: 0 },
    { year: '1994', value: 5, v2: 13, v3: 2, v4: 1 },
    { year: '1995', value: 4.9, v2: 16, v3: 7, v4: 1 },
    { year: '1996', value: 6, v2: 17, v3: 7, v4: 1 },
    { year: '1997', value: 7, v2: 8, v3: 17, v4: 12 },
    { year: '1998', value: 9, v2: 19, v3: 1, v4: 5 },
    { year: '1999', value: 13, v2: 20, v3: 7, v4: 15 },
]

export default {
    render(){
        return <div ref="chart" style="height:100vh;width:100vw; padding:2rem"></div>
    },
    data() {
        return {
            a: {b: 1}
        }
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
            chart.axis('v3', false)
            chart.axis('v4', false)
            chart.axis('v2', false)

            chart.interval().adjust('stack').position('year*value');
            chart.area().adjust('stack').position('year*value');

            chart.render()
        })
    }
}

</script>