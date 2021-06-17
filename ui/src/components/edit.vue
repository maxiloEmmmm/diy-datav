<script lang="tsx">
export default {
    render() {
        let context = this.$slots.default()
        return <div>
            {context}
        </div>
    }
}
</script>