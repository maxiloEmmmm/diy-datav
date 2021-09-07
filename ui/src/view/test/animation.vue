<script lang="jsx">
export default {
    render() {
        return <div class="t-a">Text</div>
    }
}
</script>

<style lang="scss">
.t-a {
    animation: t-s1 1s;
    animation-iteration-count: infinite;
    overflow: hidden;
}
@keyframes t-s1 {
    0% {
        color: red
    }
    25% {
        color: green;
    }
    50% {
        color: red
    }
    75% {
        color: green
    }
    100% {
        color: red;
    }
}
</style>