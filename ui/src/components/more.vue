<script lang="tsx">
export default {
    name: 'more',
    render(){
        let cs = []
        for(let i = 0; i < this.count; i++) {
            cs.push(<ysz-list-item start={true} v-slots={{
                left: () => i + 1
            }}><a-space>{this.component(i)}<a-button size="small" onClick={e => this.onRemove(i)}>移除</a-button></a-space></ysz-list-item>)
        }
        return <div>
            <a-divider orientation="right"><a-button size="small" onClick={this.onNew}>新增</a-button></a-divider>
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
        initCount: {
            type: Number,
            default: 0
        },
        value: {
            type: Array,
            default() {
                return []
            }
        }
    },
    emits: ['add', 'remove'],
    watch: {
        initCount: {
            immediate: true,
            handler() {
                this.count = this.initCount
            }
        }
    },
    data() {
        return {count: 0}
    },
    methods: {
        onNew() {
            this.$emit('add', {
                count: this.count,
                done: () => {
                    this.count++
                }
            })
        },
        onRemove(index) {
            this.$emit('remove', {
                index,
                done: () => {
                    this.count--
                }
            })
        }
    }
}
</script>