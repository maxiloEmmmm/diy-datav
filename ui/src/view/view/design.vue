<script lang="jsx">
export default {
    render() {
        return this.visible ? <a-alert
            message="Alert Message Text"
            type="success"
            closable
            after-close={this.handleClose}
        /> : null
    },
    data() {
        return {
            visible: true,
        };
    },
    methods: {
        handleClose() {
            this.visible = false;
        },
    },
}
</script>