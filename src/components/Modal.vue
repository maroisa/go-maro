<script setup>
import { ref } from 'vue';

</script>

<template>
    <div class="modal-container" :class="type.length > 0 ? '' : 'hide'">
        <div class="modal row">
            <span>Lorem</span>
            <button @click="onClose" class="modal-btn">X</button>
        </div>
    </div>
</template>

<style>

.modal-container {
    pointer-events: none;
    position: absolute;
    height: 100%;
    width: 100%;

    display: flex;
    justify-content: flex-end;
    align-items: flex-start;

    margin-top: 0;

    transition: all 100ms;
}

.modal {
    pointer-events:visible;
    border-radius: 0.8rem;
    box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.5);
    padding: 1rem;
    margin: 1rem;
    background-color: white;

    display: flex;
    align-items: center;
    gap: 1rem;
}

.hide {
    display: none;
}

.modal-btn {
    border: none;
    width: 40px;
    height: 40px;
}

</style>