<script lang="tsx">
export default {
    name: 'select-item',
    props: {
        icon: {
            type: Function,
            default() {
                return () => null
            }
        },
        active: {
            type: Boolean,
            default: false
        }
    },
    render() {
        const label = this.$slots.default()
        const attrs = {
            class: ['type-item-wrap']
        }

        if(this.active) {
            attrs.class.push('active')
        }

        return <div {...attrs}>
            <div style="flex: 0 0 60%; text-align:center">{this.icon()}</div>
            <div style="flex: 0 0 40%; text-align:center">{label}</div>
        </div>
    }
}
</script>

<style lang="scss" scoped>
.type-item-wrap {
    display: flex; flex-direction: column; height: 100%; cursor: pointer;
    &.active {
         border: 1px dashed pink; border-radius: 6px; padding: 2px
    }
}
</style>