import util from 'pkg/util'

const state = {
    dragBlockId: "",
    focus: {
        item: "",
        map: {}
    },
    help: {}
}

const getters = {
    help(state) {
        let rets = {}
        for (let typ in state.help) {
            rets[typ] = state.help[typ].filter(val => !val.disable)
        }
        return rets
    }
}

function normalHelp(help) {
    if(help.disable === undefined) {
        help.disable = false
    }
    return help
}

const mutations = {
    setDragBlockID(state, payload) {
        state.dragBlockId = payload
    },
    clearDragBlockID(state) {
        state.dragBlockId = ''
    },
    addHelp(state, payload = {typ: "", helps: []}) {
        if(!state.help[payload.typ]) {
            state.help[payload.typ] = []
        }

        state.help[payload.typ].push(...payload.helps.map(normalHelp))
    },
    setHelp(state, payload = {typ: "", helps: []}) {
        state.help[payload.typ] = payload.helps.map(normalHelp)
    },
    disableHelp(state, payload = {typ: "", key: ""}) {
        (state.help[payload.typ] || []).forEach((val) => {
            if(val.key === payload.key) {
                val.disable = true
            }
        })
    },
    activeHelp(state, payload = {typ: "", key: ""}) {
        (state.help[payload.typ] || []).forEach((val) => {
            if(val.key === payload.key) {
                val.disable = false
            }
        })
    },
    removeHelp(state, payload = {typ: "", key: ""}) {
        state.help[payload.typ] = (state.help[payload.typ] || []).filter((val) => {
           return  val.key !== payload.key
        })
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
    getters,
    mutations,
    actions,
}

