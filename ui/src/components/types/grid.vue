<script lang="jsx">
import {GridConfig, GridConfigParse} from 'type'
import {mapState} from 'vuex'
import componentCommon from './component-common'
import * as resource from 'type/resource.js'
import { BlockPositions } from '../../type/resource'
export default {
    mixins: [componentCommon],
    props: {
        config: {
            type: Object,
            default() {
                return GridConfig()
            }
        },
        data: {
            type: Array,
            default() {
                return []
            }
        },
        edit: {
            type: Boolean,
            default: false
        }
    },
    created() {
        this.$sub.addHandle(resource.BlockMouseUp, this.onBlockMouseUp)
    },
    render() {
        // if block dragend and grid col is mark (when block is close to grid col, col will be mark)
        // resize and move block to grid col

        // if grid resize and move end
        // same resize and move block
        const editBorder = this.edit ? row => `1px ${row ? 'dashed' : 'solid'} red` : () => 'unset'
        return <div style={{display: 'flex', flexDirection: 'column', height: '100%', padding: `${this.cfg.padding.top}% ${this.cfg.padding.right}% ${this.cfg.padding.bottom}% ${this.cfg.padding.left}%`}}>
            {this.cfg.rows.map((row, ri) => <div style={{border: editBorder(true), display: 'flex', flexDirection: 'row', flex: `0 0 ${row.height}%`, padding: `${row.padding.top}% ${row.padding.right}% ${row.padding.bottom}% ${row.padding.left}%`}}>
                {row.rowCols.map((col, ci) => <div ref={`${this.makeRCKey(ri, ci)}`} style={{
                    border: editBorder(false),
                    flex: `0 0 ${col.width}%`,
                    padding: `${col.padding.top}% ${col.padding.right}% ${col.padding.bottom}% ${col.padding.left}%`,
                    borderColor: this.edit ? (this.shouldHighlightCol(ri, ci) ? 'red' : '#000' ) : 'unset'
                }}
                    onMousemove={() => this.onColMousemove(ri, ci)}
                    onMouseleave={() => this.onColMouseleave(ri, ci)}
                >
                </div>)}
            </div>)}
        </div>
    },
    data() {
        return {
            cfg: GridConfig(),
            moveFocusCol: ''
        }
    },
    computed: {
        ...mapState('view', ['blockMoving']),
    },
    watch: {
        config: {
            immediate: true,
            deep: true,
            handler() {
                this.parse()
            }
        },
        data: {
            deep: true,
            immediate: true,
            handler(val) {
            }
        },
    },
    methods: {
        onBlockMouseUp(payload) {
            if(this.blockKey !== payload.blockKey) {
                return
            }
            const req = []
            this.cfg.rows.forEach((r, ri) => {
                r.rowCols.forEach((c, ci) => {
                    if(c.keys.length !== 0) {
                        c.keys.forEach(key => {
                            req.push({
                                blockKey: key,
                                rect: this.computedRCRect(ri, ci)
                            })
                        })
                    }
                })
            })
            this.$sub.dispatch(BlockPositions, req)
        },
        shouldHighlightCol(r, c) {
            return this.blockMoving !== "" && this.makeRCKey(r, c) === this.moveFocusCol
        },
        parse() {
            try {
                this.cfg = GridConfigParse(this.config)
            }catch(e) {
                console.log('grid config parse failed', e, this.config)
            }
        },
        makeRCKey(r, c) {
            return `r${r}c${c}`
        },
        onColMouseleave(r, c) {
            if(!this.blockMoving) {
                return
            }
            this.$store.commit("view/setAdsorptionGrid", {blockKey: ""})
        },
        computedRCPosition(r, c) {
            let obj = this.$refs[`${this.makeRCKey(r, c)}`]
            let l = 0, t = 0
            while(obj) {
                l += obj.offsetLeft
                t += obj.offsetTop
                obj = obj.offsetParent
            }
            return {top: t, left: l}
        },
        computedRCRect(r, c) {
            const {top, left} = this.computedRCPosition(r, c)
            return {
                left: parseFloat((left / document.body.clientWidth).toFixed(3)) * 100,
                top: parseFloat((top / document.body.clientHeight).toFixed(3)) * 100,
                width: parseFloat((this.$refs[`${this.makeRCKey(r, c)}`].clientWidth / document.body.clientWidth).toFixed(3)) * 100,
                height: parseFloat((this.$refs[`${this.makeRCKey(r, c)}`].clientHeight / document.body.clientHeight).toFixed(3)) * 100,
            }
        },
        onColMousemove(r, c) {
            if(!this.blockMoving) {
                return
            }
            this.moveFocusCol = this.makeRCKey(r, c)
            this.$store.commit("view/setAdsorptionGrid", {
                blockKey: this.blockKey, row: r, col: c, meta: {
                    r, c
                }, pos: this.computedRCRect(r, c)
            })
        }
    }
}
</script>