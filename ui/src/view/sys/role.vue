<script lang="jsx">
export default {
    render(){
        return <>
            <tool-curd
                columns={this.cols}
                models={this.models}
                title="角色"
                fetchTransform={this.fetchTransform}
                fetchUrl={this.fetchUrl}
                onUser={this.onUser}
                onApi={this.onApi}
                onMenu={this.onMenu}
                onMap={this.onMap}
            />
            <a-modal destroyOnClose={true} title={this.current.title} vModel={[this.current.show, 'visible']} onOk={() => this.current.show = false}>
                <permission
                    {...this.current}
                />
            </a-modal>
            <a-modal destroyOnClose={true} title="分配角色" vModel={[this.user.show, 'visible']} onOk={() => this.user.show = false}>
                <role-user
                    role={this.user.role}
                />
            </a-modal>
        </>
    },
    data(){
        return {
            cols: [
                {title: "名字", field: "role", key: true},
            ],
            models: [
                {title: "新增", dispatchArea: "topBar", key: "new", xhr: {method: "post"}},
                {title: "删除", hide: '*', key: "delete", xhr: {method: "delete"}},
                {title: "管理用户", type: "api", key: "user", api: "event"},
                {title: "接口权限", type: "api", key: "api", api: "event"},
                {title: "菜单权限", type: "api", key: "menu", api: "event"},
                {title: "可视化权限", type: "api", key: "map", api: "event"}
            ],
            fetchUrl: `/casbin/role`,
            user: {
                show: false,
                role: ''
            },
            current: {
                show: false,
                role: '',
                prefix: '',
                objKey: '',
                actKey: '',
                nodes: [],
                keys: [],
                allNode: false,
                title: '',
                width: 600
            }
        }
    },
    methods: {
        onUser({item}){
            this.user = {show: true, role: item.role}
        },
        onMap({item}){
            this.$api[this.$apiType.PermissionView](item.role).then(response => {
                let nodes = []
                let keys = []
                response.data.forEach(view => {
                    let m = {title: view.title, key: view.id, eft: view.eft, act: "access"}
                    if(m.eft) {
                        keys.push(m.key)
                    }
                    nodes.push(m)
                })
                this.current = {
                    prefix: "pr-view", sub: item.role, objKey: "key", actKey: "act", nodes, keys, allNode: false,
                    title: `可视化权限, 对象: ${item.role}`, width: 600, show: true
                }
            })
        },
        onApi({item}){
            const h = this.$createElement;
            this.$api[this.$apiType.PermissionApi](item.role).then(response => {
                let nodes = []
                let keys = []
                Object.keys(response.data).forEach(mod => {
                    let m = {title: mod, key: mod, eft: true}
                    m.children = response.data[mod].map(api => {
                        m.eft = m.eft && api.eft
                        api.title = `${api.method} - ${api.path}`
                        api.key = `${mod}-${api.title}`
                        if(api.eft) {
                            keys.push(api.key)
                        }
                        return api
                    })
                    if(m.eft) {
                        keys.push(m.key)
                    }
                    nodes.push(m)
                })
                this.current = {
                    prefix: "pr-router", sub: item.role, objKey: "path", actKey: "method", nodes, keys, allNode: false,
                    title: `接口权限, 对象: ${item.role}`, width: 600, show: true
                }
            })
        },
        depMenu(menus, keys) {
            return menus.map(d => {
                let m = {...d, title: d.title, key: d.id, eft: d.eft, act: "access"}
                if (m.eft) {
                    keys.push(m.key)
                }
                if(d.edges.children) {
                    m.children = this.depMenu(d.edges.children, keys)
                }
                return m
            })
        },
        onMenu({item}){
            const h = this.$createElement;
            this.$api[this.$apiType.PermissionMenu](item.role).then(response => {
                let keys = []
                let nodes = this.depMenu(response.data, keys)
                this.current = {
                    prefix: "pr-menu", sub: item.role, objKey: "url", actKey: "act", nodes, keys, allNode: false,
                    title: `菜单权限, 对象: ${item.role}`, width: 600, show: true
                }
            })
        },
        fetchTransform(response){
            response.data = {
                data: response.data.map(name => {
                    return {role: name}
                }),
                total: response.data.length
            }
            return response
        }
    }
}
</script>