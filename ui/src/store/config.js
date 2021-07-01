const state = {
    global: {},
    block: {
        type: '',
        config: '',
        key: '',
        history: [],
        current: -1,
    },
    show: false
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
        state.block.history = []
        state.block.key = key
    },
    SetActive(state, active) {
        state.show = active
    },
    ClearHistory() {
        state.block.history = []
    },
    AddHistory(state) {
        if(state.block.current < state.block.history.length - 1) {
            state.block.history = state.block.history.slice(0, state.block.current)
        }

        state.block.history.push({
            type: state.block.type,
            config: state.block.config
        })

        state.block.current = state.block.history.length - 1
    },
    GoHistory(state, step) {
        if(!state.block.key || step === 0) {
            return
        }

        const hl = state.block.history.length
        const targetIndex = hl - 1 + step
        if (targetIndex > hl - 1) {
            return
        }

        state.block.current = targetIndex
        const record = state.block.history[state.block.current]
        state.block.type = record.type
        state.block.config = record.config
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

