<script lang="tsx">
export default {
    name: 'more',
    render(){
        let cs = []
        for(let i = 0; i < this.count; i++) {
            cs.push(this.component(i))
        }
        return <div>
            <a-divider orientation="right"><a-button onClick={this.onNew}>新增</a-button></a-divider>
            {cs}
            {this.$slots.default}
        </div>
    },
    props: {
        component: {
            type: Function,
            default() {
                return <div/>
            }
        },
        value: {
            type: Array,
            default() {
                return []
            }
        }
    },
    emits: ['change'],
    data() {
        return {count: 0}
    },
    methods: {
        onNew() {
            this.$emit('change', {
                count: this.count,
                done: () => {
                    this.count++
                }
            })
        }
    }
}
</script>