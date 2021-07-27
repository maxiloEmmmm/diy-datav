import util from 'pkg/util'

export default {
    props: {
        config: {
            type: String,
            default: ''
        },
        edit: {
            type: Boolean,
            default: false
        },
        style: Object,
        option: Object,
        disabled: Boolean,
        ref: String,
        size: String,
        value: String,
        change: Function,
        item: Object,
    },
    watch: {
        config: {
            immediate: true,
            handler: 'transformConfig'
        },
    },
    render() {
        return <tool-form
            ref="form"
            view={true}
            viewEdit={this.edit}
            onChange={this.onChange}
        />
    },
    data() {
        return {
            cfg: {}
        }
    },
    methods: {
        transformConfig() {
            console.log('impl `transformConfig` function')
        },
        onChange: util.debounce(function(data){
            try {
                this.cfg = data
                this.change && this.change(JSON.stringify(this.cfg))
            }catch(e) {
                console.log('config change json parse err', e, this.cfg)
            }
        }, 100)
    }
}