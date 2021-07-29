import 'ant-design-vue/dist/antd.css'
import antd from 'ant-design-vue'
import blockWrap from './block-wrap.vue'
import block from './block.vue'
import selectItem from './select-item.vue'
import more from './more.vue'
import colorPick from './color-pick.vue'
import configBar from './config-bar.vue'
import antdTool from 'pkg/antd-tool'
import bgPick from './bg-pick.vue'
import permission from './permission.vue'
import roleUser from './role_user.vue'
export default {
    install(app) {
        app.use(antdTool)
        app.use(antd)
        app.component(blockWrap.name, blockWrap)
        app.component(block.name, block)
        app.component(selectItem.name, selectItem)
        app.component(more.name, more)
        app.component(colorPick.name, colorPick)
        app.component(configBar.name, configBar)
        app.component(bgPick.name, bgPick)
        app.component(permission.name, permission)
        app.component(roleUser.name, roleUser)
    }
}
