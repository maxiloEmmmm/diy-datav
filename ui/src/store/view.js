import util from 'pkg/util'

const state = {
    dragBlockId: "",
    focus: {
        item: "",
        map: {}
    }
}

const mutations = {
    setDragBlockID(state, payload) {
        state.dragBlockId = payload
    },
    clearDragBlockID(state) {
        state.dragBlockId = ''
    },
}

const actions = {
    addFocusItem({state}, cb) {
        let uuid = util.uuid()
        state.focus.map[uuid] = cb
        return () => {
            state.focus.item = uuid

            for (let id in state.focus.map) {
                state.focus.map[id](id === uuid)
            }
        }
    }
}

export default {
    namespaced: true,
    state,
    mutations,
    actions,
}

