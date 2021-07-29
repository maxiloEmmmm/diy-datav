<script lang="jsx">
export default {
    render() {
        return <dashboard navs={this.navs} onLogout={this.onLogout}/>
    },
    methods: {
        onLogout(){
            this.$store.commit("auth/SetToken", "")
            this.$router.push("/auth/login")
        },
        depMenu(menus){
            return menus.map(m => {
                m.href = `${m.url}`
                m.children = m.edges.children ? this.depMenu(m.edges.children) : []
                return m
            })
        }
    },
    created(){
        this.$api[this.$apiType.AuthMenu]().then(r => this.navs = this.depMenu(r.data))
    },
    data() {
        return {
            navs: []
        }
    }
}
</script>