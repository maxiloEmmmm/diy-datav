import util from 'pkg/util'

export default {
    props: {
        config: {
            type: String,
            default: ''
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
    emits: ['update:config', 'change'],
    methods: {
        transformConfig() {
            console.log('impl `transformConfig` function')
        },
        onChange: util.debounce(function(){
            try {
                const c = JSON.stringify(this.cfg)
                this.$emit('update:config', c)
                this.$emit('change', c)
            }catch(e) {
                console.log('config change json parse err', e, this.cfg)
            }
        }, 100)
    }
}