<script setup>
import { onBeforeMount } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()

onBeforeMount(() => {
    fetch(import.meta.env.VITE_DATABASE_URL + "/" + route.params.alias)
        .then(res => {
            console.log(res)
            res.text()
        })
        .then(json => {
            console.log(json)
        })
})

</script>

<template>
    <div style="padding: 1rem;">
        Redirecting to {{ $route.params.alias }}...
    </div>
</template>