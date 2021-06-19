<script lang="tsx">
import move from './move.vue'
import edit from './edit.vue'
import {mapGetters} from "vuex"
import {Module as HelpModule} from '@/mixins/help'
export default {
    name: 'block-wrap',
    components: {
        edit, move
    },
    render(){
        let context = this.$slots.default()
        let moveAttrs = {class: ['ext-wrap', ...this.$attrs.class]}

        let help = this.hasHelp ? <div class="ext-help">
            {this.helps.map(help => help.component())}
        </div> : null

        return <move
            {...moveAttrs}
            onMousedown={this.onMouseDown}
            enable={this.app_mixin.focus.in}
        >
            <edit>
                {help}
                {context}
                {this.app_mixin.focus.in ? 'focus' : 'no-focus'}
            </edit>
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
        blockKey: {type: Number}
    },
    created() {
        this.mixinInitFocus()
    },
    methods: {
        onMouseDown(e) {
            e.stopPropagation()
            this.mixinDoFocus()
        },
    }
}
</script>

<style lang="scss" scoped>
.ext-wrap {
    position: relative;
    .ext-help {
        position: absolute;
        left: 0; right: 0; top: 0; bottom: 0;
        display: flex; justify-content: center; align-items: center;
    }
}
</style>