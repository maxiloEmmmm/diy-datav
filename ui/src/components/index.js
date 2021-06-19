import 'ant-design-vue/dist/antd.css'
import antd from 'ant-design-vue'
import blockWrap from './block-wrap.vue'

export default {
    install(app) {
        app.use(antd)
        app.component(blockWrap.name, blockWrap)
    }
}
