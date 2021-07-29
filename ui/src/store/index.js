import view from './view'
import config from './config'
import auth from './auth'
import util from 'pkg/util'

import { createStore } from 'vuex'

export default createStore({
    modules: {
        auth,
        config,
        view,
    },
    actions: {
        resize: util.debounce(() => {
            const e = document.createEvent('Event')
            e.initEvent('resize', true, true)
            window.dispatchEvent(e)
        }, 100)
    }
})