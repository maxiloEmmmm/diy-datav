const state = {
    global: {},
    block: {
        type: "antv",
        config: '',
        key: ''
    },
    show: true
}

const getters = {}

const mutations = {
    SetBlockConfig(state, config) {
        state.block.config = config
    },
    SetBlockType(state, type) {
        state.block.type = type
    },
    SetBlockKey(state, key) {
        state.block.key = key
    },
    SetActive(state, active) {
        state.show = active
    }
}

const actions = {}

export default {
    namespaced: true,
    state,
    getters,
    mutations,
    actions,
}

