<script lang="jsx">
export default {
    name: 'block-wrap',
    render(){
        let context = this.$slots.default()

        let attrs = {}
        attrs.draggable = true
        attrs.class = ['wrap']

        if(this.status.draging) {
            attrs.class.push('bw-draging')
        }

        return <div 
            {...this.$attrs} 
            {...attrs}

            onDragstart={this.onDrag}
            onDragend={this.onDragEnd}>
            {context}
        </div>
    },
    data() {
        return {
            status: {
                draging: false
            }
        }
    },
    props: {
        blockKey: {type: Number}
    },
    methods: {
        onDrag(de) {
            this.status.draging = true
        },
        onDragEnd(de) {
            this.status.draging = false
        }
    }
}
</script>

<style lang="scss" scoped>
 .wrap {
    position: absolute; z-index: 3; border: 1px red dashed
 }
 .bw-draging {
     border-color: pink;
 }
</style>