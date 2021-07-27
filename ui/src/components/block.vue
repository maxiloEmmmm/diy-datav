<script lang="tsx">
import components from './types/component.js'
import util from 'pkg/util'
import {ViewBLockTypeCommonInputItemParse, ViewBlockTypeConfig} from 'type'
export default {
    name: 'view-block',
    render() {
        let Component = components[this.type]

        return !Component
            ? <div>unknown block type: {this.type}</div>
            : <Component config={this.cfg.type} data={this.data}/>
    },
    data() {
        return {
            cfg: ViewBlockTypeConfig(),
            data: []
        }
    },
    props: {
        type: {
            type: String,
            default: "",
        },
        config: {
            type: String,
            default() {
                return ""
            },
        }
    },
    watch: {
        config: {
            immediate: true,
            handler: 'transformTypeConfig'
        }
    },
    methods: {
        transformTypeConfig() {
            const cfg = JSON.parse(this.config)
            cfg.common.input = this.normalInput(cfg.common.input)
            this.cfg = cfg
            this.fetch()
        },
        normalInput(inputs) {
            if(!Array.isArray(inputs)) {
                return []
            }

            return inputs.map(ViewBLockTypeCommonInputItemParse)
        },
        fetch() {
            this.cfg.common.input.forEach((input, index) => {
                this.$store.commit('view/loadData', {
                    input,
                    refresh: this.cfg.common.refresh,
                    cb: (data) => {
                        this.data[index] = data
                    }
                })
            })
        }
    }
}
</script>