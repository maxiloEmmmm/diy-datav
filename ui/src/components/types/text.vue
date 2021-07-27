<script lang="jsx">
import {TextConfig, TextConfigParse} from 'type'
export default {
    props: {
        config: {
            type: Object,
            default() {
                return TextConfig()
            }
        }
    },
    render() {
        let c = {
            'st': true,
            'st-bold': this.cfg.bold,
            'st-italic': this.cfg.italic,
            'st-underline': this.cfg.underline
        }
        let s = {}

        let sup, sub = null
        if(this.cfg.sup) {
            sup = <sup>{this.cfg.sup}</sup>
        }

        if(this.cfg.sub) {
            sub = <sub>{this.cfg.sub}</sub>
        }

        s.alignItems = this.cfg.align.v
        s.justifyContent = this.cfg.align.h
        s.color = this.cfg.color
        s.fontSize = `${this.cfg.size}rem`

        let items = []
        const itemLength = this.cfg.text.length
        for (let i = 0; i < itemLength; i++) {
            items.push(<span>{this.cfg.text[i]}{i === itemLength - 1 ? [sup, sub] : []}</span>)
        }

        return <div class={c} style={s}>
            {items}
        </div>
    },
    data() {
        return {
            cfg: TextConfig()
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
                this.cfg = TextConfigParse(this.config)
            }catch(e) {
                console.log('static text config parse failed', e, this.config)
            }
        },
    }
}
</script>

<style lang="scss" scoped>
    .st {
        width: 100%; height: 100%; display: flex; flex-direction: row;

        &.st-bold {font-weight: 800}
        &.st-italic {font-style: italic}
        &.st-underline {text-decoration: underline;}
    }
</style>