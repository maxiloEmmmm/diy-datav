import focus from './focus'
import help from './help'
export default {
    install(app) {
        app.mixin(focus)
        app.mixin(help)
    }
}