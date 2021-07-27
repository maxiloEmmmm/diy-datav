
<script lang="jsx">
import tcMod from './tc-mod/mod.js'
export default {
    render() {
        return <tool-curd layout={[
            {key: 'default', col: 2},
            {key: 'single', col: 1}
        ]} models={this.models} columns={this.columns} fetch-url="typeconfig"/>
    },
    data() {
        return {
            columns: [
                {title: 'ID', field: 'id', key: true},
                {title: '标题', field: 'title', filter: true},
                {title: '配置', field: 'config', type: 'customer', layout_key: 'single', option: {
                    customer_form: (props) => {
                        let mod = tcMod[props.item.type]
                        if (mod === undefined) {
                            return <a-alert message="请选择具体类型" type="info" />
                        }

                        return <mod config={props.value} {...props}/>
                    },
                }, default: '{}'},
                {title: '类型', field: 'type', type: 'select', option: {
                    selectOptions: () => this.store.typeOptions
                }, filter: true}
            ],
            store: {
                typeOptions: []
            },
            models: [
                {title: '更新', disabled: ['id'], xhr: {method: 'patch'}},
                {title: '新增', hide: ['id'], omit: ['id'], xhr: {method: 'post'}, dispatchArea: 'topBar'},
                {title: '删除', hide: '*', xhr: {method: 'delete'}}
            ]
        }
    },
    created() {
        this.$api[this.$apiType.KindList]()
            .then(response => {
                this.store.typeOptions = response.data
            })
    }
}
</script>