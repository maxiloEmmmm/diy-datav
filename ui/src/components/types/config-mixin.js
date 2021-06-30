export default {
    emits: ['change'],
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
        onChange() {
            try {
                this.$emit('config', JSON.stringify(this.cfg))
            }catch(e) {
                console.log('config change json parse err', e, this.cfg)
            }
        }
    }
}