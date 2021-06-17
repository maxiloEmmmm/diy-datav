export default {
    data() {
        return {
            app_mixin: {
                focus: {
                    in: false,
                    do: null
                }
            }
        }
    },
    methods: {
        async _initFocus() {
            this.app_mixin.focus.do = await this.$store.dispatch("view/addFocusItem", (isFocus) => {
                this.app_mixin.focus.in = isFocus
            })
        },
        _doFocus() {
            if(this.app_mixin.focus.do != null && !this.app_mixin.focus.in) {
                this.app_mixin.focus.do()
            }
        }
    }
}