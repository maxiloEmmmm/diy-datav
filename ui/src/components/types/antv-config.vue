<script lang="tsx">
import {AntVConfig, AntVConfigParse} from 'type/types/index.js'
import configMixin from './config-mixin'
export default {
    mixins: [configMixin],
    data() {
        return {
            cfg: AntVConfig()
        }
    },
    render() {
        return <div>antv config</div>
    },
    methods: {
        transformConfig() {
            try {
                this.cfg = AntVConfigParse(JSON.parse(this.config))
            }catch(e) {
                console.log('AntV config parse failed in antv-config', e, this.config)
            }
        },
    }
}
</script>