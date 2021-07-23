<script lang="jsx">
import typeConfigMixin from './type-config-mixin'
import {sqlTypeConfig, sqlTypeConfigParse, sqlTypeOption} from "./types";
export default {
    mixins: [typeConfigMixin],
    data() {
        return {
            cfg: sqlTypeConfig()
        }
    },
    mounted() {
        this.$nextTick(() => {
            this.$refs.form.setFields([
                {title: '类型', field: 'type', type: 'select', option: {
                    selectOptions: sqlTypeOption.map(t => ({label: t, value: t}))
                }},
                {title: '地址', field: 'host'},
                {title: '端口', field: 'port', type: 'number'},
                {title: '用户', field: 'user'},
                {title: '密码', field: 'pass'},
                {title: '数据库', field: 'db'},
            ])
            this.$refs.form.setModels([
                {key: 'default'}
            ])
            this.$refs.form.setData(this.cfg)
            this.$refs.form.setModel('default')
        })
    },
    methods: {
        transformConfig() {
            try {
                this.cfg = sqlTypeConfigParse(JSON.parse(this.config))
                if(this.$refs.form) {
                    this.$refs.form.setData(this.cfg)
                }
            }catch(e) {
                console.log('sql config parse failed in sql-config', e, this.config)
            }
        },
    }
}
</script>