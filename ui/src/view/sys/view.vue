<script lang="jsx">
export default {
    render() {
        return <tool-curd
            models={this.models}
            columns={this.columns}
            fetch-url="view"
            onView={this.onView}
            onDesign={this.onDesign}
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
                {title: '删除', hide: '*', xhr: {method: 'delete'}}
            ]
        }
    },
    methods: {
        onDesign() {
            console.log(...arguments)
        },
        onView() {
            console.log(...arguments)
        }
    }
}
</script>