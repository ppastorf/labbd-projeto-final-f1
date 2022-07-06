import {createRouter, createWebHistory}  from 'vue-router'

const routes = [
    {
        path:'/SignupPiloto',
        name:"SignupPiloto",
        component:()=>import('../views/SignupPiloto.vue')
    },
    {
        path:'/SignupEscuderia',
        name:"SignupEscuderia",
        component:()=>import('../views/SignupEscuderia.vue')
    },
    {
        path:'/loginPage',
        name:"LoginPage",
        component:()=>import('../views/LoginPage.vue')
    },
    {
        path:'/user-profile',
        name:"User",
        component:()=>import('../views/User.vue')
    },
    {
        path:'/admin-profile',
        name:"Admin",
        component:()=>import('../views/Admin.vue')
    },
    {
        path:'/overview',
        name:"overview",
        component:()=>import('../views/Overview.vue')
    }
]

const router = createRouter({
    history:createWebHistory(),
    routes
})

export default router