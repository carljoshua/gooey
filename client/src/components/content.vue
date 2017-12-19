<template>
    <div class='main-content'>
        <navbar>
            <slot name='navbar'></slot>
        </navbar>
        <slot></slot>
    </div>
</template>

<script>
    import Navbar from './navbar.vue'

    export default {
        components: {
            'navbar': Navbar
        },
        data() {
            return {

            }
        }
    }
</script>

<style scoped>
    .main-content {
        width: 80%;
        background: #f4f3ef
    }
</style>
