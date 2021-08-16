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
    },
    render() {
        return <div style={{display: 'flex', flexDirection: 'column', padding: `${this.cfg.padding.top}% ${this.cfg.padding.right}% ${this.cfg.padding.bottom}% ${this.cfg.padding.left}%`}}>
            {this.cfg.rows.map(row => <div style={{display: 'flex', flexDirection: 'row', flex: `0 0 ${row.height}%`, padding: `${row.padding.top}% ${row.padding.right}% ${row.padding.bottom}% ${row.padding.left}%`}}>
                {row.rowCols.map(col => <div style={{flex: `0 0 ${col.width}%`, padding: `${row.padding.top}% ${row.padding.right}% ${row.padding.bottom}% ${row.padding.left}%`}}>
                    ???
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
    }
}
</script>