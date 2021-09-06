import util from 'pkg/util'
import {api, type as apiType} from '@/api'
import {PositionType} from 'type'
const clockStop = -1
const clockStart = 0

const state = {
    radio: 1,
    focus: {
        item: "",
        map: {}
    },
    help: {},
    blockMoving: "",
    adsorption: {
        design: {
            lineIndex: -1,
            pos: PositionType()
        },
        grid: {
            blockKey: '',
            row: 0,
            col: 0,
            cb: (blockKey) => {},
            pos: PositionType()
        }
    },
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
        return state.clock !== clockStop
    }
}

function normalHelp(help) {
    if(help.disable === undefined) {
        help.disable = false
    }
    return help
}

const mutations = {
    setBlockMoving(state, moving) {
        // TODO: add getter to like bool
        state.blockMoving = moving
    },
    setAdsorptionDesign(state, payload = {lineIndex: 0, pos: PositionType()}) {
        state.adsorption.design.lineIndex = payload.lineIndex
        state.adsorption.design.pos = payload.pos
    },
    setAdsorptionGrid(state, payload = {blockKey: "", row: 0, col: 0, blockOn: Function, pos: PositionType()}) {
        state.adsorption.grid.blockKey = payload.blockKey
        state.adsorption.grid.row = payload.row
        state.adsorption.grid.col = payload.col
        state.adsorption.grid.cb = payload.blockOn
        state.adsorption.grid.pos = payload.pos
    },
    addHelp(state, payload = {typ: "", helps: []}) {
        if(!state.help[payload.typ]) {
            state.help[payload.typ] = []
        }

        state.help[payload.typ].push(...payload.helps.map(normalHelp).filter(help => !(state.help[payload.typ] || []).some((val) => {
            if(val.key === help.key) {
                return true
            }
        })))
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
    updateDataRefresh(state, payload = {id: "", refresh: 10}) {

    },
    loadData(state, payload = {input: {key: String}, refresh: 10, cb: Function}) {
        state.dataSet[payload.input.key] = payload
    },
    clearLoadData(state) {
        state.dataSet = {}
    },
    addClock(state) {
        if (state.clock === clockStop) {
            return
        }

        if (state.clock >= 3600) {
            state.clock = clockStart
        }

        state.clock += 1
    },
    activeClock(state) {
        state.clock = state.clock === clockStop ? clockStart : state.clock
    },
    stopClock(state) {
        state.clock = clockStop
    },
    setRadio(state, radio) {
        state.radio = radio
    }
}

let fetchHandler = null

// TODO: support view config engine, websocket or xhr
let fetchEngine = null

export const SetFetchEngine = t => fetchEngine = t

let fetchTmpEchoEngine = null

export const SetFetchTmpEchoEngine = t => fetchTmpEchoEngine = t

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
            commit('activeClock')
            fetchHandler = setInterval(() => {
                if(getters.fetchAble) {
                    for(let key in state.dataSet) {
                        const ds = state.dataSet[key]
                        if(state.clock % ds.refresh === 0) {
                            let pro = null
                            if(ds.input.id) {
                                if(fetchEngine !== null) {
                                    pro = fetchEngine(ds.input.id)
                                }
                            }else {
                                if(fetchTmpEchoEngine !== null) {
                                    pro = fetchTmpEchoEngine({
                                        type: ds.input.type,
                                        config: ds.input.config
                                    })
                                }
                            }

                            pro.then(response => {
                                ds.cb(response.data)
                            })
                        }
                    }
                }
                commit('addClock')
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

