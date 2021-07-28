<script lang="jsx">
import util from 'pkg/util'
export default {
    render() {
        return <tool-curd
            models={this.models}
            columns={this.columns}
            fetch-url="view"
            onView={this.onView}
            onDesign={this.onDesign}
            onNew={this.onNew}
        />
    },
    data() {
        return {
            columns: [
                {title: 'ID', field: 'id', key: true},
                {title: '描述', field: 'desc', filter: true},
            ],
            store: {
                typeOptions: []
            },
            models: [
                {title: '阅览', type: 'api', api: 'event', key: 'view'},
                {title: '设计', type: 'api', api: 'event', key: 'design'},
                {title: '删除', hide: '*', xhr: {method: 'delete'}},
                {title: '新建', type: 'api', api: 'event', key: 'new', dispatchArea: 'topBar'}
            ]
        }
    },
    methods: {
        onDesign({item}) {
            util.newPage(`/view/design/${item.id}`)
        },
        onView({item}) {
            util.newPage(`/view/show/${item.id}`)
        },
        onNew() {
            util.newPage(`/view/design`)
        }
    }
}
</script>