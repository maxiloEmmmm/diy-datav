<script lang="tsx">
import components from './types/component.js'
import util from 'pkg/util'
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
            cfg: {
                common: {
                    input: []
                },
                type: ""
            },
            data: []
        }
    },
    props: {
        type: {
            type: String,
            default: "",
        },
        config: {
            type: [String, Object],
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
    created() {
        this.fetch()
    },
    methods: {
        transformTypeConfig() {
            // TODO: wait alert
            // TODO: parse config from props, get common and type config
            switch (typeof this.config) {
                case 'object':
                    this.mergeConfig(this.config)
                    break
                case 'string':
                    try {
                        this.mergeConfig(JSON.parse(this.config))
                    }catch(e) {
                        return
                    }
                    break
                default:
                    return
            }
        },
        mergeConfig(config) {
            let common = config.common
            this.cfg.common.input = this.normalInput(common.input)
            this.cfg.type = config.type || {}
        },
        normalInput(inputs) {
            if(!Array.isArray(inputs)) {
                return []
            }

            return inputs.filter(input => !!input.id)
        },
        fetch() {
            this.cfg.common.input.forEach((input, index) => {
                this.$store.commit('view/loadData', {
                    id: input.id,
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