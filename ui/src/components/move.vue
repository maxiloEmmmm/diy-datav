<script lang="tsx">
import {Module as HelpModule} from '@/mixins/help'
import util from 'pkg/util'
import {PositionType} from 'type'
import {mapState} from "vuex";
import {
    UpCircleOutlined, DownCircleOutlined, LeftCircleOutlined,
    RightCircleOutlined,
    RadiusBottomrightOutlined,
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
                border: this.edit ? '1px red dashed' : 'unset',
            }
        }

        // bar view in body or current move?
        // move must e.stopPropagation

        let moveMiniBar = this.edit && this.enable ? <div style="position:absolute; inset: 0;" style={{pointerEvents: this.pointerEventsNone ? 'none' : 'all'}}>
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
            {moveMiniBar}
            {context}
        </div>
    },
    props: {
        blockKey: {type: String},
        enable: {
            type: Boolean, default: false
        },
        edit: {
            type: Boolean, default: false
        },
        // don't need to watch position change
        // just use once on init
        position: {
            default: () => PositionType(),
            type: Object
        }
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
                    left: this.$props.position.left,
                    top: this.$props.position.top,
                    width: this.$props.position.width,
                    height: this.$props.position.height,
                    old: {
                        width: 10,
                        height: 10
                    }
                }
            }
        }
    },
    watch: {
        'status.box': {
            deep: true,
            handler: util.debounce(function() {
                this.$emit('position', {
                    left: this.status.box.left,
                    top: this.status.box.top,
                    width: this.status.box.width,
                    height: this.status.box.height
                })
            }, 100)
        }
    },
    created() {
        this.mixinAddHelp(HelpModule.ViewBlock, [
            {key: `resize-${this.blockKey}`, component() {
                    return <RadiusBottomrightOutlined twoToneColor="red"/>
                }, cb: (payload) => {
                    if(payload.type !== 'mousedown') {
                        return
                    }

                    this.onBarDown(payload.e)
                },
            }
        ])
    },
    emits: ['mousedown', 'position'],
    inject: ['pointerEventsNone'],
    computed: {
        ...mapState('view', ['adsorption']),
        hasAdsorptionDesign() {
            return this.adsorption.design.lineIndex >= 0
        },
        hasAdsorptionGrid() {
            return this.adsorption.grid.blockKey !== ""
        }
    },
    methods: {
        onMouseDown(e) {
            e.preventDefault()
            e.stopPropagation()

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

                this.status.box.left = parseFloat((x / document.body.clientWidth).toFixed(3)) * 100
                this.status.box.top = parseFloat((y / document.body.clientHeight).toFixed(3)) * 100
                this.status.mouse.move = true
                this.$store.commit("view/setBlockMoving", this.blockKey)
            }

            let upCb = () => {
                this.status.mouse.down = false
                this.status.mouse.move = false
                this.$store.commit("view/setBlockMoving", "")
                document.removeEventListener('mouseup', upCb)
                document.removeEventListener('mousemove', moveCb)
                if (this.hasAdsorptionDesign || this.hasAdsorptionGrid) {
                    if(this.hasAdsorptionGrid) {
                        this.status.box.left = this.adsorption.grid.pos.left
                        this.status.box.top = this.adsorption.grid.pos.top
                        this.status.box.width = this.adsorption.grid.pos.width
                        this.status.box.height = this.adsorption.grid.pos.height
                        this.$store.commit('view/setAdsorptionGrid', {
                            blockKey: ""
                        })
                    }else {
                        this.status.box.left = this.adsorption.design.pos.left
                        this.status.box.top = this.adsorption.design.pos.top
                        this.$store.commit('view/setAdsorptionDesign', {
                            lineIndex: -1
                        })
                    }
                }
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

                this.status.box.width = parseFloat((width / document.body.clientWidth).toFixed(3)) * 100
                this.status.box.height = parseFloat((height / document.body.clientHeight).toFixed(3)) * 100
                this.mixinDispatchWindowResize()
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