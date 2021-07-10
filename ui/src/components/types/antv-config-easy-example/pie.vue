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
            chart.coordinate('theta', {
                radius: 0.75
            })
            chart.tooltip(false).legend(false)
            // TODO: adjust,fit-view
            chart.interval().adjust('stack').position('y')
                .color('x', ['#063d8a', '#1770d6', '#47abfc', '#38c060'])
            chart.data(data).render()
        })
    },
    emits: ['patch'],
    methods: {
        onPatch() {
            this.$emit('patch', {
                coordinate: {
                    type: 'theta',
                    cfg: {
                        radius: 0.75
                    }
                },
                layers: [
                    {
                        type: 'interval',
                        adjust: {
                            enable: true,
                            type: 'stack'
                        },
                        cat: {
                            color: {
                                single: false,
                                enum: ['#063d8a', '#1770d6', '#47abfc', '#38c060']
                            }
                        }
                    }
                ]
            })
        }
    }
}
</script>