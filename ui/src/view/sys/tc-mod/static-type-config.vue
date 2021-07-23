<script lang="jsx">
import typeConfigMixin from './type-config-mixin'
import {staticTypeConfig, staticTypeConfigParse} from "./types";
export default {
    mixins: [typeConfigMixin],
    data() {
        return {
            cfg: staticTypeConfig()
        }
    },
    mounted() {
        this.$nextTick(() => {
            this.$refs.form.setFields([
                {title: '数据', field: 'data', type: 'code', option: {
                    language: 'json'
                }},
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
                this.cfg = staticTypeConfigParse(JSON.parse(this.config))
                if(this.$refs.form) {
                    this.$refs.form.setData(this.cfg)
                }
            }catch(e) {
                console.log('static config parse failed in static-config', e, this.config)
            }
        },
    }
}
</script>