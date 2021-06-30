<script lang="tsx">
import move from './move.vue'
import {mapGetters} from "vuex"
import {Module as HelpModule} from '@/mixins/help'
export default {
    name: 'block-wrap',
    components: {
        move
    },
    render(){
        let context = this.$slots.default()
        let moveAttrs = {
            class: ['ext-wrap', ...this.$attrs.class.split(' ')],
        }

        let help = this.hasHelp ? <div class="ext-help">
            {this.helps.map(help => {
                const Component = help.component()
                return <Component onClick={e => help.cb({
                    e,
                    blockKey: this.blockKey
                })}/>
            })}
        </div> : null

        return <move
            {...moveAttrs}
            onMousedown={this.onMouseDown}
            enable={this.app_mixin.focus.in}
        >
            {help}
            {context}
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
        onMouseDown() {
            this.mixinDoFocus()
        }
    }
}
</script>

<style lang="scss" scoped>
.ext-wrap {
    .ext-help {
        position: absolute;
        left: 0; right: 0; top: 0; bottom: 0;
        display: flex; justify-content: center; align-items: center;
    }
}
</style>