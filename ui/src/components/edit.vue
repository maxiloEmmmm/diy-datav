<script lang="tsx">
export default {
    render() {
        let context = this.$slots.default()
        return <div>
            {context}
        </div>
    },
    emits: ['do'],
    methods: {
        onMove(typ) {
            this.$emit('dp', {
                // or resize
                type: 'move',
                value: typ
            })
        }
    }
}
</script>