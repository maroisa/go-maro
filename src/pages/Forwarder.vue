<script setup>
import { onBeforeMount } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()

onBeforeMount(() => {
    fetch("localhost:5000/" + route.params.alias)
        .then(res => {
            console.log(res)
        })
})

</script>

<template>
    <div style="padding: 1rem;">
        Redirecting to {{ $route.params.alias }}...
    </div>
</template>