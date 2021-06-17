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
        let moveAttrs = {class: this.$attrs.class}

        return <move
            {...moveAttrs}
            onMousedown={this.onMouseDown}
        >
            <edit>
                {context}
                {this.app_mixin.focus.in ? 'focus' : 'no-focus'}
            </edit>
        </move>
    },
    data() {
        return {}
    },
    props: {
        blockKey: {type: Number}
    },
    created() {
        this._initFocus()
    },
    methods: {
        onMouseDown(e) {
            e.stopPropagation()
            this._doFocus()
        },
    }
}
</script>