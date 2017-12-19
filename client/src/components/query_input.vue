<template>
    <div>
        <div class="textarea"></div>
        <button class="btn">Execute</button>
    </div>
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
    .textarea {
        border: 0.5px solid lightgray;
        height: 200px;
        width: 200px;
        background: white;
    }
    .btn {
        border: none;
        padding: 10px 15px;
        background: #f9690e;
        color: white;
        outline: none;
    }
</style>
