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
        name:"Login",
        component:()=>import('../views/Login.vue')
    },
    {
        path:'/overview-user',
        name:"User",
        component:()=>import('../views/User.vue')
    },
    {
        path:'/overview',
        name:"Admin",
        component:()=>import('../views/Admin.vue')
    },
    {
        path:'/escuderia-profile',
        name:"Escuderia",
        component:()=>import('../views/Escuderia.vue')
    }
]

const router = createRouter({
    history:createWebHistory(),
    routes
})

export default router