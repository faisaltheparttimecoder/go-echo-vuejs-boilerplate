<template>
    <section>
        <h1> Welcome to VueJs Home Page </h1>
        <p>
            You are currently on the home page of the VueJs !!!, to test the router works click on the
            buttons "Second VueJs Page" below
        </p>
        <button onclick="window.location='/backendhome'" style="background-color: lightblue">
            Backend Home Page
        </button>
        <button @click="$router.push('/second')" style="background-color: teal">
            Second VueJs Page
        </button>
        <p> By the way this is the data we got from the backend api </p>
        <ul>
            <li v-for="(d, i) in backendData" :key="i">{{i}} => {{d}}</li>
        </ul>
    </section>
</template>

<script>
    export default {
        name: 'home',
        data: function () {
            return {
                backendData: []
            }
        },
        created: function () {
            this.get(this.endpoints.backendapi).then(response => {
                this.backendData = response.data
            }).catch(error => {
                console.log(error)
            })
        }
    }
</script>

<style scoped>

</style>
