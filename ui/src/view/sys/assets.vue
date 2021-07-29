
<script lang="jsx">
import {CloudUploadOutlined} from '@ant-design/icons-vue'
export default {
    render() {
        return <>
            <tool-curd ref="curd" layout={[
                {key: 'default', col: 2},
                {key: 'single', col: 1}
            ]} models={this.models} columns={this.columns} fetch-url="assets" onUpload={this.onUploadTouch}/>
            <tool-form ref="form" vModel={[this.inUpload, 'show']}/>
        </>
    },
    data() {
        return {
            columns: [
                {title: 'ID', field: 'id', key: true},
                {title: '类型', field: 'type', type: 'select', option: {
                    selectOptions: () => this.store.typeOptions
                }},
                {title: '路径', field: 'path', filter: true},
                {title: '拓展', field: 'ext'},
            ],
            uploadColumns: [
                {title: '类型', field: 'type', type: 'select', option: {
                    selectOptions: () => this.store.typeOptions
                }, onChange: v => {
                    this.dataform.type = v
                }},
                {title: '文件', field: 'file', type: 'customer', option: {
                    customer_form: () => {
                        return <a-upload
                            before-upload={this.onUpload}
                            showUploadList={false}
                        >
                            <a-button size="small" disabled={!this.dataform.type}> <CloudUploadOutlined /> 点击上传 </a-button>
                            {this.loading ? <a-progress percent={this.percent} size="small" status="active" /> : null}
                        </a-upload>
                    }
                }}
            ],
            loading: false,
            percent: 0,
            inUpload: false,
            dataform: {
                type: ''
            },
            store: {
                typeOptions: []
            },
            models: [
                {title: '更新', disabled: ['id'], xhr: {method: 'patch'}},
                {title: '新增', hide: ['id'], omit: ['id'], xhr: {method: 'post'}, dispatchArea: 'topBar'},
                {title: '上传', type: 'api', api: 'event', key: 'upload', dispatchArea: 'topBar'},
                {title: '删除', hide: '*', xhr: {method: 'delete'}}
            ]
        }
    },
    created() {
        this.$api[this.$apiType.AssetsTypeList]()
            .then(response => {
                this.store.typeOptions = response.data
            })
    },
    mounted() {
        this.$nextTick(() => {
            this.$refs.form.setFields(this.uploadColumns)
            this.$refs.form.setModels([
                {key: 'default'}
            ])
            this.$refs.form.setData(this.cfg)
            this.$refs.form.setModel('default')
        })
    },
    methods: {
        onUploadTouch() {
            this.inUpload = true
        },
        async onUpload(file) {
            let fd = new FormData()
            fd.append('file', file)
            new Promise(async (ok) => {
                this.loading = true
                this.percent = 50
                await this.$api[this.$apiType.ViewUploadBG](this.dataform.type, fd)
                this.percent = 100
                this.loading = false
                this.$message.success("ok~")
                this.$refs.curd.refresh()
                this.inUpload = false
            })
            return false
        }
    }
}
</script>