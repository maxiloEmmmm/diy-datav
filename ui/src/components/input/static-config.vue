<script lang="jsx">
import {staticInputConfigParse, staticInputConfig, staticInputConfigDefault} from "type";
import configMixin from './config-mixin'
export default {
    mixins: [configMixin],
    render() {
        return <ysz-list-item v-slots={{
                left: () => '静态数据'
            }}>
            <a-select style="width:100%" options={this.refs} size="small" vModel={[this.cfg.id, 'value']} onChange={this.onChange}/>
        </ysz-list-item>
    },
    data() {
        return {
            cfg: staticInputConfig(),
            refs: []
        }
    },
    created() {
        this.$api[this.$apiType.StaticList]()
            .then(response => {
                this.refs = [
                    {label: '无', value: staticInputConfigDefault.id()},
                    ...response.data.map(ref => ({value: ref.id, label: ref.title}))
                ]
            })
    },
    methods: {
        transformConfig() {
            try {
                this.cfg = staticInputConfigParse(JSON.parse(this.config))
            }catch(e) {
                console.log('static config parse failed in static-config', e, this.config)
            }
        },
    }
}
</script>