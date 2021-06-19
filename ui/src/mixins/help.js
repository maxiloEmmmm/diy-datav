export const Module = {
    ViewBlock: "view-block"
}

export default {
    methods: {
        mixinAddHelp(typ = '', helps = []) {
            this.$store.commit("view/addHelp", {
                typ, helps
            })
        },
        mixinSetHelp(typ = '', helps = []) {
            this.$store.commit("view/setHelp", {
                typ, helps
            })
        },
        mixinDisableHelp(typ = '', key = "") {
            this.$store.commit('view/disableHelp', {
                typ, key
            })
        },
        mixinActiveHelp(typ = '', key = "") {
            this.$store.commit('view/activeHelp', {
                typ, key
            })
        },
        mixinRemoveHelp(typ = '', key = "") {
            this.$store.commit('view/removeHelp', {
                typ, key
            })
        },
    }
}