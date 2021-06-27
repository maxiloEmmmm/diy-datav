import util from 'pkg/util'
import {api, type as apiType} from '@/api'

const clockStop = -1
const clockStart = 1

const state = {
    dragBlockId: "",

    focus: {
        item: "",
        map: {}
    },
    help: {},

    dataSet: {},
    clock: -1,
}

const getters = {
    help(state) {
        let rets = {}
        for (let typ in state.help) {
            rets[typ] = state.help[typ].filter(val => !val.disable)
        }
        return rets
    },
    fetchAble(state) {
        return state.clock != clockStop
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
    setData(state, payload = {id: "", data}) {
        state.dataSet[payload.id] = payload.data
    },
    updateDataRefresh(state, payload = {id: "", refresh: 1000 * 10}) {

    },
    loadData(state, payload = {id: "", refresh: 1000 * 10, cb: Function}) {
        state.dataSet[payload.id] = payload
    },
    addClock(state) {
        if (state.clock === clockStop) {
            return
        }
        state.clock += 1
    },
    activeClock(state) {
        state.clock = state.clock === clockStop ? clockStart : state.clock
    },
    stopClock(state) {
        state.clock = clockStop
    }
}

let fetchHandler = null

// TODO: support view config engine, websocket or xhr
let fetchEngine = api[apiType.Data]

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
    },
    getData({state}, id) {
        return state.dataSet[id]
    },
    fetchData({state, commit, getters}) {
        if (fetchHandler == null) {
            fetchHandler = setInterval(() => {
                commit('addClock')
                if(getters.fetchAble) {
                    for(let id in state.dataSet) {
                        const ds = state.dataSet[id]
                        if(ds.refresh % state.clock === 0 && fetchEngine !== null) {
                            fetchEngine(id)
                                .then(data => {
                                    ds.cb(data)
                                })
                        }
                    }
                }
            }, 1000)
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

