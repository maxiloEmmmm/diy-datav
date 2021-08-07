<script lang="tsx">
import configMixin from '../config-mixin'
import {
    ViewBlockType, TableConfig, TableConfigParse,
} from 'type'

import inputChoose from './input-choose.vue'
export default {
    components: {inputChoose},
    mixins: [configMixin],
    data() {
        return {
            cfg: {
                ...ViewBlockType().config,
                type: TableConfig()
            },
        }
    },
    render() {
        return <div>
            <ysz-list-item v-slots={{
                left: () => '数据'
            }}>
                <inputChoose inputs={this.cfg.common.input}
                             value={this.cfg.type.dataIndex}
                             onChange={value => {
                                 this.cfg.type.dataIndex = value
                                 this.onChange()
                             }} />
            </ysz-list-item>
        </div>
    },
    methods: {
        transformConfig() {
            try {
                const blockCfg = JSON.parse(this.config)
                blockCfg.type = TableConfigParse(blockCfg.type)
                this.cfg = blockCfg
            }catch(e) {
                console.log('static table config parse failed', e, this.config)
            }
        },
    }
}
</script>