<script lang="tsx">
import configMixin from '../config-mixin'
import {
    ViewBlockType, TableConfig, TableConfigParse,
} from 'type'

import inputChoose from './input-choose.vue'
function SizeRange(start, end, step) {
    let rs = []
    for(let i = start; i <= end; i += step) {
        rs.push(parseFloat(i.toFixed(2)))

        if (i != end && i + step > end) {
            rs.push(parseFloat(end.toFixed(2)))
        }
    }
    return rs
}
export default {
    components: {inputChoose},
    mixins: [configMixin],
    data() {
        return {
            cfg: {
                ...ViewBlockType().config,
                type: TableConfig()
            },
            store: {
                scrollPrecentOnceOptions: SizeRange(0, 100, 10).map(val => ({label: val, value: val})),
                scrollCycleOptions: SizeRange(1, 20, 1).map(val => ({label: val, value: val}))
            }
        }
    },
    render() {
        return <div>
            <ysz-list-item v-slots={{
                left: () => '滚动百分比'
            }}>
                <a-select options={this.store.scrollPrecentOnceOptions} vModel={[this.cfg.type.scrollPrecentOnce, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '滚动周期(s)'
            }}>
                <a-select options={this.store.scrollCycleOptions} vModel={[this.cfg.type.scrollCycle, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '表头背景色'
            }}>
                <color-pick vModel={[this.cfg.type.headerBGC, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '表头字体颜色'
            }}>
                <color-pick vModel={[this.cfg.type.headerC, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '数据'
            }}>
                <inputChoose inputs={this.cfg.common.input}
                             value={this.cfg.type.dataIndex}
                             onChange={value => {
                                 this.cfg.type.dataIndex = value
                                 this.onChange()
                             }} />
            </ysz-list-item>
        </div>
    },
    methods: {
        transformConfig() {
            try {
                const blockCfg = JSON.parse(this.config)
                blockCfg.type = TableConfigParse(blockCfg.type)
                this.cfg = blockCfg
            }catch(e) {
                console.log('static table config parse failed', e, this.config)
            }
        },
    }
}
</script>