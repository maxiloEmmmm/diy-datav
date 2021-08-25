<script lang="jsx">
import {GridConfig, GridConfigParse} from 'type'
export default {
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
    render() {
        // if block dragend and grid col is mark (when block is close to grid col, col will be mark)
        // resize and move block to grid col

        // if grid resize and move end
        // same resize and move block
        const editBorder = this.edit ? row => `1px ${row ? 'dashed' : 'solid'} red` : () => 'unset'
        return <div style={{display: 'flex', flexDirection: 'column', height: '100%', padding: `${this.cfg.padding.top}% ${this.cfg.padding.right}% ${this.cfg.padding.bottom}% ${this.cfg.padding.left}%`}}>
            {this.cfg.rows.map((row, ri) => <div style={{border: editBorder(true), display: 'flex', flexDirection: 'row', flex: `0 0 ${row.height}%`, padding: `${row.padding.top}% ${row.padding.right}% ${row.padding.bottom}% ${row.padding.left}%`}}>
                {row.rowCols.map((col, ci) => <div style={{border: editBorder(false), flex: `0 0 ${col.width}%`, padding: `${col.padding.top}% ${col.padding.right}% ${col.padding.bottom}% ${col.padding.left}%`}}
                    onMousemove={() => this.onColMousemove(ri, ci)}
                >
                </div>)}
            </div>)}
        </div>
    },
    data() {
        return {
            cfg: GridConfig()
        }
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
        parse() {
            try {
                this.cfg = GridConfigParse(this.config)
            }catch(e) {
                console.log('grid config parse failed', e, this.config)
            }
        },
        onColMousemove(ri, ci) {
            console.log(ri, ci)
        // TODO: record grid row/col info to view adsorption store
        }
    }
}
</script>