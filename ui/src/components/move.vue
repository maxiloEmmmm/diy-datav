<script lang="tsx">
export default {
    render() {
        let context = this.$slots.default()
        return <div
            style="width:120px; height:120px"
        >
            {context}
        </div>
    },
    data() {
        return {
            inMove: false,
        }
    },
}
</script>