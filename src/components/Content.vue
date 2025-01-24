<template>
    <main class="col container">
        <h1>Link Shortener</h1>
        <input type="text" placeholder="https://siakad.uns.ac.id">
    </main>
</template>

<style scoped>

input:focus {
    outline: 3px solid rgba(0, 100, 0, 1);
    transition: all 100ms;
}

input {
    padding: 0.8rem;
    border: none;
    border-radius: 0.4rem;
    box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.2);
    outline: 1px solid rgba(0, 0, 0, 0);
    min-width: 400px;
    transition: all 100ms;
}

.container {
    flex-grow: 1;
    align-items: center;
    justify-content: center;
    gap: 4rem;
}

</style>