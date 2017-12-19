<template>
    <div class="container-fluid">
        <sidebar></sidebar>
        <dashboard>
            <div slot='navbar'>
                <span>Query</span>
            </div>

            <div class="container">
                <panel v-bind:table_name="this.$route.params.tab"></panel>
                <query-input></query-input>
            </div>
        </dashboard>
    </div>
</template>

<script>
    import Sidebar from '../components/sidebar.vue'
    import Dashboard from '../components/content.vue'
    import Panel from '../components/panel.vue'
    import QueryInput from '../components/query_input.vue'

    export default {
        components: {
            'sidebar': Sidebar,
            'dashboard': Dashboard,
            'panel': Panel,
            'query-input': QueryInput
        },
        data() {
            return {

            }
        }
    }
</script>

<style scoped>
    .container {
        padding: 30px 35px;
    }
</style>
