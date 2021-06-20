<script lang="tsx">
import move from './move.vue'
import edit from './edit.vue'
import {mapGetters} from "vuex"
import {Module as HelpModule} from '@/mixins/help'
import util from 'pkg/util'
export default {
    name: 'block-wrap',
    components: {
        edit, move
    },
    render(){
        let context = this.$slots.default()
        let moveAttrs = {
            class: ['ext-wrap', ...this.$attrs.class],
            style: {
                left: `${this.status.position.left}%`,
                top: `${this.status.position.top}%`
            }
        }

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
                    oldPosition: {
                        domLeft: 0,
                        domTop: 0,
                        left: 0,
                        top: 0
                    }
                },
                position: {
                    left: 0,
                    top: 0
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
        blockKey: {type: String}
    },
    created() {
        this.mixinInitFocus()
    },
    methods: {
        onMouseDown(e) {
            e.preventDefault()
            e.stopPropagation()
            this.mixinDoFocus()
            this.status.mouse.down = true
            this.status.mouse.oldPosition.left = e.clientX
            this.status.mouse.oldPosition.top = e.clientY
            this.status.mouse.oldPosition.domLeft = this.$refs.move.$el.offsetLeft
            this.status.mouse.oldPosition.domTop = this.$refs.move.$el.offsetTop

            let moveCb = util.throttle((e) => {
                this.$nextTick(() => {
                    let x = this.status.mouse.oldPosition.domLeft + e.clientX - this.status.mouse.oldPosition.left
                    let y = this.status.mouse.oldPosition.domTop + e.clientY - this.status.mouse.oldPosition.top
                    x = x < 0 ? 0 : (x > document.body.clientWidth ? document.body.clientWidth : x)
                    y = y < 0 ? 0 : (y > document.body.clientHeight ? document.body.clientHeight : y)

                    this.status.position.left = `${(x / document.body.clientWidth).toFixed(3) * 100}`
                    this.status.position.top = `${(y / document.body.clientHeight).toFixed(3) * 100}`
                    this.status.mouse.move = true
                })
            }, 25)

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
    .ext-help {
        position: absolute;
        left: 0; right: 0; top: 0; bottom: 0;
        display: flex; justify-content: center; align-items: center;
    }
}
</style>