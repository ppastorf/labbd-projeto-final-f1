import {createRouter, createWebHistory}  from 'vue-router'

const routes = [
    {
        path:'/Signup',
        name:"Signup",
        component:()=>import('../views/Signup.vue')
    },
    {
        path:'/login',
        name:"Login",
        component:()=>import('../views/Login.vue')
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