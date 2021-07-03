export default {
    methods: {
        mixinConfigShow() {
            this.$store.commit('config/SetActive', true)
        },
        mixinConfigHidden() {
            this.$store.commit('config/SetActive', false)
        },
        mixinSetConfigTypeAndConfig(typ = "", config = "") {
            this.$store.commit('config/SetBlockType', typ)
            this.$store.commit('config/SetBlockConfig', config)
            this.$store.commit('config/AddHistory')
        },
        mixinSetConfigConfig(config = "") {
            this.$store.commit('config/SetBlockConfig', config)
            this.$store.commit('config/AddHistory')
        },
        mixinSetConfigKey(key = "") {
            this.$store.commit("config/SetBlockKey", key)
        },
        mixinClearConfig(key = "") {
            this.mixinSetConfigKey()
            this.mixinSetConfigTypeAndConfig()
            this.mixinClearHistory()
        },
        mixinClearHistory() {
            this.$store.commit("config/ClearHistory")
        },
        mixinBackHistory() {
            this.$store.commit("config/GoHistory", -1)
        },
        mixinForwardHistory() {
            this.$store.commit("config/GoHistory", 1)
        }
    }
}