<script lang="tsx">
import components from './types/component.js'
import util from 'pkg/util'
import {toRefs, provide} from 'vue'
import textComponent from './types/text.vue'
import {ViewBLockTypeCommonInputItemParse, ViewBlockTypeConfig, ViewBLockTypeCommonParse} from 'type'
export default {
    name: 'view-block',
    components: {
        innerTextComponent: textComponent
    },
    render() {
        let Component = components[this.type]
        let hasDesc = !!this.cfg.common.desc.textConfig.text
        let isTop = this.cfg.common.desc.positionType

        let desc = <div style={{
            flexBasis: `${this.cfg.common.desc.position.height}%`,
            flexGrow: 0,
            flexShrink: 0
        }}>
            <inner-text-component config={this.cfg.common.desc.textConfig}/>
        </div>

        return !Component
            ? <div style={{pointerEvents: this.pointerEventsNone.value ? 'none' : 'all'}}>unknown block type: {this.type}</div>
            : <div style="display:flex; flex-direction: column; height: 100%;" style={{
                backgroundColor: !!this.cfg.common.bg ? this.cfg.common.bg : 'transparent',
                border: `${this.cfg.common.border.width}rem ${this.cfg.common.border.color} ${this.cfg.common.border.style}`
            }}>
                {hasDesc && isTop ? desc : null}
                <Component style={{flexGrow: 1}} common={this.cfg.common} config={this.cfg.type} data={this.data} edit={this.edit}/>
                {hasDesc && !isTop ? desc : null}
            </div>
    },
    data() {
        return {
            cfg: ViewBlockTypeConfig(),
            data: []
        }
    },
    inject: ['pointerEventsNone'],
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
        },
        edit: {
            type: Boolean,
            default: false
        },
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
            cfg.common = ViewBLockTypeCommonParse(this.type, cfg.common)
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