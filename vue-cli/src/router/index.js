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
        },
        {
            path: '/mall',
            name: 'mall',
            component: () => import('@/views/Mall/Mall')
        },
        {
            path: '/user',
            name: 'user',
            component: () => import('@/views/User/User')
        }
    ]
}]
const router = new VueRouter({
    routes
})
export default router
