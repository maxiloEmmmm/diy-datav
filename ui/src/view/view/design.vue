<script lang="jsx">
import { ViewType, ViewBlockType } from 'type'
import bgAssetsDev from '@/assets/bg_design.png'
export default {
    render() {
        let blocks = this.view.blocks.map((block, blockIndex) => {
            return <block-wrap block-key={blockIndex}>{block.config.common.refresh}</block-wrap>
        })
        let bg = <div id='diy-data-view_bg' style={this._bg_style}></div>
        let util = <div id='diy-data-view_util'>
            <a-button onClick={this.newBlock}>添加块</a-button>
            current drag id: { this.dragBlockID }
        </div>

        return <div id='diy-datav-view'
            onDrop={this.onDrop}
        >
            {bg}
            {util}
            {blocks}
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
        }
    },
    methods: {
        newBlock() {
            this.view.newBlockAndStore()
        },
        onDrop() {

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
}
</style>