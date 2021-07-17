<script lang="tsx">
export default {
    render() {
        return <a-radio-group options={this.options} value={this.value} onChange={this.onChange}/>
    },
    computed: {
        options() {
            return [
                {value: -1, label: '无数据'},
                ...(this.inputs || []).map((input, index) => ({label: input.title ? input.title : index + 1, value: index}))
            ]
        }
    },
    emits: ['change'],
    props: {
        inputs: {
            type: Array,
            default: () => []
        },
        value: {
            type: Number,
            default: -1
        }
    },
    methods: {
        onChange(e) {
            this.$emit('change', e.target.value)
        }
    }
}
</script>