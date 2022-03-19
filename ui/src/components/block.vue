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
        let hasDesc = !!this.config.common.desc.textConfig.text
        let isTop = this.config.common.desc.positionType

        let desc = <div style={{
            flexBasis: `${this.config.common.desc.position.height}%`,
            flexGrow: 0,
            flexShrink: 0
        }}>
            <inner-text-component config={this.config.common.desc.textConfig}/>
        </div>

        return !Component
            ? <div style={{pointerEvents: this.pointerEventsNone.value ? 'none' : 'all'}}>unknown block type: {this.type}</div>
            : <div style="display:flex; flex-direction: column; height: 100%;" style={{
                backgroundColor: !!this.config.common.bg ? this.config.common.bg : 'transparent',
                border: `${this.config.common.border.width}rem ${this.config.common.border.color} ${this.config.common.border.style}`
            }}>
                {hasDesc && isTop ? desc : null}
                <Component ref="component"  style={{flexGrow: 1}} common={this.config.common} config={this.config.type} data={this.data} edit={this.edit}/>
                {hasDesc && !isTop ? desc : null}
            </div>
    },
    data() {
        return {
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
            type: Object
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
            this.fetch()
        },
        normalInput(inputs) {
            if(!Array.isArray(inputs)) {
                return []
            }

            return inputs.map(ViewBLockTypeCommonInputItemParse)
        },
        fetch() {
            this.config.common.input.forEach((input, index) => {
                this.$store.commit('view/loadData', {
                    input,
                    refresh: this.config.common.refresh,
                    cb: (data) => {
                        this.data[index] = data
                    }
                })
            })
        }
    }
}
</script>