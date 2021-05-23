<script lang="jsx">
import { ViewType, ViewBlockType } from 'type'
export default {
    render() {
        let blocks = this.view.blocks.map(block => {
            return <block-wrap class='view-block'>{block.config.common.refresh}</block-wrap>
        })
        let bg = <div id='diy-data-view_bg'></div>
        let util = <div id='diy-data-view_util'>
            <a-button onClick={this.newBlock}>添加块</a-button>
        </div>

        return <div id='diy-datav-view'>
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
    methods: {
        newBlock() {
            this.view.newBlockAndStore()
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

    .view-block {
        position: absolute; z-index: 3;
    }
}
</style>