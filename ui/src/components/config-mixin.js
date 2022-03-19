import util from 'pkg/util'

export default {
    props: {
        config: {
            type: Object,
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
    methods: {
        transformConfig() {
            this.cfg = this.$util.merge({}, this.config)
        },
        onChange: util.debounce(function(){
            if(this.handleChange) {
                this.$emit('change', this.cfg)
                return
            }

            try {
                this.mixinSetConfigConfig(this.cfg)
            }catch(e) {
                console.log('config change json parse err', e, this.cfg)
            }
        }, 100)
    }
}