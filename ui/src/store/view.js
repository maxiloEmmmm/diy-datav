const state = {
    dragBlockId: "",
}

const mutations = {
    setDragBlockID(state, payload) {
        state.dragBlockId = payload
    },
    clearDragBlockID(state) {
        state.dragBlockId = ''
    }
}

export default {
    namespaced: true,
    state,
    mutations,
}

