import view from './view'
import config from './config'

import { createStore } from 'vuex'

export default createStore({
    modules: {
        view,
        config
    }
})