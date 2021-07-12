<script lang="tsx">
import {StaticText, StaticTextParse} from 'type'
export default {
    props: {
        config: {
            type: Object,
            default() {
                return StaticText()
            }
        }
    },
    render() {
        return <span>{this.cfg.text}</span>
    },
    data() {
        return {
            cfg: StaticText()
        }
    },
    watch: {
        config: {
            immediate: true,
            deep: true,
            handler() {
                this.parse()
            }
        }
    },
    methods: {
        parse() {
            try {
                this.cfg = StaticTextParse(this.config)
            }catch(e) {
                console.log('static text config parse failed', e, this.config)
            }
        },
    }
}
</script>