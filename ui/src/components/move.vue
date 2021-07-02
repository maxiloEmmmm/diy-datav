<script lang="tsx">
import {Module as HelpModule} from '@/mixins/help'
import { DragOutlined } from '@ant-design/icons-vue';
import util from 'pkg/util'
import {
    UpCircleOutlined, DownCircleOutlined, LeftCircleOutlined,
    RightCircleOutlined, RadiusUpleftOutlined, RadiusBottomleftOutlined,
    RadiusBottomrightOutlined, RadiusUprightOutlined
} from '@ant-design/icons-vue'

export default {
    render() {
        let context = this.$slots.default()
        let attrs = {
            class: this.$attrs.class,
            style: {
                left: `${this.status.box.left}%`,
                top: `${this.status.box.top}%`,
                height: `${this.status.box.height}%`,
                width: `${this.status.box.width}%`,
                border: '1px red dashed'
            }
        }

        // bar view in body or current move?
        // move must e.stopPropagation

        let bar = this.enable ? <div style="position:absolute; inset: 0;">
            <RadiusBottomrightOutlined onMousedown={this.onBarDown} style="position:absolute; right: -2rem; bottom: -2rem; cursor: pointer"/>
        </div> : null

        let moveMiniBar = this.enable ? <div style="position:absolute; inset: 0;">
            <UpCircleOutlined onMousedown={e => e.stopPropagation()} onMouseup={() => this.onBarMove('up')} style="position:absolute; left: 50%; top: -2rem; cursor: pointer"/>
            <DownCircleOutlined onMousedown={e => e.stopPropagation()} onMouseup={() => this.onBarMove('down')} style="position:absolute; left: 50%; bottom: -2rem; cursor: pointer"/>
            <LeftCircleOutlined onMousedown={e => e.stopPropagation()} onMouseup={() => this.onBarMove('left')} style="position:absolute; left: -2rem; top: 50%; cursor: pointer"/>
            <RightCircleOutlined onMousedown={e => e.stopPropagation()} onMouseup={() => this.onBarMove('right')} style="position:absolute; right: -2rem; bottom: 50%; cursor: pointer"/>
        </div> : null

        return <div
            ref="move"
            {...attrs}

            onMousedown={this.onMouseDown}
            onMousemove={this.onMouseMove}
            onMouseup={this.onMouseUp}
            onMouseenter={this.onMouseEnter}
            onMouseleave={this.onMouseLeave}
        >
            {bar}
            {moveMiniBar}
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
                barMouse: {
                    down: false,
                    move: false,
                    icon: null
                },
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
                box: {
                    width: 10,
                    height: 10,
                    old: {
                        width: 10,
                        height: 10
                    }
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

            this.mixinAddHelp(HelpModule.ViewBlock, [
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
                let x = this.status.mouse.oldPosition.domLeft + e.clientX - this.status.mouse.oldPosition.left
                let y = this.status.mouse.oldPosition.domTop + e.clientY - this.status.mouse.oldPosition.top
                x = x < 0 ? 0 : (x > document.body.clientWidth ? document.body.clientWidth : x)
                y = y < 0 ? 0 : (y > document.body.clientHeight ? document.body.clientHeight : y)

                this.status.box.left = `${(x / document.body.clientWidth).toFixed(3) * 100}`
                this.status.box.top = `${(y / document.body.clientHeight).toFixed(3) * 100}`
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
        onMouseLeave(e) {},
        onBarDown(e) {
            this.status.mouse.oldPosition.left = e.clientX
            this.status.mouse.oldPosition.top = e.clientY
            this.status.mouse.oldPosition.domLeft = this.$refs.move.offsetLeft
            this.status.mouse.oldPosition.domTop = this.$refs.move.offsetTop
            this.status.box.old.width = this.$refs.move.clientWidth
            this.status.box.old.height = this.$refs.move.clientHeight

            const resizeCtrl = util.debounce(() => {
                const e = document.createEvent('Event')
                e.initEvent('resize', true, true)
                window.dispatchEvent(e)
            }, 200)

            e.preventDefault()
            e.stopPropagation()
            this.status.barMouse.down = true
            let moveCb = (e) => {
                let width = this.status.box.old.width + e.clientX - this.status.mouse.oldPosition.left
                let height = this.status.box.old.height + e.clientY - this.status.mouse.oldPosition.top
                let maxWidth = document.body.clientWidth - this.status.mouse.oldPosition.domLeft
                width = width < 0 ? 0 : (width > maxWidth ? maxWidth : width)
                let maxHeight = document.body.clientHeight - this.status.mouse.oldPosition.domTop
                height = height < 0 ? 0 : (height > maxHeight ? maxHeight : height)

                this.status.box.width = `${(width / document.body.clientWidth).toFixed(3) * 100}`
                this.status.box.height = `${(height / document.body.clientHeight).toFixed(3) * 100}`
                this.status.barMouse.move = true
                resizeCtrl()
            }

            let upCb = () => {
                this.status.barMouse.down = false
                this.status.barMouse.move = false
                document.removeEventListener('mouseup', upCb)
                document.removeEventListener('mousemove', moveCb)
            }

            document.addEventListener('mouseup', upCb)
            document.addEventListener('mousemove', moveCb)
        },
        onBarMove(typ) {
            let tmp
            let step = 0.1
            switch (typ) {
                case 'up':
                    tmp = parseFloat(this.status.box.top) - step
                    this.status.box.top = tmp < 0 ? 0 : tmp
                    break
                case 'down':
                    tmp = parseFloat(this.status.box.top) + step
                    this.status.box.top = tmp > 100 ? 100 : tmp
                    break
                case 'left':
                    tmp = parseFloat(this.status.box.left) - step
                    this.status.box.left = tmp < 0 ? 0 : tmp
                    break
                case 'right':
                    tmp = parseFloat(this.status.box.left) + step
                    this.status.box.left = tmp > 100 ? 100 : tmp
                    break
            }
        }
    }
}
</script>