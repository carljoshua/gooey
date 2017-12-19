<template>
    <div id='sidebar'>
        <ul>
            <center>GOOEY</center><hr />
            <p>TABLES</p>
            <li v-for="table in tab_list.tables" v-on:click="changeTab">
                <router-link v-bind:to="'/gooey/browse/' + table">
                    {{ table }}
                </router-link>
            </li>
        </ul>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                tab_list: []
            }
        },
        created() {
            this.$http.get('http://localhost:8000/tables/').then(function(data){
                this.tab_list = data.body
            })
        },
        mounted() {
            document.getElementById('sidebar').style.height = window.innerHeight + "px"
        },
        methods: {
            changeTab: function() {
                this.$emit('changeTable');
            }
        }
    }
</script>

<style scoped>
    ul {
        list-style: none;
        padding: 0;
        background: #212120;
        height: 100%;
    }
    li {
        padding: 0px 0px 20px 30px;
    }
    a {
        -webkit-transition: color 0.3s;
        color: #66615b;
        text-decoration: none;
        cursor: pointer;
    }
    a:hover {
        color: white
    }
    #sidebar {
        height: 100px;
        width: 20%;
    }
    p {
        color: white;
        font-size: 14px;
        padding: 20px;
    }
    center {
        padding: 20px;
        color: white;
        font-weight: bold;
    }
    hr {
        margin: 0px 10px;
        padding: 0px;
        border: none;
        height: 1px;
        background: #444;
    }
</style>
