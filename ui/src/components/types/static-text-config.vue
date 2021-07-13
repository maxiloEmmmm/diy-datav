<script lang="tsx">
import configMixin from './config-mixin'
import {ViewBlockType, StaticTextConfig, StaticTextConfigParse} from 'type'
export default {
    mixins: [configMixin],
    data() {
        return {
            cfg: {
                ...ViewBlockType().config,
                type: StaticTextConfig()
            },
        }
    },
    render() {
        return <div>
            <ysz-list-item v-slots={{
                left: () => '内容'
            }}>
                <a-input size="small" vModel={[this.cfg.type.text, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '水平居中'
            }}>
                <a-switch size="small" vModel={[this.cfg.type.center.h, 'checked']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '垂直居中'
            }}>
                <a-switch size="small" vModel={[this.cfg.type.center.v, 'checked']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '加粗'
            }}>
                <a-switch size="small" vModel={[this.cfg.type.bold, 'checked']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '倾斜'
            }}>
                <a-switch size="small" vModel={[this.cfg.type.italic, 'checked']} onChange={this.onChange}/>
            </ysz-list-item>
        </div>
    },
    methods: {
        transformConfig() {
            try {
                const blockCfg = JSON.parse(this.config)
                blockCfg.type = StaticTextConfigParse(blockCfg.type)
                this.cfg = blockCfg
            }catch(e) {
                console.log('static text config parse failed', e, this.config)
            }
        },
    }
}
</script>