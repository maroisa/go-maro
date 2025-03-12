import { lazy } from "solid-js"
import { Router, Route } from "@solidjs/router"

const Home = lazy(() => import("./pages/Home"))
const Forwarder = lazy(() => import("./pages/Forwarder"))

export default function App(){
    return <Router>
        <Route path="/" component={Home} />
        <Route path="/:alias" component={Forwarder} />
    </Router>
}