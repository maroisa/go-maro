import { createRouter, createWebHistory } from "vue-router"

const Home = () => import("./pages/Home.vue")
const Forwarder = () => import("./pages/Forwarder.vue")

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", component: Home },
        { path: "/:alias", component: Forwarder }
    ]
})

export default router