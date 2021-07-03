<script lang="tsx">
import move from './move.vue'
import {mapGetters} from "vuex"
import {Module as HelpModule} from '@/mixins/help'
export default {
    name: 'block-wrap',
    components: {
        move
    },
    emits: ['mousedown'],
    render(){
        let context = this.$slots.default()
        let moveAttrs = {
            class: ['ext-wrap', ...this.$attrs.class.split(' ')],
        }

        let help = this.hasHelp ? <div class="ext-help">
            {this.helps.map(help => {
                const Component = help.component()
                return <Component
                    onMousedown={e => e.stopPropagation()}
                    onClick={e => this.onHelpComponentClick(e, help)}/>
            })}
        </div> : null

        return <move
            {...moveAttrs}
            onMousedown={this.onMouseDown}
            enable={this.app_mixin.focus.in}
        >
            {help}
            <div class="content">{context}</div>
            {this.app_mixin.focus.in ? 'focus' : 'no-focus'}
        </move>
    },
    data() {
        return {}
    },
    computed: {
        ...mapGetters('view', {
            appHelp: 'help'
        }),
        hasHelp() {
            return this.helps.length > 0
        },
        helps() {
            return this.appHelp[HelpModule.ViewBlock] || []
        }
    },
    props: {
        blockKey: {type: String}
    },
    created() {
        this.mixinInitFocus()
    },
    methods: {
        onMouseDown(e) {
            this.mixinDoFocus()
            this.$emit('mousedown', e)
        },
        onHelpComponentClick(e, help) {
            e.stopPropagation()
            help.cb({
                e,
                blockKey: this.blockKey
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