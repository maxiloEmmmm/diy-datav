
<script lang="jsx">
export default {
    render() {
        return <tool-curd models={this.models} columns={this.columns} fetch-url="user"/>
    },
    data() {
        return {
            columns: [
                {title: 'ID', field: 'id', type: 'number', key: true},
                {title: '用户名', field: 'username', filter: true},
                {title: '密码', field: 'password', hidden: true},
                {title: '启用', field: 'enable', type: 'switch', default: true}
            ],
            models: [
                {title: '更新', disabled: ['id', 'username'], omit: ['id'], xhr: {method: 'patch'}},
                {title: '新增', hide: ['id'], omit: ['id'], xhr: {method: 'post'}, dispatchArea: 'topBar'},
                {title: '删除', hide: '*', xhr: {method: 'delete'}}
            ]
        }
    },
}
</script>