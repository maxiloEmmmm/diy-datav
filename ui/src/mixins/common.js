export default {
    methods: {
        mixinDispatchWindowResize() {
            this.$store.dispatch('resize')
        }
    }
}