<script lang="jsx">
import {mapState} from "vuex";
import configComponent from '@/components/types/config.js'
import inputComponent from '@/components/input/config.js'
import {ViewBLockTypeCommon, ViewBLockTypeCommonInputItem} from 'type'
import typeConfigComponent from '@/components/types/type-config.vue'
import inputConfigComponent from '@/components/input/type-config.vue'
export default {
    components: {
        typeConfigComponent,
        inputConfigComponent
    },
    name: 'config-bar',
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
        }
    },
    methods: {
        transformConfig() {
            try {
                const config = JSON.parse(this.currentConfigBlockConfig)
                config.common.input = this.normalInput(config.common.input)
                this.cfg = config.common
            }catch (e) {
                console.log('config-bar parse config err', e)
            }
        },
        normalInput(inputs) {
            if(!Array.isArray(inputs)) {
                return []
            }

            return inputs
        },
        onChange() {
            try {
                const config = JSON.parse(this.currentConfigBlockConfig)
                this.mixinSetConfigConfig(JSON.stringify({
                    common: this.cfg,
                    type: config.type
                }))
            }catch(e) {
                console.log('config change json parse err', e, this.cfg)
            }
        }
    },
    render() {
        let cm = configComponent[this.currentConfigBlockType]
        return <a-tabs size="small">
            <a-tab-pane key="type" tab="类型">
                <type-config-component/>
                <a-divider/>
                {!!cm ? <cm config={this.currentConfigBlockConfig}/> : null}
            </a-tab-pane>
            <a-tab-pane key="common" tab="通用">
                <a-tabs class='common-config' size="small" tab-position="left">
                    <a-tab-pane key="type" tab="数据">
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
                                    this.onChange()
                                }} currentType={this.cfg.input[index].type}/>,
                                {cm ? [
                                    <a-divider/>,
                                    <ysz-list-item v-slots={{
                                        left: () => '标题'
                                    }}>
                                        <a-input size="small" vModel={[this.cfg.input[index].title, 'value']} onChange={this.onChange}/>
                                    </ysz-list-item>,
                                    <cm vModel={[this.cfg.input[index].config, 'config']} onChange={this.onChange}/>
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