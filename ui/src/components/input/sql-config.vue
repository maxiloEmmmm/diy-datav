<script lang="jsx">
import {sqlInputConfigParse, sqlInputConfig, sqlInputConfigDefault} from "type";
import configMixin from './config-mixin'
export default {
    mixins: [configMixin],
    render() {
        return <div>
            <ysz-list-item v-slots={{
                left: () => '查询语句'
            }}>
                <a-input size="small" vModel={[this.cfg.sql, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '数据库'
            }}>
                <a-select style="width:100%" options={this.engines} size="small" vModel={[this.cfg.engine, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
        </div>
    },
    data() {
        return {
            cfg: sqlInputConfig(),
            engines: []
        }
    },
    created() {
        this.$api[this.$apiType.SqlList]()
            .then(response => {
                this.engines = [
                    {label: '请选择', value: sqlInputConfigDefault.engine()},
                    ...response.data.map(ref => ({value: ref.id, label: ref.title}))
                ]
            })
    },
    methods: {
        transformConfig() {
            try {
                this.cfg = sqlInputConfigParse(JSON.parse(this.config))
            }catch(e) {
                console.log('sql config parse failed in sql-config', e, this.config)
            }
        },
    }
}
</script>