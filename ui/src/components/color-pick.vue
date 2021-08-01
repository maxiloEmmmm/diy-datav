<script lang="tsx">
// TODO: vue-color wait use official
import {Sketch} from '@ckpack/vue-color'
export default {
    name: "color-pick",
    components: {cp: Sketch},
    render(){
        return <a-popover v-slots={{
            content: () => <cp modelValue={this.color} onUpdate:modelValue={this.onInput}/>
        }} title="颜色选择" trigger="click" visible={this.visible} onVisibleChange={e => this.visible = e}>
            <a-button size="small" type="primary">
                {this.color ? <div style={{backgroundColor: this.color, borderRadius: "4px", width: "1rem", height: "1rem"}}/> : "选择"}
            </a-button>
        </a-popover>
    },
    props: {
        value: String
    },
    watch: {value(val, oldvar){
            this.color = this.init(val)
        }},
    data(){return {color: this.init(this.$props.value), visible: false}},
    emits: ["update:value", "change"],
    methods: {
        onInput(x){
            console.log('?')
            this.color = x.hex
            this.$emit("update:value", this.color)
            this.$emit("change", this.color)
            this.visible = false
        },
        init(color){
            return color ? color : "#fff"
        }
    }
}
</script>

<style lang="scss" scoped>
.icon-pick {
&:hover {
     border: 1px pink dashed; padding: 2px; border-radius: 6px;
 }
&.active {
     border: 1px pink dashed; background: pink; color: #fff;border-radius: 6px;padding: 4px;
 }
}
</style>