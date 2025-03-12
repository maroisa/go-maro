import { A } from "@solidjs/router"

import Content from "../components/Content"
import "./Home.css"

export default function Home(){
    return <>
        <header>
            <h1>
                <A href="/">Go Maro</A>
            </h1>
        </header>
        <Content />
    </>
}