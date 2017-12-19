<template>
    <div>
        <ul>
            <li v-for="button in buttons">
                <router-link v-bind:to="'/gooey/' + button + '/' + table_name">
                    {{ button }}
                </router-link>
            </li>
        </ul>
    </div>
</template>

<script>
    export default {
        props: ['table_name'],
        data() {
            return {
                buttons: ['browse', 'structure', 'query']
            }
        }
    }
</script>

<style scoped>
    ul {
        list-style: none;
        padding: 0px;
        margin-bottom: 30px;
    }
    ul > li{
        display: inline-block;
        margin-right: 15px;
        cursor: pointer;
    }
    a {
        text-transform: uppercase;
        text-decoration: none;
        color: #777;
        -webkit-transition: color 0.3s;
    }
    a:hover {
        color: #111;
    }
</style>
