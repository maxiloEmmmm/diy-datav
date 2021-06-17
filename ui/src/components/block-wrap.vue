<script lang="tsx">
import move from './move.vue'
import edit from './edit.vue'

export default {
    name: 'block-wrap',
    components: {
        edit, move
    },
    render(){
        let context = this.$slots.default()
        let attrs = {}

        return <move>
            <edit>
                <div
                    {...attrs}>
                    {context}
                </div>
            </edit>
        </move>
    },
    data() {
        return {}
    },
    props: {
        blockKey: {type: Number}
    },
    methods: {}
}
</script>