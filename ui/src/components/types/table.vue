<script lang="jsx">
import {TableConfig, TableConfigParse} from 'type'
import util from 'pkg/util'
export default {
    props: {
        config: {
            type: Object,
            default() {
                return TableConfig()
            }
        },
        data: {
            type: Array,
            default() {
                return []
            }
        },
    },
    render() {
        return <tool-curd preview columns={this.fields} ref="table"/>
    },
    data() {
        return {
            fields: [],
            cfg: TableConfig(),
        }
    },
    watch: {
        config: {
            immediate: true,
            deep: true,
            handler() {
                this.parse()
            }
        },
        data: {
            deep: true,
            immediate: true,
            handler(val) {
                console.log(val)
                this.getData()
            }
        },
    },
    methods: {
        getData() {
            let data = []
            if(util.isArray(this.data)) {
                data = this.data[this.cfg.dataIndex]
            }

            let t = []

            if(util.isArray(data)) {
                t = data
            }

            let fields = []
            if(t.length > 0) {
                const obj = t[0]

                for(let k in obj) {
                    fields.push({
                        field: k,
                        title: k
                    })
                }
            }
            this.fields = fields

            this.$nextTick(() => {
                this.$refs.table.render({
                    status: 200,
                    data: {
                        data: t,
                        total: t.length
                    }
                })
                this.$refs.table.pageRender({
                    current: 1,
                    pageSize: t.length
                }, {})
            })
        },
        parse() {
            try {
                this.cfg = TableConfigParse(this.config)
            }catch(e) {
                console.log('static table config parse failed', e, this.config)
            }
        },
    }
}
</script>

<style lang="scss" scoped>
.st {
    width: 100%; height: 100%; display: flex; flex-direction: row;

    &.st-bold {font-weight: 800}
    &.st-italic {font-style: italic}
    &.st-underline {text-decoration: underline;}
}
</style>