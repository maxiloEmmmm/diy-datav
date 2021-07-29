<script lang="jsx">
export default {
    name: "permission",
    render(){
        return <ysz-module-widget
            v-slots={{
                title: () => <ysz-list-item>
                    <a-button onClick={this.onEmpty}>全空</a-button>
                    <a-divider type="vertical" />
                    <a-button onClick={this.onAll}>全选</a-button>
                </ysz-list-item>,
            }}
        >
            <a-tree
                checkable
                expanded-keys={this.expandedKeys}
                auto-expand-parent
                checked-keys={this.checkedKeys}
                tree-data={this.nodes}
                onSelect={this.onSelect}
                onExpand={this.onExpand}
                onCheck={this.onCheck}
            ></a-tree>
        </ysz-module-widget>
    },
    data(){
        return {
            expandedKeys: [],
            checkedKeys: this.keys
        }
    },
    props: {
        sub: String,
        objKey: String,
        actKey: String,
        nodes: Array,
        keys: Array,
        prefix: String,
        allNode: {
            type: Boolean,
            default: true
        }
    },
    methods: {
        onSelect(){
            console.log(arguments)
        },
        onExpand(keys){
            this.expandedKeys = keys
        },
        fetchAllKey(node){
            let keys = []
            node.forEach(n => {
                keys.push(n.key)
                if(Array.isArray(n.children) && n.children.length > 0) {
                    keys.push(...this.fetchAllKey(n.children))
                }
            })
            return keys
        },
        onAll(){
            this.checkedKeys = this.fetchAllKey(this.nodes)
            this.change(this.checkedKeys)
        },
        onEmpty(){
            this.checkedKeys = []
            this.change(this.fetchAllKey(this.nodes))
        },
        change(keys){
            let tmp = this.getNodeByKeys(this.nodes, [...keys])
            this.$api[this.$apiType.Policy]({
                add: tmp.nodes.map(n => `${this.sub},${this.prefix}.${n[this.objKey]},${n[this.actKey]},allow`),
                remove: tmp.xnodes.map(n => `${this.sub},${this.prefix}.${n[this.objKey]},${n[this.actKey]},allow`)
            })
        },
        getNodeByKeys(node, keys){
            let nodes = []
            let xnodes = []
            node.forEach(n => {
                let hasChild = Array.isArray(n.children) && n.children.length > 0
                let can = this.allNode || !hasChild
                if(keys.includes(n.key)) {
                    delete keys[keys.indexOf(n.key)]
                    if(can) {
                        nodes.push(n)
                    }
                }else {
                    if(can) {
                        xnodes.push(n)
                    }
                }
                if(hasChild) {
                    let tmp = this.getNodeByKeys(n.children, keys)
                    nodes.push(...tmp.nodes)
                    xnodes.push(...tmp.xnodes)
                }
            })
            return {nodes, xnodes}
        },
        onApply(){},
        onCheck(keys, state){
            this.checkedKeys = keys
            this.change(keys)
        }
    }
}
</script>