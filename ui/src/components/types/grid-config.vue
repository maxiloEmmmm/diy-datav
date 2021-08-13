<script lang="tsx">
import configMixin from '../config-mixin'
import {
    ViewBlockType, GridConfig, GridConfigParse, GridConfigDefault
} from 'type'

import inputChoose from './input-choose.vue'
export default {
    components: {inputChoose},
    mixins: [configMixin],
    data() {
        return {
            cfg: {
                ...ViewBlockType().config,
                type: GridConfig()
            },
        }
    },
    render() {
        return <a-tabs
            type="editable-card"
            onEdit={this.onEdit}
            tabPosition="left"
        >
            {this.cfg.type.rows.map((row, index) => <a-tab-pane key={index} tab={`第${index + 1}行`} closable={true}>
                <ysz-list-item-top
                    v-slots={{
                        top: () => '高百分比'
                    }}
                >
                    <a-input-number size="small" vModel={[this.cfg.type.rows[index].height, 'value']} onChange={this.onChange}/>
                </ysz-list-item-top>
                <ysz-list-item-top
                    v-slots={{
                        top: () => '行配置'
                    }}
                >
                    <more onAdd={payload => {
                        this.cfg.type.rows[index].rowCols[payload.count] = GridConfigDefault.rowCol()
                        this.onChange()
                        payload.done()
                    }} onRemove={payload => {
                        this.cfg.type.rows[index].rowCols = this.cfg.rows[index].rowCols.filter((v, i) => i !== payload.index)
                        this.onChange()
                        payload.done()
                    }} component={rowColIndex => {
                        return <ysz-list-item-top
                            v-slots={{
                                top: () => '宽百分比'
                            }}
                        >
                            <a-input-number size="small" vModel={[this.cfg.type.rows[index].rowCols[rowColIndex].width, 'value']} onChange={this.onChange}/>
                        </ysz-list-item-top>
                    }}/>
                </ysz-list-item-top>
            </a-tab-pane>)}
        </a-tabs>
    },
    methods: {
        onEdit(targetKey: string | MouseEvent, action: string) {
            if (action === 'add') {
                this.cfg.type.rows.push(GridConfigDefault.row())
            } else {
                this.cfg.type.rows = this.cfg.type.rows.filter((row, index) => index !== targetKey)
            }
            this.onChange()
        },
        transformConfig() {
            try {
                const blockCfg = JSON.parse(this.config)
                blockCfg.type = GridConfigParse(blockCfg.type)
                this.cfg = blockCfg
            }catch(e) {
                console.log('grid config parse failed', e, this.config)
            }
        },
    }
}
</script>