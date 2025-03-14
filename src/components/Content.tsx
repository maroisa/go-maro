import { createSignal } from "solid-js"
import "./Content.css"

export default function Content(){
    const [source, setSource] = createSignal<string>("")
    const [alias, setAlias] = createSignal<string>("")

    const [sourceMessage, setSourceMessage] = createSignal<string>(" ")
    const [aliasMessage, setAliasMessage] = createSignal<string>(" ")
    const [apiMessage, setApiMessage] = createSignal<string>("")
    
    const [shortenedUrl, setShortenedUrl] = createSignal<string>("")
    
    const validateSource = (value: string) => {
        // check empty
        if (value.length == 0){
            setSourceMessage(" ")
            return
        }

        if (value.length > 128){
            setSourceMessage("URL terlalu panjang!")
            return
        }

        // URL validation
        const pattern = /[-a-zA-Z0-9@:%_\+.~#?&//=]{2,256}\.[a-z]{2,4}\b(\/[-a-zA-Z0-9@:%_\+.~#?&//=]*)?/gi
        let valid: boolean = pattern.test(value)
        setSourceMessage(valid ? "" : "URL tidak valid!")
        if (!valid) return

    }

    const validateAlias = (value: string) => {
        setAliasMessage("")

        // check length
        if (value.length == 0){
            setAliasMessage(" ")
            return
        }

        if (value.length > 32){
            setAliasMessage("URL terlalu panjang! (maks 32 karakter)")
            return
        }

        if (/\s/gi.test(value)){
            setAliasMessage("Alias contain whitespace!")
            return
        }
    }

    const handleInput = (event: InputEvent) => {
        const target = event.target as HTMLInputElement
        if (!target) return null

        if (target.name == "source"){
            setSource(target.value)
            validateSource(target.value)
        }

        if (target.name == "alias"){
            setAlias(target.value)
            validateAlias(target.value)
        }
    }

    const submitForm = async () => {
        if (aliasMessage()) return
        if (sourceMessage()) return

        const data = {
            "source": source(),
            "alias": alias()
        }

        const result = await fetch("/api/", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data)
        })

        if (result.status != 200){
            const text = await result.text()
            setApiMessage(text)
            return
        }

        setShortenedUrl("https://go.maroisa.org/" + alias())
        setTimeout(() => {
            setShortenedUrl("")
        }, 7000);
    }

    const copyToClipboard = (event: Event) => {
        event.preventDefault()
        navigator.clipboard.writeText(shortenedUrl())
        alert("Copied to clipboard")
    }

    return <main>
        <h2 class="subtitle">URL Shortener</h2>
        <div class="row item item-form">
            <label>Source</label>
            <input 
                onInput={handleInput}
                class={sourceMessage() ? "item-input red" : "item-input"}
                placeholder="https://www.example.com" 
                name="source" 
                type="text" />
        </div>
        <p class="danger-text">{ sourceMessage() }</p>

        <div class="row item item-form">
            <label>Alias</label>
            <label>https://go.maroisa.org/</label>
            <input 
                onInput={handleInput}
                class={aliasMessage() ? "item-input red" : "item-input"}
                placeholder="(cannot blank)"
                name="alias"
                type="text" />
        </div>
        <p class="danger-text">{ aliasMessage() }</p>

        <button onClick={submitForm} class="item">Shorten!</button>
        <p class="danger-text">{ apiMessage() }</p>
        <a href="" onclick={copyToClipboard}>{ shortenedUrl() }</a>
    </main>
}