<template>
    <div class="container-fluid">
        <sidebar v-on:changeTable="getData"></sidebar>
        <dashboard>
            <div slot='navbar'>
                <span>Structure</span>
            </div>

            <div class="container">
                <panel v-bind:table_name="this.$route.params.tab"></panel>
                <table-view v-bind:dataset="table_data"></table-view>
            </div>
        </dashboard>
    </div>
</template>

<script>
    import Sidebar from '../components/sidebar.vue'
    import Dashboard from '../components/content.vue'
    import Panel from '../components/panel.vue'
    import Table from '../components/table.vue'

    export default {
        components: {
            'sidebar': Sidebar,
            'dashboard': Dashboard,
            'panel': Panel,
            'table-view': Table
        },
        data() {
            return {
                table_data: []
            }
        },
        created() {
            this.$http.get('http://localhost:8000/describe?table=' + this.$route.params.tab).then(function(data){
                this.table_data = data.body
            })
        },
        methods: {
            getData: function() {
                this.$http.get('http://localhost:8000/describe?table=' + this.$route.params.tab).then(function(data){
                    this.table_data = data.body
                })
            }
        }
    }
</script>

<style scoped>
    .container {
        padding: 30px 35px;
    }
</style>
