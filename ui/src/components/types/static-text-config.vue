<script lang="tsx">
import configMixin from './config-mixin'
import {ViewBlockType, StaticText, StaticTextParse} from 'type'
export default {
    mixins: [configMixin],
    data() {
        return {
            cfg: {
                ...ViewBlockType().config,
                type: StaticText()
            },
        }
    },
    render() {
        return <ysz-list-item v-slots={{
            left: () => '内容'
        }}>
            <a-input size="small" vModel={[this.cfg.type.text, 'value']} onChange={this.onChange}/>
        </ysz-list-item>
    },
    methods: {
        transformConfig() {
            try {
                const blockCfg = JSON.parse(this.config)
                blockCfg.type = StaticTextParse(blockCfg.type)
                this.cfg = blockCfg
            }catch(e) {
                console.log('static text config parse failed', e, this.config)
            }
        },
    }
}
</script>