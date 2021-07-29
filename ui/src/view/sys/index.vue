<script lang="jsx">
export default {
    render() {
        return <dashboard navs={this.navs}/>
    },
    data() {
        return {
            navs: [
                {title: '可视化', href: '/sys/view'},
                {title: '素材', href: '/sys/assets'},
                {title: '数据源', href: '/sys/tc'},
                {title: '用户', href: '/sys/user'}
            ]
        }
    }
}
</script>