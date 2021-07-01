<script lang="tsx">
import { BarChartOutlined, ReadOutlined } from '@ant-design/icons-vue';
import * as componentType from '@/components/types/type.js'
import {AntVConfig} from 'type'
import {mapState} from "vuex";

const typeItem = {
    props: {
        icon: {
            type: Function,
            default() {
                return () => null
            }
        },
        active: {
            type: Boolean,
            default: false
        }
    },
    render() {
        const label = this.$slots.default()
        const attrs = {
            class: ['type-item-wrap']
        }

        if(this.active) {
            attrs.class.push('active')
        }

        return <div {...attrs}>
            <div style="flex: 0 0 60%; text-align:center">{this.icon()}</div>
            <div style="flex: 0 0 40%; text-align:center">{label}</div>
        </div>
    }
}

export default {
    render() {
        // TODO: tranform to config list
        return <a-row gutter={[16,16]}>
            <a-col span={4}><typeItem active={this.currentType === componentType.AntV} onClick={this.onAntVTypeClick} icon={() => <BarChartOutlined style="font-size:1.4rem"/>}>图表</typeItem></a-col>
            <a-col span={4}><typeItem active={this.currentType === componentType.StaticText} onClick={this.onStaticTextTypeClick} icon={() => <ReadOutlined style="font-size:1.4rem"/>}>静态文本</typeItem></a-col>
            <a-col span={4}><typeItem active={this.currentType === componentType.DynamicText} onClick={this.onDynamicTextTypeClick} icon={() => <ReadOutlined style="font-size:1.4rem" spin/>}>动态文本</typeItem></a-col>
        </a-row>
    },
    computed: {
        ...mapState('config', {
            currentType: state => state.block.type
        })
    },
    methods: {
        onAntVTypeClick() {
            this.mixinSetConfigTypeAndConfig(componentType.AntV, JSON.stringify(AntVConfig()))
        },
        onStaticTextTypeClick() {
            this.mixinSetConfigTypeAndConfig(componentType.StaticText, JSON.stringify(AntVConfig()))
        },
        onDynamicTextTypeClick() {
            this.mixinSetConfigTypeAndConfig(componentType.DynamicText, JSON.stringify(AntVConfig()))
        }
    }
}
</script>

<style lang="scss" scoped>
.type-item-wrap {
    display: flex; flex-direction: column; height: 100%; cursor: pointer;
    &.active {
        border: 1px dashed pink; border-radius: 6px; padding: 2px
    }
}
</style>