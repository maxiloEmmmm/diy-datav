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
        async mixinInitFocus() {
            this.app_mixin.focus.do = await this.$store.dispatch("view/addFocusItem", (isFocus) => {
                this.app_mixin.focus.in = isFocus
            })
        },
        mixinDoFocus() {
            if(this.app_mixin.focus.do != null && !this.app_mixin.focus.in) {
                this.app_mixin.focus.do()
            }
        }
    }
}