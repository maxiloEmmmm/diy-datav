<script lang="tsx">
import configMixin from '../config-mixin'
import {
    ViewBlockType, TextConfig, TextConfigParse,
    TextHorizontalAlignOptions, TextVerticalAlignOptions
} from 'type'

function SizeRange(end, step) {
    let rs = []
    for(let i = 1; i <= end; i += step) {
        rs.push(parseFloat(i.toFixed(2)))

        if (i != end && i + step > end) {
            rs.push(parseFloat(end.toFixed(2)))
        }
    }
    return rs
}

import inputChoose from './input-choose.vue'
export default {
    components: {inputChoose},
    mixins: [configMixin],
    data() {
        return {
            cfg: {
                ...ViewBlockType().config,
                type: TextConfig()
            },
            store: {
                size: SizeRange(5, 0.04).map(val => ({label: val, value: val})),
                horizontalAlignOptions: TextHorizontalAlignOptions,
                verticalAlignOptions: TextVerticalAlignOptions
            }
        }
    },
    render() {
        return <div>
            <ysz-list-item v-slots={{
                left: () => '内容'
            }}>
                <a-input size="small" vModel={[this.cfg.type.text, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '水平排列'
            }}>
                <a-radio-group options={this.store.horizontalAlignOptions} size="small" vModel={[this.cfg.type.align.h, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '垂直排列'
            }}>
                <a-radio-group options={this.store.verticalAlignOptions} size="small" vModel={[this.cfg.type.align.v, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '加粗'
            }}>
                <a-switch size="small" vModel={[this.cfg.type.bold, 'checked']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '倾斜'
            }}>
                <a-switch size="small" vModel={[this.cfg.type.italic, 'checked']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '下划线'
            }}>
                <a-switch size="small" vModel={[this.cfg.type.underline, 'checked']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '上标'
            }}>
                <a-input size="small" vModel={[this.cfg.type.sup, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '下标'
            }}>
                <a-input size="small" vModel={[this.cfg.type.sub, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '颜色'
            }}>
                <color-pick size="small" vModel={[this.cfg.type.color, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '大小'
            }}>
                <a-select style="width:200px" options={this.store.size} size="small" vModel={[this.cfg.type.size, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '动态数据'
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
                blockCfg.type = TextConfigParse(blockCfg.type)
                this.cfg = blockCfg
            }catch(e) {
                console.log('static text config parse failed', e, this.config)
            }
        },
    }
}
</script>