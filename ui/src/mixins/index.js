import focus from './focus'
import help from './help'
import config from './config'
import common from './common'
export default {
    install(app) {
        app.mixin(focus)
        app.mixin(help)
        app.mixin(config)
        app.mixin(common)
    }
}