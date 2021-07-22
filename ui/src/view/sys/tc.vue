
<script lang="jsx">
import tcMod from './tc-mod/mod.js'
export default {
    render() {
        return <tool-curd layout={[
            {key: 'default', col: 2},
            {key: 'single', col: 1}
        ]} models={this.models} columns={this.columns} fetch-url="tc"/>
    },
    data() {
        const configMod = function(edit = true) {
            return (value, row, option) => {
                let mod = tcMod[row.type]
                if (mod === undefined) {
                    return '未知tc类型'
                }

                return <mod config={value} edit={edit}/>
            }
        }
        return {
            columns: [
                {title: 'ID', field: 'id'},
                {title: '标题', field: 'title', filter: true},
                {title: '配置', field: 'config', type: 'customer', layout_key: 'single', option: {
                    customer_form: configMod(true),
                    customer_view: configMod(false)
                }},
                {title: '类型', field: 'type', type: 'select', option: {
                    selectOptions: () => this.store.typeOptions
                }, filter: true}
            ],
            store: {
                typeOptions: []
            },
            models: [
                {title: '更新', disabled: ['id'], xhr: {method: 'patch'}},
                {title: '新增', hide: ['id'], xhr: {method: 'post'}, dispatchArea: 'topBar'},
                {title: '删除', hide: '*', xhr: {method: 'delete'}}
            ]
        }
    },
    created() {
        this.$api[this.$apiType.KindList]("1")
            .then(response => {
                this.store.typeOptions = response.data.data
            })
    }
}
</script>