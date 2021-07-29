<script lang="jsx">
export default {
    render(){
        return <div><tool-form onDone={this.onDone} show={true} ref="form"/></div>
    },
    mounted(){
        this.$refs.form.setFields([
            {title: "用户名", type: "string", field: "username", validate: "required"},
            {title: "密码", type: "string", field: "password", validate: "required"}
        ])
        this.$refs.form.setModels([
            {title: "登录", key: "token", layout_group: 1}
        ])
        this.$refs.form.setModel("token")
    },
    methods: {
        onDone(data){
            this.$api[this.$apiType.AuthLogin](data)
                .then(response => {
                    this.$store.commit("auth/SetToken", response.data.token)
                    let forward = this.$store.state.auth.forward
                    if(forward) {
                        this.$store.commit("auth/SetForward", "")
                        window.location.hash = forward
                    }else {
                        this.$router.push("/")
                    }
                })
        }
    }
}
</script>