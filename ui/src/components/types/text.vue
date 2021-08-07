<script lang="jsx">
import {TextConfig, TextConfigParse} from 'type'
import util from 'pkg/util'
export default {
    props: {
        config: {
            type: Object,
            default() {
                return TextConfig()
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
        const itemLength = this._text.length
        for (let i = 0; i < itemLength; i++) {
            items.push(<span>{this._text[i]}{i === itemLength - 1 ? [sup, sub] : []}</span>)
        }

        return <div class={c} style={s}>
            {items}
        </div>
    },
    data() {
        return {
            remoteData: '',
            cfg: TextConfig()
        }
    },
    computed: {
        _text() {
            return this.remoteData ? this.remoteData : this.cfg.text
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
                this.getData()
            }
        },
    },
    methods: {
        getData() {
            let data = []
            if(util.isArray(this.data)) {
                data = this.data[this.cfg.dataIndex]
            }

            let t = ""

            if(util.isObject(data)) {
                t = data.text || ''
            }

            if(util.isString(data)) {
                t = data
            }

            this.remoteData = t
        },
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