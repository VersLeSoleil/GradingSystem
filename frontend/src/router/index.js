import{createRouter, createWebHistory} from 'vue-router'
import homePage from '../view/homePage.vue'
import LoginPage from '../view/loginPage.vue'
import Error404 from '../view/error404.vue'
import Dataset from '../view/Dataset.vue'
import Online from '../view/Grading.vue'
import { compile } from 'vue'

const routes = [
    {
        path: '/',
        redirect: '/login', 
    },
    {
        path: '/home',
        name: 'Home',
        component: homePage
    },
    {
        path: '/login',
        name: 'Login',
        component:LoginPage
    },
    {
        path: '/404',
        name: 'Error404',
        component: Error404,
        hidden: true,
    },
    {
        path:'/dataset',
        name: 'Dataset',
        component: Dataset
    },
    {
        path:'/online',
        name: 'OnlineGrading',
        component: Online
    }
]

const router = createRouter({
    history: createWebHistory(), // 使用 history 模式
    routes,
});

export default router



