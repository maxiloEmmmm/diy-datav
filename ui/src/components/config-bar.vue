<script lang="jsx">
import {mapState} from "vuex";
import configComponent from '@/components/types/config.js'
import inputComponent from '@/components/input/config.js'
import * as configComponentType from '@/components/types/type.js'
import {borderStyle, descPosition, ViewBLockTypeCommon, ViewBLockTypeCommonInputItem, ViewBLockTypeCommonInputItemDefault, ViewBLockTypeCommonParse} from 'type'
import typeConfigComponent from '@/components/types/type-config.vue'
import inputConfigComponent from '@/components/input/type-config.vue'
export default {
    components: {
        typeConfigComponent,
        inputConfigComponent
    },
    name: 'config-bar',
    props: {
        maxZIndex: {
            type: Number,
            default: 1
        }
    },
    computed: {
        ...mapState('config', {
            currentConfigBlockType: state => state.block.type,
            currentConfigBlockConfig: state => state.block.config,
        })
    },
    watch: {
        currentConfigBlockConfig: {
            immediate: true,
            handler: 'transformConfig'
        }
    },
    data() {
        return {
            cfg: ViewBLockTypeCommon(),
            store: {
                descTypes: descPosition.map(t => ({label: t, value: t})),
                borderStyles: borderStyle.map(t => ({label: t, value: t})),
            }
        }
    },
    methods: {
        transformConfig() {
            this.cfg = this.$util.merge({}, this.currentConfigBlockConfig.common)
        },
        normalInput(inputs) {
            if(!Array.isArray(inputs)) {
                return []
            }

            return inputs
        },
        // data is dirty, so can't use old id to fetch data on block in design model
        hookDataToTmpEchoModel(index) {
            this.cfg.input[index].id = ViewBLockTypeCommonInputItemDefault.id()
        },
        onChange() {
            this.mixinSetConfigConfig(this.$util.merge(this.currentConfigBlockConfig, {common: this.cfg}))
        },
        onBottomZIndex() {
            this.cfg.zIndex = 1
            this.onChange()
        },
        onTopZIndex() {
            this.cfg.zIndex = this.maxZIndex + 1
            this.onChange()
        }
    },
    render() {
        let cm = configComponent[this.currentConfigBlockType]
        let textCM = configComponent[configComponentType.Text]
        return <a-tabs size="small">
            <a-tab-pane key="type" tab="类型">
                <type-config-component/>
                <a-divider/>
                {!!cm ? <cm config={this.currentConfigBlockConfig}/> : null}
            </a-tab-pane>
            <a-tab-pane key="common" tab="通用">
                <a-tabs class='common-config' size="small" tab-position="left">
                    <a-tab-pane key="bg" tab="背景">
                        <ysz-list-item-top v-slots={{
                            top: () => '颜色'
                        }}>
                            <color-pick size="small" vModel={[this.cfg.bg, 'value']} onChange={this.onChange}/>
                        </ysz-list-item-top>
                    </a-tab-pane>
                    <a-tab-pane key="desc" tab="描述">
                        <textCM config={this.cfg.desc.textConfig} onChange={this.onChange} objectValue handleChange/>
                        <ysz-list-item-top v-slots={{
                            top: () => '高度占比'
                        }}>
                            <a-input-number size="small" vModel={[this.cfg.desc.position.height, 'value']} onChange={this.onChange}/>
                        </ysz-list-item-top>
                        <ysz-list-item-top v-slots={{
                            top: () => '位置'
                        }}>
                            <a-select size="small" options={this.store.descTypes} vModel={[this.cfg.desc.positionType, 'value']} onChange={this.onChange}/>
                        </ysz-list-item-top>
                    </a-tab-pane>
                    <a-tab-pane key="border" tab="边框">
                        <ysz-list-item-top v-slots={{
                            top: () => '粗细'
                        }}>
                            <a-input-number size="small" vModel={[this.cfg.border.width, 'value']} onChange={this.onChange}/>
                        </ysz-list-item-top>
                        <ysz-list-item-top v-slots={{
                            top: () => '颜色'
                        }}>
                            <color-pick size="small" vModel={[this.cfg.border.color, 'value']} onChange={this.onChange}/>
                        </ysz-list-item-top>
                        <ysz-list-item-top v-slots={{
                            top: () => '风格'
                        }}>
                            <a-select size="small" options={this.store.borderStyles} vModel={[this.cfg.border.style, 'value']} onChange={this.onChange}/>
                        </ysz-list-item-top>
                    </a-tab-pane>
                    <a-tab-pane key="position" tab="布局">
                        <ysz-list-item-top v-slots={{
                            top: () => '层叠位置'
                        }}>
                            {this.currentConfigBlockType === configComponentType.Grid ? this.cfg.zIndex
                                : <ysz-list row group={3}>
                                    <a-input-number size="small" vModel={[this.cfg.zIndex, 'value']} onChange={this.onChange}/>
                                    <a-button size="small" onClick={this.onBottomZIndex}>最下面</a-button>
                                    <a-button size="small" onClick={this.onTopZIndex}>最上面</a-button>
                                </ysz-list>}
                        </ysz-list-item-top>
                    </a-tab-pane>
                    <a-tab-pane key="type" tab="数据">
                        <ysz-list-item v-slots={{
                            left: () => '刷新时间'
                        }}>
                            <a-input-number formatter={v => `${v}秒`} size="small" vModel={[this.cfg.refresh, 'value']} onChange={this.onChange}/>
                        </ysz-list-item>
                        <more initCount={this.cfg.input.length} onAdd={payload => {
                            this.cfg.input[payload.count] = ViewBLockTypeCommonInputItem()
                            this.onChange()
                            payload.done()
                        }} onRemove={payload => {
                            this.cfg.input = this.cfg.input.filter((v, i) => i !== payload.index)
                            this.onChange()
                            payload.done()
                        }} component={index => {
                            const cm = inputComponent[this.cfg.input[index].type]
                            return <div>
                                <inputConfigComponent onChange={payload => {
                                    this.cfg.input[index].type = payload.type
                                    this.cfg.input[index].config = payload.config
                                    this.hookDataToTmpEchoModel(index)
                                    this.onChange()
                                }} currentType={this.cfg.input[index].type}/>,
                                {cm ? [
                                    <a-divider/>,
                                    <ysz-list-item v-slots={{
                                        left: () => '标题'
                                    }}>
                                        <a-input size="small" vModel={[this.cfg.input[index].title, 'value']} onChange={this.onChange}/>
                                    </ysz-list-item>,
                                    <cm vModel={[this.cfg.input[index].config, 'config']} onChange={() => {
                                        this.hookDataToTmpEchoModel(index)
                                        this.onChange()
                                    }}/>
                                ] : null}
                            </div>
                        }}/>
                    </a-tab-pane>
                </a-tabs>
            </a-tab-pane>
        </a-tabs>
    },

}
</script>

<style lang="scss" scoped>
.common-config {
    &:deep .ant-tabs-tab {
        padding: 0; margin: 0; padding-right: 1rem; margin-bottom: 1rem
    }
    &:deep .ant-tabs-left-content ant-tabs-content {
        padding-left: 0
    }
}
</style>