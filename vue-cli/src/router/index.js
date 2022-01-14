import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../views/Main.vue'

Vue.use(VueRouter)

const routes = [{
    path: '/',
    name: 'Main',
    component: Main,
    children: [
        {
            path: '/',
            name: 'home',
            component: () => import('@/views/Home/Home.vue')
        }
    ]
}]
const router = new VueRouter({
    routes
})
export default router
