import view from './view'
import config from './config'
import util from 'pkg/util'

import { createStore } from 'vuex'

export default createStore({
    modules: {
        view,
        config
    },
    actions: {
        resize: util.debounce(() => {
            const e = document.createEvent('Event')
            e.initEvent('resize', true, true)
            window.dispatchEvent(e)
        }, 100)
    }
})