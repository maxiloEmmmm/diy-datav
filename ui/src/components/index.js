import 'ant-design-vue/dist/antd.css'
import antd from 'ant-design-vue'
import blockWrap from './block-wrap.vue'
import block from './block.vue'
import selectItem from './select-item.vue'
import listItem from './list-item.vue'
import listItemTop from './list-item-top.vue'
import more from './more.vue'
import colorPick from './color-pick.vue'
export default {
    install(app) {
        app.use(antd)
        app.component(listItem.name, listItem)
        app.component(listItemTop.name, listItemTop)
        app.component(blockWrap.name, blockWrap)
        app.component(block.name, block)
        app.component(selectItem.name, selectItem)
        app.component(more.name, more)
        app.component(colorPick.name, colorPick)
    }
}
