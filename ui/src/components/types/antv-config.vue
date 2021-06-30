<script lang="tsx">
import {AntVConfig} from 'type'
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
            cfg: AntVConfig()
        }
    },
    render() {
        return <div>antv config</div>
    },
    methods: {
        transformConfig() {

        }
    }
}
</script>