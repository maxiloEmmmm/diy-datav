<script type="tsx">

import antvType from './types/antv'

export default {
    components: {
        antvType,
    },
    render() {
        let Component = {
            antv: antvType
        }[this.type]
        return <Component config={this.cfg.type} data={this.data}/>
    },
    data() {
        return {
            cfg: {
                common: {},
                type: ""
            },
            data: {}
        }
    },
    props: {
        type: {
            type: String,
            default: "",
        },
        id: {
            type: Number,
            default: "",
        },
        config: {
            type: String,
            default: "",
        }
    },
    watch: {
        config: {
            immediate: true,
            handler: 'transformTypeConfig'
        }
    },
    methods: {
        transformTypeConfig() {
        // TODO: parse config from props, get common and type config
        // TODO: use common.input config to fetch data
        }
    }
}
</script>