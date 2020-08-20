<template>
    <section>
        <h1> Second VueJS Page !!! </h1>
        <p> This is the second page of the vuejs </p>
        <button @click="$router.push('/')" style="background-color: teal">
            VueJs Home
        </button>
    </section>
</template>

<script>
    export default {
        name: "second"
    }
</script>

<style scoped>

</style>