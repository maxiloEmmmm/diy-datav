
import blockWrap from './block-wrap.vue'

export default {
    install(app) {
        app.component(blockWrap.name, blockWrap)
    }
}
