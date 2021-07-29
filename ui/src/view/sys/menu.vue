<script lang="jsx">
export default {
    render(){
        return <ysz-module-widget>
            <ysz-list-item
                v-slots={{
                    left: () => "菜单"
                }}
            >
                <a-button size="small" onClick={() => this.addModal = true}>新增</a-button>
            </ysz-list-item>
            <a-tree
                showLine
                tree-data={this.nodes}
                replaceFields={{key: "id"}}></a-tree>

            <a-modal title="顶级菜单" vModel={[this.addModal, 'visible']} title="保存" onOk={this.onAddTop}>
                <ysz-list-item v-slots={{
                    left: () => '名字'
                }}>
                    <a-input value={this.dataform.title} onChange={e => this.dataform.title = e.target.value } size="small"/>
                </ysz-list-item>
                <ysz-list-item v-slots={{
                    left: () => '链接'
                }}>
                    <a-input value={this.dataform.url} onChange={e => this.dataform.url = e.target.value } size="small"/>
                </ysz-list-item>
            </a-modal>
        </ysz-module-widget>
    },
    created(){this.fetch()},
    data(){
        return {
            nodes: [],
            dataform: {
                title: "",
                url: ""
            },
            addModal: false,
        }
    },
    methods: {
        onAddTop(){
            if(this.dataform.title && this.dataform.url) {
                this.do("add", null)
            }
        },
        depMenu(data){
            data.forEach(d => {
                d._title = d.title
                d.title = (obj) => {
                    return <a-popover
                        placement="bottomRight" trigger="click" title="操作"
                        v-slots={{
                            content: () => <tw-list-item1 index indexBorder items={[
                                {title: <ysz-module-widget title="修改节点">
                                        <ysz-list-item
                                            v-slots={{
                                                left: () => "名字"
                                            }}
                                        >
                                            <a-input value={this.dataform.title} onChange={e => this.dataform.title = e.target.value } size="small"/>
                                        </ysz-list-item>
                                        <ysz-list-item
                                            v-slots={{
                                                left: () => "链接"
                                            }}
                                        >
                                            <a-input value={this.dataform.url} onChange={e => this.dataform.url = e.target.value } size="small"/>
                                        </ysz-list-item>
                                        <a-button size="small"  onClick={() => this.do("update", obj.id)}>修改!</a-button>
                                    </ysz-module-widget>},
                                {title: <a-button size="small" onClick={() => this.do("del", obj.id)}>删除节点</a-button>},
                                {title: <ysz-module-widget title="新增子节点">
                                        <ysz-list-item
                                            v-slots={{
                                                left: () => "名字"
                                            }}
                                        >
                                            <a-input value={this.dataform.title} onChange={e => this.dataform.title = e.target.value } size="small"/>
                                        </ysz-list-item>
                                        <ysz-list-item
                                            v-slots={{
                                                left: () => "链接"
                                            }}
                                        >
                                            <a-input value={this.dataform.url} onChange={e => this.dataform.url = e.target.value } size="small"/>
                                        </ysz-list-item>
                                        <a-button size="small" onClick={() => this.do("add", obj.id)}>新增!</a-button>
                                    </ysz-module-widget>},
                            ]}></tw-list-item1>
                        }}
                    >
                        <a-button size="small" onClick={() => {
                            this.dataform = {title: obj._title, url: obj.url}
                        }}>
                            {obj._title}
                        </a-button>
                    </a-popover>
                }
                d.children = d.edges.Children
                if(d.children != null) {
                    this.depMenu(d.children)
                }
            })
            return data
        },
        do(action, id){
            let p = null
            switch(action) {
                case "add":
                    p = this.$api[this.$apiType.MenuAdd]({...this.dataform,
                        edges: {
                            parent: id ? {id} : undefined
                        }
                    })
                    break
                case "update":
                    p = this.$api[this.$apiType.MenuUpdate](id, this.dataform)
                    break
                case "del":
                    p = this.$api[this.$apiType.MenuDelete](id)
                    break
            }
            p.then(() => {
                this.fetch()
                this.$message.success("ok")
                this.dataform = {title: "", url: ""}

                switch(action) {
                    case "add":
                        this.addModal = false
                        break
                    case "update":
                    case "del":
                }
            })
        },
        fetch(){
            this.$api[this.$apiType.MenuList]().then(response => {
                this.nodes = this.depMenu(response.data.data)
            })
        }
    }

}
</script>