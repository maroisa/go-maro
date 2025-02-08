<script setup>
import { onBeforeMount } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()

onBeforeMount(() => {
    fetch("/api/" + route.params.alias)
        .then(res => res.json())
        .then(json => window.location.href = json.source)
})

</script>

<template>
    <div style="padding: 1rem;">
        Redirecting to {{ $route.params.alias }}...
    </div>
</template>