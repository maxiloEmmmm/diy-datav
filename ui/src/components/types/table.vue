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
        return <no-scroll style="height: 100%" scroll-top={this.st} ref="scroll" onScrollTop={this.onScrollTop}>
            <table style="width: 100%">
                <thead style={{transform: `translateY(${this.stNumber}px)`, backgroundColor: this.cfg.headerBGC, color: this.cfg.headerC}}>
                    <tr>
                        {this.fields.map(field => <th>{field.title}</th>)}
                    </tr>
                </thead>
                <tbody>
                {this.d.map(row => <tr>
                    {this.fields.map(field => <td>{row[field.field]}</td>)}
                </tr>)}
                {!this.has ? <a-empty/> : null}
                </tbody>
            </table>
        </no-scroll>
    },
    data() {
        return {
            fields: [],
            d: [],
            cfg: TableConfig(),
            st: 0,
            stNumber: 0,
            sth: null
        }
    },
    watch: {
        config: {
            immediate: true,
            deep: true,
            handler() {
                this.parse()
                this.launchScroll()
            }
        },
        data: {
            deep: true,
            immediate: true,
            handler(val) {
                this.getData()
            }
        },
    },
    computed: {
        has() {
            return this.d.length !== 0
        }
    },
    methods: {
        launchScroll() {
            if(this.sth !== null) {
                clearInterval(this.sth)
            }

            this.sth = setInterval(() => {
                if(this.st >= 100) {
                    this.st = 0
                    return
                }

                this.st += this.cfg.scrollPrecentOnce
            }, this.cfg.scrollCycle * 1000)
        },
        onScrollTop(val) {
            this.stNumber = val
        },
        getData() {
            this.d = []

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
            this.d = t
            this.$nextTick(() => {
                this.$refs.scroll.reScrollTop()
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