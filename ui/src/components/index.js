import 'ant-design-vue/dist/antd.css'
import antd from 'ant-design-vue'
import blockWrap from './block-wrap.vue'
import block from './block.vue'
import selectItem from './select-item.vue'
export default {
    install(app) {
        app.use(antd)
        app.component(blockWrap.name, blockWrap)
        app.component(block.name, block)
        app.component(selectItem.name, selectItem)
    }
}
