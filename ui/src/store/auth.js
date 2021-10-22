const TokenKey = "diy-datav.store.token"
const ForwardKey = "diy-datav.store.forward"
const storageEngine = window.localStorage
const state = {
    token: storageEngine.getItem(TokenKey),
    forward: storageEngine.getItem(ForwardKey)
}

const getters = {}

const mutations = {
    SetToken(state, token) {
        state.token = token
        storageEngine.setItem(TokenKey, token)
    },
    SetForward(state, forward) {
        state.forward = forward
        storageEngine.setItem(ForwardKey, forward)
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

