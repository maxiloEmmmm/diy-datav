<script lang="tsx">
import move from './move.vue'
import {mapGetters, mapState} from "vuex"
import {Module as HelpModule} from '@/mixins/help'
import {ViewBlockTypeConfig, ViewBLockTypeCommonParse} from 'type'
import {provide, toRefs} from "vue";
import * as configComponentType from '@/components/types/type.js'

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
            onAdsorptionEnd={this.onAdsorptionEnd}
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
        ...mapState('view', {
            adsorptionGridBlockKey: state => state.adsorption.grid.blockKey,
        }),
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
        type: String,
        edit: {
            type: Boolean,
            default: true
        },
        pointerEventsNone: {
            type: Boolean,
            default: false
        },
        view: {type: Object}
    },
    created() {
        this.mixinInitFocus()
        // TODO: design model add ctrl+z / ctrl shift z
    },
    methods: {
        transformTypeConfig() {
            let cfg = JSON.parse(this.config)
            cfg.common = ViewBLockTypeCommonParse(this.type, cfg.common)
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
        },
        onAdsorptionEnd() {
            // grid1落到了grid2上 保证grid1.zIndex > grid2.zIndex 防止同级dom带来block无法降落到grid1的问题
            if(this.type === configComponentType.Grid) {
                let targetBlock = this.view.blocks.filter(block => this.adsorptionGridBlockKey === block.getKey())[0]
                if(targetBlock) {
                    try {
                        let cfg = JSON.parse(targetBlock.config)
                        if(cfg.common.zIndex >= this.cfg.common.zIndex) {
                            this.cfg.common.zIndex = cfg.common.zIndex + 1
                            this.mixinSetConfigKey(this.blockKey)
                            this.mixinSetConfigConfig(JSON.stringify(this.cfg))
                        }
                    }catch(e) {
                        console.log('update adsorption grid index failed', e)
                    }
                }
            }
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