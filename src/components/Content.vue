<script setup>
import { ref, watch } from 'vue';

const urlValid = ref(false)
const urlValue = ref("")
const aliasValue = ref("")

const aliasLabel = ref("")
const urlLabel = ref("")

function submit(){
    if (!urlValid.value){
        urlLabel.value = "URL invalid"
        return
    }

    if (aliasValue.value.length >= 32){
        aliasLabel.value = "Alias too long"
        return
    }

    let data = JSON.stringify({ "source": urlValue.value, "alias": aliasValue.value })

    fetch("/api/", {
        method: "POST",
        body: data
    })
}

watch(urlValue, (newValue) => {
    const pattern = /[-a-zA-Z0-9@:%_\+.~#?&//=]{2,256}\.[a-z]{2,4}\b(\/[-a-zA-Z0-9@:%_\+.~#?&//=]*)?/gi
    urlValid.value = pattern.test(newValue)
    if (urlValid.value) {
        urlLabel.value = ""
    }
})

</script>

<template>
    <main>
        <div class="container col">
            <h2>URL Shortener</h2>
            <div class="row">
                <label class="input input-left">From:</label>
                <input 
                    placeholder="https://example.com"
                    v-model="urlValue" 
                    class="input input-right" 
                    :class="!urlValid ? 'red' : ''"
                    type="text">
            </div>

            <span v-show="urlLabel" v-text="urlLabel" style="color: red;"></span>

            <div class="row">
                <label class="input input-left">To: </label>
                <label class="input">go.maroisa.org/</label>
                <input v-model="aliasValue" placeholder="(blank for random)" class="input input-right" type="text">
            </div>
            <button @click="submit" class="btn">Shorten!</button>
            <span v-show="aliasLabel" v-text="aliasLabel" style="color: red;"></span>
        </div>
    </main>
</template>

<style scoped>

main {
    height: 100%;
    width: 100%;
    
    display: flex;
    justify-content: center;
    align-items: center;

}

.container {
    width: 500px;
    text-align: center;

    gap: 1.5rem;
}

.btn {
    padding: 0.8rem;
    border: none;
    border-radius: 10px;
    background-color: var(--primary);
    color: white;
    box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.2);
}

.btn:active {
    background-color: rgb(0, 100, 0);
}

.input {
    flex-grow: 1;
    padding: 0.8rem;
    background-color: white;
    border: none;
    box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.2);
    outline: 1px solid rgba(0, 0, 0, 0);
    transition: all 100ms;
}

.input:focus {
    outline: 3px solid rgba(0, 100, 0, 1);
}

.input-left {
    flex-grow: 0;
    min-width: 100px;
    border-radius: 10px 0 0 10px;
    background-color: var(--primary);
    color: var(--secondary);
}

.input-right {
    border-radius: 0 10px 10px 0;
}

.red:focus {
    outline-color: red;
}

@media screen and (max-width: 500px) {
    .row {
        flex-direction: column;
        text-align: left;
    }

    .input-left {
        border-radius: 10px 10px 0 0;
    }

    .input-right {
        border-radius: 0 0 10px 10px;
    }
}

</style>