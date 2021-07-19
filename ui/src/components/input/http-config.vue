<script lang="jsx">
import {httpInputConfigParse, httpInputConfig} from "type";
import configMixin from './config-mixin'
export default {
    mixins: [configMixin],
    render() {
        return <div>
            {this.cfg.ref > 0 ? <ysz-list-item v-slots={{
                left: () => '地址'
            }}>
                <a-input size="small" vModel={[this.cfg.url, 'value']} onChange={this.onChange}/>
            </ysz-list-item> : null}
            <ysz-list-item v-slots={{
                left: () => '已有接口'
            }}>
                <a-select options={this.refs} size="small" vModel={[this.cfg.ref, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
        </div>
    },
    data() {
        return {
            cfg: httpInputConfig(),
            refs: []
        }
    },
    created() {
        this.$api[this.$apiType.HttpList]()
            .then(response => {
                this.refs = response.data.map(ref => ({value: ref.id, label: ref.title}))
            })
    },
    methods: {
        transformConfig() {
            try {
                this.cfg = httpInputConfigParse(JSON.parse(this.config))
            }catch(e) {
                console.log('http config parse failed in http-config', e, this.config)
            }
        },
    }
}
</script>