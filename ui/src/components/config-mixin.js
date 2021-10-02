import util from 'pkg/util'

export default {
    props: {
        config: {
            type: [String, Object],
            default: ''
        },
        handleChange: {
            type: Boolean,
            default: false
        }
    },
    watch: {
        config: {
            immediate: true,
            handler: 'transformConfig'
        }
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
        onChange: util.debounce(function(){
            if(this.handleChange) {
                this.$emit('change', this.cfg)
                return
            }

            try {
                this.mixinSetConfigConfig(JSON.stringify(this.cfg))
            }catch(e) {
                console.log('config change json parse err', e, this.cfg)
            }
        }, 100)
    }
}