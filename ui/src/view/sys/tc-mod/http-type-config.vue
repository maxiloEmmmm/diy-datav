<script lang="jsx">
import typeConfigMixin from './type-config-mixin'
import {httpTypeConfigParse, httpTypeConfig} from "./types";
export default {
    mixins: [typeConfigMixin],
    data() {
        return {
            cfg: httpTypeConfig()
        }
    },
    mounted() {
        this.$nextTick(() => {
            this.$refs.form.setFields([
                {title: '接口地址', field: 'url'}
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
                this.cfg = httpTypeConfigParse(JSON.parse(this.config))
                if(this.$refs.form) {
                    this.$refs.form.setData(this.cfg)
                }
            }catch(e) {
                console.log('http config parse failed in http-config', e, this.config)
            }
        },
    }
}
</script>