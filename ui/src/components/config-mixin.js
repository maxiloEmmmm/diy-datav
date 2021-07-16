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
    methods: {
        transformConfig() {
            console.log('impl `transformConfig` function')
        },
        onChange: util.debounce(function(){
            try {
                this.mixinSetConfigConfig(JSON.stringify(this.cfg))
            }catch(e) {
                console.log('config change json parse err', e, this.cfg)
            }
        }, 100)
    }
}