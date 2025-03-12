import { useParams } from "@solidjs/router"
import { onMount, Show } from "solid-js"

export default function Forwarder(){
    const params = useParams()
    
    onMount(async () => {
        if (params.alias.includes("%20")) return

        const result = await fetch("/api/" + params.alias)
        if (result.status != 200) window.location.href = "/"

        const json = await result.json()
        let source = json.source

        if (!/http(s?)/gi.test(source)){
            source = "https://" + source
        }

        window.location.href = source
    })

    return <>
        <Show 
            when={!params.alias.includes("%20")}
            fallback={<p>Invalid! (Terdapat whitespace)</p>}>
            Redirecting to {params.alias}...
        </Show>
    </>
}