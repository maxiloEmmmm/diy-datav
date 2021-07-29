<script lang="jsx">
export default {
    name: "bg-pick",
    render(){
        return <a-space>
            <a-popover
                title="背景选择"
                trigger="click"
                visible={this.visible}
                onVisibleChange={e => this.visible = e}
                v-slots={{
                    content: () => <a>
                        <no-scroll style="width:400px;height: 200px">
                            <ysz-list row group={4}>
                                {this.bgs.map(bg => <div
                                    style={`width:180px; height: 180px; background-size: 100% 100%;background-image:url(${this.$api_url}/assets-file/${bg.id}?token=${this.$store.state.auth.token})`}
                                    onClick={() => this.onClick(bg)}
                                    className={`bg-pick ${bg.id === this.value ? "active" : ""}`}/>)}
                            </ysz-list>
                            {this.bgs.length === 0 ? <a-empty/> : null}
                        </no-scroll>
                    </a>
                }}
            >
                <a-button size="small" type="primary">
                    {this.value !== 0 ? "已选" : "选择"}
                </a-button>
            </a-popover>
        </a-space>
    },
    props: {
        value: {
            type: Number,
            default: 0
        }
    },
    created(){
        this.fetch()
    },
    emits: ['update:value'],
    data(){return {bgs: [], visible: false}},
    methods: {
        onClick(bg){
            this.$emit("update:value", bg.id)
            this.visible = false
        },
        fetch(){
            this.$api[this.$apiType.ViewBGAssets]()
                .then(response => this.bgs = response.data)
        },
    }
}
</script>

<style lang="scss" scoped>
.bg-pick {
    &:hover {
        border: 1px pink dashed; padding: 2px; border-radius: 6px;
    }
    &.active {
        border: 1px pink dashed; background: pink; color: #fff;border-radius: 6px;padding: 4px;
    }
}
</style>