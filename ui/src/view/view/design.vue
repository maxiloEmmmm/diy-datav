<script lang="jsx">
import { ViewType, ViewBlockType } from 'type'
import bgAssetsDev from '@/assets/bg_design.png'
import {mapState} from 'vuex'
import {Module as HelpModule} from '@/mixins/help'
import { BarChartOutlined } from '@ant-design/icons-vue';
export default {
    render() {
        let blocks = this.view.blocks.map(block => {
            let blockKey = block.getKey()
            return <block-wrap class="diy-data-view_block" key={blockKey} block-key={blockKey} onMousedown={e => this.onBlockMouseDown(blockKey)}>
                <view-block type={block.type} config={block.config} />
            </block-wrap>
        })
        let bg = <div id='diy-data-view_bg' style={this._bg_style} />
        let util = <div id='diy-data-view_util'>
            <a-button onClick={this.newBlock}>添加块</a-button>
            current drag id: { this.dragBlockID }, focus: {this.app_mixin.focus.in ? 'focus' : 'no-focus'}
        </div>

        // TODO: 考虑要不要加个预览在配置旁边
        // let configView = this.configShow ? <div id='config-view'>
        //     <view-block type={this.currentConfigBlockType} config={this.currentConfigBlockConfig} />
        // </div> : null

        let configBar = <a-drawer mask={false} width="40vw" visible={this.configShow} onClose={this.onConfigBarClose}>
            <config-bar></config-bar>
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
                    this.blockConfigReplace(payload.blockKey)
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
    watch: {
        currentConfigBlockType: 'onTypeChange',
        currentConfigBlockConfig: 'onConfigChange'
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
            this.mixinClearConfig()
        },
        onTypeChange(typ) {
            const block = this.view.blocks.filter(block => block.getKey() === this.currentConfigBlockKey)[0]
            if (!!block) {
                block.type = typ
            }
        },
        onConfigChange(cfg) {
            const block = this.view.blocks.filter(block => block.getKey() === this.currentConfigBlockKey)[0]
            if (!!block) {
                block.config = cfg
            }
        },
        onBlockMouseDown(blockKey) {
            if(this.configShow) {
                this.blockConfigReplace(blockKey)
            }
        },
        blockConfigReplace(blockKey) {
            const block = this.view.blocks.filter(block => block.getKey() === blockKey)[0]
            if (!!block) {
                this.mixinSetConfigKey(blockKey)
                // TODO: how to get config edit action to edit block config
                this.mixinSetConfigTypeAndConfig(block.type, block.config)
                this.mixinConfigShow()
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