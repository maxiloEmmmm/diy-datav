<script lang="jsx">
import util from 'pkg/util'
export default {
    render() {
        return <tool-curd models={this.models} columns={this.columns} onAccess={this.onAccess} fetch-url="share"/>
    },
    data() {
        return {
            store: {
                viewOptions: []
            },
            columns: [
                {title: 'ID', field: 'id', type: 'number', key: true},
                {title: '图', field: 'edges.view.id', type: 'select', option: {
                    selectOptions: () => this.store.viewOptions
                }},
                {title: '操作人员', field: 'edges.creator.username', filter: true},
                {title: '结束日期', field: 'end_at', type: 'datetimepick', option: {showTime: true}},
            ],
            models: [
                {title: '更新', hide: ['id', 'edges.creator.username', 'edges.view.id'], omit: ['id', 'edges.creator.username'], xhr: {method: 'patch'}},
                {title: '新增', hide: ['id', 'edges.creator.username'], omit: ['id', 'edges.creator.username'], xhr: {method: 'post'}, dispatchArea: 'topBar'},
                {title: '删除', hide: '*', xhr: {method: 'delete'}},
                {title: "访问连接", type: "api", key: "access", api: "event"}
            ]
        }
    },
    created() {
        this.$api[this.$apiType.ViewList]()
            .then(response => {
                this.store.viewOptions = response.data.data.map(view => ({label: view.desc, value: view.id}))
            })
    },
    methods: {
        onAccess({item}) {
            this.$notification.info({
                message: `${util.baseURL()}#/view/share/${item.id}`
            })
        }
    }
}
</script>