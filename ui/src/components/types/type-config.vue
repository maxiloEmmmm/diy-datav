<script lang="tsx">
import { BarChartOutlined, ReadOutlined } from '@ant-design/icons-vue';
import * as componentType from '@/components/types/type.js'
import {AntVConfig, ViewBlockType, TextConfig, TableConfig, GridConfig} from 'type'
import {mapState} from "vuex";
import SelectItem from "../select-item.vue";

export default {
    components: {SelectItem},
    render() {
        // TODO: tranform to config list
        return <a-row gutter={[16,16]}>
            <a-col span={4}><select-item active={this.currentType === componentType.AntV} onClick={this.onAntVTypeClick} icon={() => <BarChartOutlined style="font-size:1.4rem"/>}>图表</select-item></a-col>
            <a-col span={4}><select-item active={this.currentType === componentType.Text} onClick={this.onTextTypeClick} icon={() => <ReadOutlined style="font-size:1.4rem"/>}>静态文本</select-item></a-col>
            <a-col span={4}><select-item active={this.currentType === componentType.Table} onClick={this.onTableTypeClick} icon={() => <ReadOutlined style="font-size:1.4rem"/>}>表格</select-item></a-col>
            <a-col span={4}><select-item active={this.currentType === componentType.Grid} onClick={this.onGridTypeClick} icon={() => <ReadOutlined style="font-size:1.4rem"/>}>分栏</select-item></a-col>
        </a-row>
    },
    computed: {
        ...mapState('config', {
            currentType: state => state.block.type,
            currentConfigBlockConfig: state => state.block.config
        })
    },
    methods: {
        onChange(typ, config) {
            try {
                const oldConfig = {...this.currentConfigBlockConfig}

                switch(typ){
                    case componentType.Grid:
                        oldConfig.common.zIndex = 1
                        break
                }

                this.mixinSetConfigTypeAndConfig(typ, {
                        common: oldConfig.common,
                        type: config
                })
            }catch(e) {
                console.log('config change json parse err', e, this.cfg)
            }
        },
        onAntVTypeClick() {
            if(this.currentType === componentType.AntV) return
            this.onChange(componentType.AntV, AntVConfig())
        },
        onTextTypeClick() {
            if(this.currentType === componentType.Text) return
            this.onChange(componentType.Text, TextConfig())
        },
        onTableTypeClick() {
            if(this.currentType === componentType.Table) return
            this.onChange(componentType.Table, TableConfig())
        },
        onGridTypeClick() {
            if(this.currentType === componentType.Grid) return
            this.onChange(componentType.Grid, GridConfig())
        }
    }
}
</script>