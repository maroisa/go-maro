<script setup>
import { onBeforeMount } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()

onBeforeMount(async () => {
    const result = await fetch("/api/" + route.params.alias)
    
    if (result.status != 200) window.location.href = "/"

    const json = await result.json()

    let source = json.source

    if (!/http(s?)/gi.test(source)){
        source = "https://" + source
    }

    window.location.href = source
})

</script>

<template>
    <div style="padding: 1rem;">
        Redirecting to {{ $route.params.alias }}...
    </div>
</template>