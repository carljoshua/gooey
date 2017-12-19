<template>
    <table border='0' cellspacing='0'>
        <tr class="theader">
            <th v-for="col_name in dataset.col_names">
                {{ col_name }}
            </th>
        </tr>
        <tr v-for="row in dataset.rows">
            <td v-for="cell in row">{{ cell }}</td>
        </tr>
    </table>
</template>

<script>
    export default {
        props: ['dataset'],
        data() {
            return {

            }
        }
    }
</script>

<style scoped>
    table {
        width: 98%;
        border-radius: 5px;
        background: white;
        padding: 12px;
        box-shadow: rgba(110, 110, 110, 0.1) 0px 5px 15px;
    }
    td {
        padding: 10px 15px 0px 15px;
    }
    th {
        padding: 7px 15px;
        text-align: left;
        border-bottom: 1px solid #ddd;
    }
</style>
