<script lang="jsx">
import { ViewType, ViewBlockType } from 'type'
import bgAssetsDev from '@/assets/bg_design.png'
import {mapState} from 'vuex'
import configComponent from '@/components/types/config.js'
import typeConfigComponent from '@/components/types/type-config.vue'
import {Module as HelpModule} from '@/mixins/help'
import { BarChartOutlined } from '@ant-design/icons-vue';
import * as type from "../../api/type";
export default {
    components: {
        typeConfigComponent
    },
    render() {
        let blocks = this.view.blocks.map(block => {
            let blockKey = block.getKey()
            return <block-wrap class="diy-data-view_block" key={blockKey} block-key={blockKey}>
                <view-block type={block.type} config={block.config} />
            </block-wrap>
        })
        let bg = <div id='diy-data-view_bg' style={this._bg_style} />
        let util = <div id='diy-data-view_util'>
            <a-button onClick={this.newBlock}>添加块</a-button>
            current drag id: { this.dragBlockID }, focus: {this.app_mixin.focus.in ? 'focus' : 'no-focus'}
        </div>

        let cm = configComponent[this.currentConfigBlockType]

        // TODO: 考虑要不要加个预览在配置旁边
        // let configView = this.configShow ? <div id='config-view'>
        //     <view-block type={this.currentConfigBlockType} config={this.currentConfigBlockConfig} />
        // </div> : null

        let configBar = <a-drawer mask={false} width="40vw" visible={this.configShow} onClose={this.onConfigBarClose}>
            <type-config-component />
            <a-divider />
            {!!cm ? <cm config={this.currentConfigBlockConfig} onChange={this.onConfigChange}/> : null}
        </a-drawer>

        return <div id='diy-datav-view'
            onDrop={this.onDrop}
            onMousedown={this.onMouseDown}
        >
            {bg}
            {util}
            {blocks}

            {configBar}
        </div>
    },
    data() {
        return {
            view: ViewType()
        }
    },
    created() {
        this.mixinInitFocus()

        this.mixinAddHelp(HelpModule.ViewBlock, [
            {key: "edit", component() {
                    return <BarChartOutlined twoToneColor="red"/>
                }, cb: (payload) => {
                    if(!payload.blockKey) {
                        return
                    }

                    const block = this.view.blocks.filter(block => block.getKey() === payload.blockKey)[0]
                    if (!!block) {
                        this.mixinSetConfigKey(payload.blockKey)
                        // TODO: how to get config edit action to edit block config
                        this.mixinSetConfigTypeAndConfig(block.type, block.config)
                        this.mixinConfigShow()
                    }
                },
            }
        ])

        this.fetch()
    },
    computed: {
        _bg_style() {
            return {
                backgroundImage: `url(${bgAssetsDev})`,
                backgroundSize: '100% 100%'
            }
        },
        dragBlockID() {
            return this.$store.state.view.dragBlockId
        },
        ...mapState('config', {
            configShow: 'show',
            currentConfigBlockType: state => state.block.type,
            currentConfigBlockConfig: state => state.block.config,
            currentConfigBlockKey: state => state.block.key
        })
    },
    methods: {
        fetch() {
            this.$api[this.$apiType.ViewInfo]("1")
                .then(response => {
                    this.view = {
                        ...ViewType(),
                        ...response.data.data,
                        blocks: response.data.data.blocks.map(block => {
                            return {
                                ...ViewBlockType(),
                                ...block,
                            }
                        })
                    }

                    this.$store.dispatch('view/fetchData')
                })
        },
        newBlock() {
            this.view.newBlockAndStore()
        },
        onDrop() {

        },
        onMouseDown() {
            this.mixinDoFocus()
        },
        onConfigBarClose() {
            this.mixinConfigHidden()
            this.mixinSetConfigTypeAndConfig()
        },
        onConfigChange(data) {
            const block = this.view.blocks.filter(block => block.getKey() === this.currentConfigBlockKey)[0]
            if (!!block) {
                block.config = data
            }
        }
    }
}
</script>

<style lang="scss" scoped>
#diy-datav-view {
    width: 100vw; height: 100vh; position: relative;

    #diy-data-view_bg {
        position: absolute; top: 0; right: 0; left: 0; bottom: 0; z-index: 1;
    }

    #diy-data-view_util {
        position: absolute; right: 0; top: 0; z-index: 2;
    }

    .diy-data-view_block {
        position: absolute; z-index: 3;
    }

    #config-view {
        position: absolute; z-index: 4; width:50vw; height: 60vh; top: 10vh; left: 5vw
    }
}
</style>