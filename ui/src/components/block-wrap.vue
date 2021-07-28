<script lang="tsx">
import move from './move.vue'
import {mapGetters} from "vuex"
import {Module as HelpModule} from '@/mixins/help'
import {ViewBlockTypeConfig} from 'type'
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
        }

        let help = this.edit && this.hasHelp && this.app_mixin.focus.in ? <div class="ext-help">
            {this.helps.map(help => {
                const Component = help.component()
                return <Component
                    onMousedown={e => this.onHelpComponentClick(e, help, 'mousedown')}
                    onClick={e => this.onHelpComponentClick(e, help, 'click')}/>
            })}
        </div> : null

        return <move
            {...moveAttrs}
            onMousedown={this.onMouseDown}
            onPosition={this.onPosition}
            enable={this.edit && this.app_mixin.focus.in}
            position={this.cfg.common.position}
        >
            {help}
            <div class="content">{context}</div>
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
            return this.appHelp[HelpModule.ViewBlock] || []
        }
    },
    watch: {
        config: {
            immediate: true,
            handler: 'transformTypeConfig'
        }
    },
    props: {
        blockKey: {type: String},
        config: String,
        edit: {
            type: Boolean,
            default: true
        }
    },
    created() {
        this.mixinInitFocus()
        // TODO: design model add ctrl+z / ctrl shift z
    },
    methods: {
        transformTypeConfig() {
            this.cfg = JSON.parse(this.config)
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
}
</style>