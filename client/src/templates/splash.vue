<template>
    <div class="container-fluid">
        <sidebar></sidebar>
        <dashboard>
            <div slot='navbar'>
                <span>Welcome</span>
            </div>


        </dashboard>
    </div>
</template>

<script>
    import Sidebar from '../components/sidebar.vue'
    import Dashboard from '../components/content.vue'

    export default {
        components: {
            'sidebar': Sidebar,
            'dashboard': Dashboard
        },
        data() {
            return {

            }
        }
    }
</script>
