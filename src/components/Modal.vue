<template>
    <div class="modal-container">
        <div class="modal">
            lorem
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
    justify-content: stretch;
    align-items: flex-end;

    margin-top: 0;

    transition: all 100ms;
}

.modal {
    border-radius: 0.8rem;
    box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.5);
    flex-grow: 1;
    padding: 1rem;
    margin: 1rem;
    background-color: white;
}

</style>