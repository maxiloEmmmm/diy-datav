export default {
    methods: {
        mixinConfigShow() {
            this.$store.commit('config/SetActive', true)
        },
        mixinConfigHidden() {
            this.$store.commit('config/SetActive', false)
        },
        mixinSetConfigKey(key = "") {
            this.$store.commit("config/SetBlockKey", key)
        },
        mixinSetConfigType(typ = "") {
            this.$store.commit("config/SetBlockType", typ)
        },
        mixinSetConfigTypeAndConfig(typ = "", config = "") {
            this.mixinSetConfigType(typ)
            this.mixinSetConfigConfig(config)
        },
        mixinSetConfigConfig(config = "") {
            this.$store.commit("config/SetBlockConfig", config)
        }
    }
}