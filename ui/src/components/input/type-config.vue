<script lang="jsx">
import { CloudOutlined, ConsoleSqlOutlined} from '@ant-design/icons-vue';
import * as dataType from './type.js'
import SelectItem from "../select-item.vue";
import {sqlInputConfig, httpInputConfig} from 'type'

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
            <a-col span={8}><select-item active={this.currentType === dataType.Http} onClick={this.onHttpTypeClick} icon={() => <CloudOutlined style="font-size:1.4rem"/>}>Http</select-item></a-col>
            <a-col span={8}><select-item active={this.currentType === dataType.Sql} onClick={this.onSqlTypeClick} icon={() => <ConsoleSqlOutlined style="font-size:1.4rem"/>}>数据库</select-item></a-col>
        </a-row>
    },
    methods: {
        onHttpTypeClick() {
            if(this.currentType !== dataType.Http) {
                this.$emit('change', {
                    type: dataType.Http,
                    config: JSON.stringify(httpInputConfig()),
                })
            }
        },
        onSqlTypeClick() {
            if(this.currentType !== dataType.Sql) {
                this.$emit('change', {
                    type: dataType.Sql,
                    config: JSON.stringify(sqlInputConfig()),
                })
            }
        },
    }
}
</script>