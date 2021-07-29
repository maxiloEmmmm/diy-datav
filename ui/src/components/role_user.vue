<script lang="jsx">
export default {
    name: "role-user",
    props: {
        role: {
            type: String,
            default: ""
        }
    },
    render(){
        return <ysz-module-widget
            v-slots={{
                title: () => <a-button onClick={this.fetch}>刷新</a-button>
            }}
        >
            <a-transfer
                data-source={this.users}
                show-search
                filter-option={(inputValue, option) => option.title.indexOf(inputValue) > -1}
                target-keys={this.targetKeys}
                render={item => item.title}
                onChange={this.onChange}
            />
        </ysz-module-widget>
    },
    data(){
        return {
            users: [],
            targetKeys: []
        }
    },
    computed: {
        _user(){
            return this.$utils.tool.makeKey(this.users, "id")
        }
    },
    created(){
        this.fetch()
    },
    methods: {
        fetch(){
            this.$api[this.$apiType.UserList]()
                .then(userResponse => {
                    this.$api[this.$apiType.RoleUser](this.role).then(response => {
                        let ids = response.data.map(id => parseInt(id))
                        let targets = []
                        this.users = userResponse.data.data.map(user => {
                            user.has_role = ids.includes(user.id)
                            if(user.has_role) {
                                targets.push(user.id + "")
                            }
                            user.key = user.id + ""
                            user.title = user.username
                            return user
                        })
                        this.targetKeys = targets
                    })
                })
        },
        onChange(targetKeys, direction, moveKeys){
            this.targetKeys = targetKeys

            let isAdd = direction === "right"
            moveKeys.forEach(k => {
                this.$api[!isAdd ? this.$apiType.RemoveUserForRole : this.$apiType.AddUserForRole](this.role, k)
                    .then(() => this.$message.success('ok'))
                    .catch((err) => this.$message.info(err))
            })
        }
    }
}
</script>