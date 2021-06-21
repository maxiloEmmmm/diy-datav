<script lang="tsx">
import {Module as HelpModule} from '@/mixins/help'
import { DragOutlined } from '@ant-design/icons-vue';
import util from 'pkg/util'
import {UpCircleOutlined, DownCircleOutlined, LeftCircleOutlined, RightCircleOutlined} from '@ant-design/icons-vue'

export default {
    render() {
        let context = this.$slots.default()
        let attrs = {
            class: this.$attrs.class,
            style: {
                left: `${this.status.position.left}%`,
                top: `${this.status.position.top}%`,
                border: this.enable ? '1px red dashed' : 'unset'
            }
        }

        // bar view in body or current move?
        // move must e.stopPropagation
        let bar = this.enable ? <div style="position:absolute; inset: 0;">
            <UpCircleOutlined onMousedown={e => e.stopPropagation()} onClick={() => this.onBarMove('up')} style="position:absolute; left: 50%; top: -2rem; cursor: pointer"/>
            <DownCircleOutlined onMousedown={e => e.stopPropagation()} onClick={() => this.onBarMove('down')} style="position:absolute; left: 50%; bottom: -2rem; cursor: pointer"/>
            <LeftCircleOutlined onMousedown={e => e.stopPropagation()} onClick={() => this.onBarMove('left')} style="position:absolute; left: -2rem; top: 50%; cursor: pointer"/>
            <RightCircleOutlined onMousedown={e => e.stopPropagation()} onClick={() => this.onBarMove('right')} style="position:absolute; right: -2rem; bottom: 50%; cursor: pointer"/>
        </div> : null

        return <div
            ref="move"
            {...attrs}
            style="width:120px; height:120px"

            onMousedown={this.onMouseDown}
            onMousemove={this.onMouseMove}
            onMouseup={this.onMouseUp}
            onMouseenter={this.onMouseEnter}
            onMouseleave={this.onMouseLeave}
        >
            {bar}
            {context}
        </div>
    },
    props: {
        enable: {
            type: Boolean, default: false
        },
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
    watch: {
        enable(val) {
            (val ? this.mixinActiveHelp : this.mixinDisableHelp)(HelpModule.ViewBlock, "move")
        }
    },
    emits: ['mousedown'],
    methods: {
        onMouseDown(e) {
            e.preventDefault()
            e.stopPropagation()

            this.mixinSetHelp(HelpModule.ViewBlock, [
                {key: "move", component() {
                        return <DragOutlined twoToneColor="red"/>
                    }, cb: () => {
                        console.log("move help click")
                    }, disable: !this.enable}
            ])

            this.$emit('mousedown', e)

            this.status.mouse.down = true
            this.status.mouse.oldPosition.left = e.clientX
            this.status.mouse.oldPosition.top = e.clientY
            this.status.mouse.oldPosition.domLeft = this.$refs.move.offsetLeft
            this.status.mouse.oldPosition.domTop = this.$refs.move.offsetTop

            let moveCb = (e) => {
                this.$nextTick(() => {
                    let x = this.status.mouse.oldPosition.domLeft + e.clientX - this.status.mouse.oldPosition.left
                    let y = this.status.mouse.oldPosition.domTop + e.clientY - this.status.mouse.oldPosition.top
                    x = x < 0 ? 0 : (x > document.body.clientWidth ? document.body.clientWidth : x)
                    y = y < 0 ? 0 : (y > document.body.clientHeight ? document.body.clientHeight : y)

                    this.status.position.left = `${(x / document.body.clientWidth).toFixed(3) * 100}`
                    this.status.position.top = `${(y / document.body.clientHeight).toFixed(3) * 100}`
                    this.status.mouse.move = true
                })
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
        onMouseLeave(e) {},
        onBarMove(typ) {
            let tmp
            let step = 0.1
            switch (typ) {
                case 'up':
                    tmp = parseFloat(this.status.position.top) - step
                    this.status.position.top = tmp < 0 ? 0 : tmp
                    break
                case 'down':
                    tmp = parseFloat(this.status.position.top) + step
                    this.status.position.top = tmp > 100 ? 100 : tmp
                    break
                case 'left':
                    tmp = parseFloat(this.status.position.left) - step
                    this.status.position.left = tmp < 0 ? 0 : tmp
                    break
                case 'right':
                    tmp = parseFloat(this.status.position.left) + step
                    this.status.position.left = tmp > 100 ? 100 : tmp
                    break
            }
        }
    }
}
</script>