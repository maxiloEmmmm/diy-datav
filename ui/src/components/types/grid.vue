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
        return <div/>
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