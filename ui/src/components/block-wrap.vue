<script lang="jsx">
export default {
    name: 'block-wrap',
    render(){
        let context = this.$slots.default()
        return <div {...this.$attrs}>
            {context}
        </div>
    }
}
</script>