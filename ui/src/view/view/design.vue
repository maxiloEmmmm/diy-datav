<script lang="jsx">
import { ViewBlockType } from 'type'
export default {
    render() {
        let blocks = this.blocks.map(block => {
            return <block-wrap class='view-block'>{block.refresh}</block-wrap>
        })
        let bg = <div id='diy-data-view_bg'></div>
        let util = <div id='diy-data-view_util'></div>

        return <div id='diy-datav-view'>
            {bg}
            {util}
            {blocks}
        </div>
    },
    data() {
        return {
            blocks: [ViewBlockType()]
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
        position: absolute; right: 0; left: 0; z-index: 2;
    }

    .view-block {
        position: absolute; z-index: 3;
    }
}
</style>