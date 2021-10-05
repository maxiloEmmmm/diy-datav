<script lang="tsx">
import move from './move.vue'
import {mapGetters} from "vuex"
import {Module as HelpModule} from '@/mixins/help'
import {ViewBlockTypeConfig, ViewBLockTypeCommonParse} from 'type'
import {provide, toRefs} from "vue";
export default {
    name: 'block-wrap',
    components: {
        move
    },
    emits: ['mousedown', 'config'],
    render(){
        let context = this.$slots.default()
        let moveAttrs = {
            class: ['ext-wrap', ...this.$attrs.class.split(' ')],
            style: {
                zIndex: this.cfg.common.zIndex,
                // TODO: parent style will merge `move` style
                pointerEvents: this.pointerEventsNone ? 'none' : 'all',
            },
            blockKey: this.blockKey
        }

        let help = this.edit && this.hasHelp && this.app_mixin.focus.in ? <div class="ext-help" style={{pointerEvents: this.pointerEventsNone ? 'none' : 'all'}}>
            <div style="background-color:#fff;border-radius: 6px; padding: 0 0.3rem">
                {this.helps.map(help => {
                    const Component = help.component()
                    return <Component
                        onMousedown={e => this.onHelpComponentClick(e, help, 'mousedown')}
                        onClick={e => this.onHelpComponentClick(e, help, 'click')}/>
                })}
            </div>
        </div> : null

        let extIndex = this.edit ? <span class="ext-index" style={{pointerEvents: this.pointerEventsNone ? 'none' : 'all'}}>层叠位置: {this.cfg.common.zIndex}</span> : null

        return <move
            {...moveAttrs}
            onMousedown={this.onMouseDown}
            onPosition={this.onPosition}
            enable={this.app_mixin.focus.in}
            edit={this.edit}
            position={this.cfg.common.position}
        >
            {extIndex}
            {help}
            <div class="content" style={{pointerEvents: this.pointerEventsNone ? 'none' : 'all'}}>{context}</div>
        </move>
    },
    data() {
        return {cfg: ViewBlockTypeConfig()}
    },
    computed: {
        ...mapGetters('view', {
            appHelp: 'help',
        }),
        hasHelp() {
            return this.helps.length > 0
        },
        helps() {
            return this.appHelp[HelpModule.ViewBlock].filter(help => !help.key.startsWith("resize-") || help.key === `resize-${this.blockKey}`) || []
        },
    },
    watch: {
        config: {
            immediate: true,
            handler: 'transformTypeConfig'
        }
    },
    setup(props) {
        let {pointerEventsNone} = toRefs(props)
        provide('pointerEventsNone', pointerEventsNone)
        provide('blockKey', props.blockKey)
    },
    props: {
        blockKey: {type: String},
        config: String,
        edit: {
            type: Boolean,
            default: true
        },
        pointerEventsNone: {
            type: Boolean,
            default: false
        }
    },
    created() {
        this.mixinInitFocus()
        // TODO: design model add ctrl+z / ctrl shift z
    },
    methods: {
        transformTypeConfig() {
            let cfg = JSON.parse(this.config)
            cfg.common = ViewBLockTypeCommonParse(cfg.common)
            this.cfg = cfg
        },
        onPosition(position) {
            try {
                const config = JSON.parse(this.config)
                config.common.position = position
                this.$emit('config', JSON.stringify(config))
            }catch (e) {
                console.log('block-wrap onPosition error', e)
            }
        },
        onMouseDown(e) {
            if(!this.edit) {
                return
            }

            this.mixinDoFocus()
            this.$emit('mousedown', e)
        },
        onHelpComponentClick(e, help, typ) {
            e.stopPropagation()
            help.cb({
                e,
                blockKey: this.blockKey,
                type: typ,
            })
        }
    }
}
</script>

<style lang="scss" scoped>
.ext-wrap {
    .ext-help {
        position: absolute; z-index: 3;
        left: 0; right: 0; top: 0; bottom: 0;
        display: flex; justify-content: center; align-items: center;
    }
    .content {
        position: absolute; z-index: 2;
        left: 0; right: 0; top: 0; bottom: 0;
    }
    .ext-index {
        position: absolute; z-index: 3;
        left: 0.5rem; top: -1.5rem; color: pink
    }
}
</style>