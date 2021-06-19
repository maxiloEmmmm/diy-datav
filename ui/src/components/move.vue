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
        },
        moving: {
            type: Boolean, default: false
        }
    },
    data() {
        return {
        }
    },
    watch: {
        enable(val) {
            (val ? this.mixinActiveHelp : this.mixinDisableHelp)(HelpModule.ViewBlock, "move")
        }
    },
    created(){
        this.mixinAddHelp(HelpModule.ViewBlock, [
            {key: "move", component() {
                return <DragOutlined/>
            }, cb: () => {
                console.log("move help click")
            }, disable: !this.enable}
        ])
    }
}
</script>