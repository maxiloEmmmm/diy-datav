<script lang="jsx">
import typeConfigMixin from './type-config-mixin'
import {staticTypeConfig, staticTypeConfigParse} from "./types";
import util from 'pkg/util'
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
                }, onChange(val, item) {
                    try {
                        // 去除换行符等不必要的字符
                        let vk = util.uuid("staticTc").replaceAll("-", "")
                        window.eval(`window.${vk}=${val};`)
                        item.data = JSON.stringify(window[vk])
                    }catch(e) {
                    // 不处理
                    }
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