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
            quick: {
                row: 1, col: 1
            }
        }
    },
    render() {
        return <ysz-list>
            <ysz-list-item-top
                v-slots={{
                    top: () => '间距'
                }}
            >
                <ysz-list row group={2}>
                    <ysz-list-item-top
                        v-slots={{
                            top: () => 'top'
                        }}
                    >
                        <a-input-number size="small" vModel={[this.cfg.type.padding.top, 'value']} onChange={this.onChange}/>
                    </ysz-list-item-top>
                    <ysz-list-item-top
                        v-slots={{
                            top: () => 'bottom'
                        }}
                    >
                        <a-input-number size="small" vModel={[this.cfg.type.padding.bottom, 'value']} onChange={this.onChange}/>
                    </ysz-list-item-top>
                    <ysz-list-item-top
                        v-slots={{
                            top: () => 'left'
                        }}
                    >
                        <a-input-number size="small" vModel={[this.cfg.type.padding.left, 'value']} onChange={this.onChange}/>
                    </ysz-list-item-top>
                    <ysz-list-item-top
                        v-slots={{
                            top: () => 'right'
                        }}
                    >
                        <a-input-number size="small" vModel={[this.cfg.type.padding.right, 'value']} onChange={this.onChange}/>
                    </ysz-list-item-top>
                </ysz-list>
            </ysz-list-item-top>
            <ysz-list-item
                v-slots={{
                    left: () => '快速分割',
                }}
            >
                <a-input size="small" vModel={[this.quick.row, 'value']}/>
                <a-input size="small" vModel={[this.quick.col, 'value']}/>
                <a-button size="small" onClick={() => {
                    if(this.cfg.type.rows.length >= this.quick.row) {
                        this.cfg.type.rows = this.cfg.type.rows.splice(0, this.quick.row)
                    }
                    for (let i = 0; i < this.quick.row; i++) {
                        let row
                        if(this.cfg.type.rows[i] === undefined) {
                            row = GridConfigDefault.row()
                            this.cfg.type.rows.push(row)
                        }else {
                            row = this.cfg.type.rows[i]
                        }
                        if(row.rowCols.length >= this.quick.col) {
                            row.rowCols = row.rowCols.splice(0, this.quick.col)
                        }
                        for (let c = 0; c < this.quick.col; c++) {
                            if(row.rowCols[c] === undefined) {
                                row.rowCols[c] = GridConfigDefault.rowCol()
                            }
                            this.colWidthAvg(i)
                        }
                    }
                    this.rowHeightAvg()
                }}>分割</a-button>
            </ysz-list-item>
            <ysz-list-item
                v-slots={{
                    left: () => '工具'
                }}
            >
                <a-button size="small" onClick={this.rowHeightAvg}>行等高</a-button>
            </ysz-list-item>
            <ysz-list-item-top
                v-slots={{
                    top: () => '布局'
                }}
            >
                <a-tabs
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
                        <ysz-list row group={2}>
                            <ysz-list-item-top
                                v-slots={{
                                    top: () => '内间距 - top'
                                }}
                            >
                                <a-input-number size="small" vModel={[this.cfg.type.rows[index].padding.top, 'value']} onChange={this.onChange}/>
                            </ysz-list-item-top>
                            <ysz-list-item-top
                                v-slots={{
                                    top: () => '内间距 - bottom'
                                }}
                            >
                                <a-input-number size="small" vModel={[this.cfg.type.rows[index].padding.bottom, 'value']} onChange={this.onChange}/>
                            </ysz-list-item-top>
                            <ysz-list-item-top
                                v-slots={{
                                    top: () => '内间距 - left'
                                }}
                            >
                                <a-input-number size="small" vModel={[this.cfg.type.rows[index].padding.left, 'value']} onChange={this.onChange}/>
                            </ysz-list-item-top>
                            <ysz-list-item-top
                                v-slots={{
                                    top: () => '内间距 - right'
                                }}
                            >
                                <a-input-number size="small" vModel={[this.cfg.type.rows[index].padding.right, 'value']} onChange={this.onChange}/>
                            </ysz-list-item-top>
                        </ysz-list>
                        <ysz-list-item
                            v-slots={{
                                left: () => '工具'
                            }}
                        >
                            <a-button size="small" onClick={() => this.colWidthAvg(index)}>列等宽</a-button>
                        </ysz-list-item>
                        <ysz-list-item-top
                            v-slots={{
                                top: () => '列配置'
                            }}
                        >
                            <more initCount={this.cfg.type.rows[index].rowCols.length} onAdd={payload => {
                                this.cfg.type.rows[index].rowCols[payload.count] = GridConfigDefault.rowCol()
                                this.onChange()
                                payload.done()
                            }} onRemove={payload => {
                                this.cfg.type.rows[index].rowCols = this.cfg.type.rows[index].rowCols.filter((v, i) => i !== payload.index)
                                this.onChange()
                                payload.done()
                            }} component={rowColIndex => {
                                return <ysz-list row group={2}>
                                    <ysz-list-item-top
                                        v-slots={{
                                            top: () => '内间距 - top'
                                        }}
                                    >
                                        <a-input-number size="small" vModel={[this.cfg.type.rows[index].rowCols[rowColIndex].padding.top, 'value']} onChange={this.onChange}/>
                                    </ysz-list-item-top>
                                    <ysz-list-item-top
                                        v-slots={{
                                            top: () => '内间距 - bottom'
                                        }}
                                    >
                                        <a-input-number size="small" vModel={[this.cfg.type.rows[index].rowCols[rowColIndex].padding.bottom, 'value']} onChange={this.onChange}/>
                                    </ysz-list-item-top>
                                    <ysz-list-item-top
                                        v-slots={{
                                            top: () => '内间距 - left'
                                        }}
                                    >
                                        <a-input-number size="small" vModel={[this.cfg.type.rows[index].rowCols[rowColIndex].padding.left, 'value']} onChange={this.onChange}/>
                                    </ysz-list-item-top>
                                    <ysz-list-item-top
                                        v-slots={{
                                            top: () => '内间距 - right'
                                        }}
                                    >
                                        <a-input-number size="small" vModel={[this.cfg.type.rows[index].rowCols[rowColIndex].padding.right, 'value']} onChange={this.onChange}/>
                                    </ysz-list-item-top>
                                    <ysz-list-item-top
                                        v-slots={{
                                            top: () => '宽百分比'
                                        }}
                                    >
                                        <a-input-number size="small" vModel={[this.cfg.type.rows[index].rowCols[rowColIndex].width, 'value']} onChange={this.onChange}/>
                                    </ysz-list-item-top>
                                </ysz-list>
                            }}/>
                        </ysz-list-item-top>
                    </a-tab-pane>)}
                </a-tabs>
            </ysz-list-item-top>
        </ysz-list>
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
        rowHeightAvg() {
            const len = this.cfg.type.rows.length

            if(len === 0) {
                return
            }
            const h = parseFloat((100 / len).toFixed(6))
            this.cfg.type.rows.forEach((row, rowIndex) => {
                this.cfg.type.rows[rowIndex].height = h
            })
            this.onChange()
        },
        colWidthAvg(index) {
            const len = this.cfg.type.rows[index].rowCols.length

            if(len === 0) {
                return
            }
            const w = parseFloat((100 / len).toFixed(6))
            this.cfg.type.rows[index].rowCols.forEach((col, colIndex) => {
                this.cfg.type.rows[index].rowCols[colIndex].width = w
            })
            this.onChange()
        }
    }
}
</script>