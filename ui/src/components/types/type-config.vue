<script lang="tsx">
import { BarChartOutlined, ReadOutlined } from '@ant-design/icons-vue';
import * as componentType from '@/components/types/type.js'
import {AntVConfig} from 'type'
import {mapState} from "vuex";
import SelectItem from "../select-item.vue";

export default {
    components: {SelectItem},
    render() {
        // TODO: tranform to config list
        return <a-row gutter={[16,16]}>
            <a-col span={4}><select-item active={this.currentType === componentType.AntV} onClick={this.onAntVTypeClick} icon={() => <BarChartOutlined style="font-size:1.4rem"/>}>图表</select-item></a-col>
            <a-col span={4}><select-item active={this.currentType === componentType.StaticText} onClick={this.onStaticTextTypeClick} icon={() => <ReadOutlined style="font-size:1.4rem"/>}>静态文本</select-item></a-col>
            <a-col span={4}><select-item active={this.currentType === componentType.DynamicText} onClick={this.onDynamicTextTypeClick} icon={() => <ReadOutlined style="font-size:1.4rem" spin/>}>动态文本</select-item></a-col>
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