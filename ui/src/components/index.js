import 'ant-design-vue/dist/antd.css'
import 'maxilo-vue-ysz-ui/lib/ysz-ui.css'
import antd from 'ant-design-vue'
import blockWrap from './block-wrap.vue'
import block from './block.vue'
import selectItem from './select-item.vue'
import yszUITool from 'maxilo-vue-ysz-ui/lib/ysz-ui'
export default {
    install(app) {
        app.use(antd)
        app.use(yszUITool)
        app.component(blockWrap.name, blockWrap)
        app.component(block.name, block)
        app.component(selectItem.name, selectItem)
    }
}
