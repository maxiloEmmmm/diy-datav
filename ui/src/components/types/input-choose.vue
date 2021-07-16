<script lang="tsx">
export default {
    render() {
        return <a-checkbox options={this.options} value={this.value} onChange={this.onChange}/>
    },
    computed: {
        options() {
            return (this.inputs || []).map(input => ({label: input.title, value: input.id}))
        }
    },
    emits: ['update:value'],
    props: {
        inputs: {
            type: Array,
            default: () => []
        },
        value: {
            type: Array,
            default: () => []
        }
    },
    methods: {
        onChange(vals) {
            console.log(vals)
        }
    }
}
</script>