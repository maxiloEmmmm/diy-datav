<script lang="jsx">
import {httpInputConfigParse, httpInputConfig} from "type";
import configMixin from './config-mixin'
export default {
    mixins: [configMixin],
    render() {
        return <ysz-list-item v-slots={{
            left: () => '地址'
        }}>
            <a-input size="small" vModel={[this.cfg.url, 'value']} onChange={this.onChange}/>
        </ysz-list-item>
    },
    data() {
        return {
            cfg: httpInputConfig()
        }
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