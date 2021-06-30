const state = {
    global: {},
    block: {
        type: "antv",
        config: ""
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

