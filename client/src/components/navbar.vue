<template>
    <nav>
        <slot></slot>
    </nav>
</template>

<script>
    export default {
        data() {
            return {

            }
        }
    }
</script>

<style scoped>
    nav {
        font-size: 20px;
        font-weight: bold;
        color: rgb(119, 119, 119);
        padding: 17.5px 35px;
        border-bottom: 1px solid rgba(102, 97, 91, 0.45);
    }
    .info {

    }
</style>
