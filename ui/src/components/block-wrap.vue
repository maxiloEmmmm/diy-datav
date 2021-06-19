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
            {this.helps.map(help => {
                let Component = help.component()
                return <Component onClick={help.cb}/>
            })}
        </div> : null

        return <move
            ref="move"
            {...moveAttrs}
            onMousedown={this.onMouseDown}
            onMousemove={this.onMouseMove}
            onMouseup={this.onMouseUp}
            onMouseenter={this.onMouseEnter}
            onMouseleave={this.onMouseLeave}
            enable={this.app_mixin.focus.in}
            moving={this.status.mouse.move}
        >
            <edit>
                {help}
                {context}
                {this.app_mixin.focus.in ? 'focus' : 'no-focus'}
            </edit>
        </move>
    },
    data() {
        return {
            status: {
                mouse: {
                    down: false,
                    move: false,
                    downPosition: {
                        left: null,
                        right: null
                    }
                }
            }
        }
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
            this.status.mouse.down = true

            // TODO: save position

            let moveCb = () => {
                // TODO: move block by current position
                this.status.mouse.move = true
            }

            let upCb = () => {
                this.status.mouse.down = false
                this.status.mouse.move = false
                document.removeEventListener('mouseup', upCb)
                document.removeEventListener('mousemove', moveCb)
            }

            document.addEventListener('mouseup', upCb)
            document.addEventListener('mousemove', moveCb)
        },
        onMouseMove(e) { },
        onMouseUp(e) {},
        onMouseEnter(e) {},
        onMouseLeave(e) {}
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