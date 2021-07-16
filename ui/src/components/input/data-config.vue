<script lang="tsx">
import { BarChartOutlined, ReadOutlined } from '@ant-design/icons-vue';
import * as dataType from './type.js'
import SelectItem from "../select-item.vue";

export default {
    components: {SelectItem},
    props: {
        currentType: {
            default: dataType.Http,
            type: String
        }
    },
    emits: ['change'],
    render() {
        // TODO: tranform to config list
        return <a-row gutter={[16,16]}>
            <a-col span={4}><select-item active={this.currentType === dataType.Http} onClick={this.onHttpTypeClick} icon={() => <BarChartOutlined style="font-size:1.4rem"/>}>Http</select-item></a-col>
            <a-col span={4}><select-item active={this.currentType === dataType.Sql} onClick={this.onSqlTypeClick} icon={() => <ReadOutlined style="font-size:1.4rem"/>}>数据库</select-item></a-col>
        </a-row>
    },
    methods: {
        onHttpTypeClick() {
            this.$emit('change', dataType.Http)
        },
        onSqlTypeClick() {
            this.$emit('change', dataType.Sql)
        },
    }
}
</script>