<script lang="tsx">
import {Module as HelpModule} from '@/mixins/help'
import { DragOutlined } from '@ant-design/icons-vue';

export default {
    render() {
        let context = this.$slots.default()
        let attrs = {class: this.$attrs.class}

        return <div
            {...attrs}
            style="width:120px; height:120px"
        >
            {context}
        </div>
    },
    props: {
        enable: {
            type: Boolean, default: false
        }
    },
    data() {
        return {
        }
    },
    watch: {
        enable: {
            immediate: true,
            handler(val) {
                (val ? this.mixinActiveHelp : this.mixinDisableHelp)(HelpModule.ViewBlock, "move")
            }
        }
    },
    // TODO: created is after watch immediate, so use beforeCreate to add help
    // but can't use mixin addHelp func, because in beforeCreate mixins is not complete
    // wait resolve to use mixin addHelp func
    beforeCreate(){
        this.$store.commit('view/addHelp', {
            typ: HelpModule.ViewBlock,
            helps: [
                {key: "move", component() {
                    return <DragOutlined/>
                }, cb: () => {
                    console.log("move help click")
                }}
            ]
        })
    }
}
</script>