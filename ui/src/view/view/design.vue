<script lang="jsx">
import { ViewType, ViewBlockType } from 'type'
import bgAssetsDev from '@/assets/bg_design.png'
import {mapState} from 'vuex'
import configComponent from '@/components/types/config.js'
export default {
    render() {
        let blocks = this.view.blocks.map(block => {
            let blockKey = block.getKey()
            return <block-wrap class="diy-data-view_block" key={blockKey} block-key={blockKey}>{block.config.common.refresh}</block-wrap>
        })
        let bg = <div id='diy-data-view_bg' style={this._bg_style}></div>
        let util = <div id='diy-data-view_util'>
            <a-button onClick={this.newBlock}>添加块</a-button>
            current drag id: { this.dragBlockID }, focus: {this.app_mixin.focus.in ? 'focus' : 'no-focus'}
        </div>

        let cm = configComponent[this.currentConfigBlockType]
        let configBar = <a-drawer visible={this.configShow && !!cm} onClick={this.onConfigBarClose}>
            <cm config={this.currentConfigBlockConfig}/>
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
        this.$api[this.$apiType.ViewStore]({id: 1})
            .then(response => {
                console.log(response.data.body)
            })

        this.mixinInitFocus()
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
            currentConfigBlockConfig: state => state.block.config
        })
    },
    methods: {
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
}
</style>