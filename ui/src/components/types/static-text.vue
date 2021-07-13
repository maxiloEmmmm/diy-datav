<script lang="jsx">
import {StaticTextConfig, StaticTextConfigParse} from 'type'
export default {
    props: {
        config: {
            type: Object,
            default() {
                return StaticTextConfig()
            }
        }
    },
    render() {
        let c = {
            'st': true,
            'st-center-v': this.cfg.center.v,
            'st-center-h': this.cfg.center.h,
            'st-bold': this.cfg.bold,
            'st-italic': this.cfg.italic
        }

        return <div class={c}>{this.cfg.text}</div>
    },
    data() {
        return {
            cfg: StaticTextConfig()
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
                this.cfg = StaticTextConfigParse(this.config)
            }catch(e) {
                console.log('static text config parse failed', e, this.config)
            }
        },
    }
}
</script>

<style lang="scss" scoped>
    .st {
        width: 100%; height: 100%; display: flex; flex-direction: column;

        &.st-center-v {justify-content: center}
        &.st-center-h {align-items: center}
        &.st-bold {font-weight: 800}
        &.st-italic {font-style: italic}
    }
</style>