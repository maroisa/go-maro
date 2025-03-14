import { A } from "@solidjs/router"

import { DarkIcon, LightIcon } from "../components/Icon"

import Content from "../components/Content"
import "./Home.css"
import { createSignal, Show } from "solid-js"

export default function Home(){
    const [isDark, setIsDark] = createSignal<boolean>(true)

    function toggleTheme(){
        document.documentElement.setAttribute("data-theme", isDark() ? "light" : "dark")
        setIsDark(!isDark())
    }

    return <>
        <header class="row">
            <h1>
                <A href="/">Go Maro</A>
            </h1>
            <button onClick={toggleTheme} class="toggle-btn">
                <Show when={isDark()} fallback={<LightIcon size={20} />}>
                    <DarkIcon size={20} />
                </Show>
            </button>
        </header>
        <Content />
    </>
}

