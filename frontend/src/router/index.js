import{createRouter, createWebHistory} from 'vue-router'
import homePage from '../view/homePage.vue'
import LoginPage from '../view/loginPage.vue'
import Error404 from '../view/error404.vue'
import Dataset from '../view/aichat.vue'
import Grading from '../view/Grading.vue'
import ModelDetail from '../view/ModelDetail.vue'
import Profile from '../view/components/userInfo.vue'
import { compile } from 'vue'
import RegisterPage from '@/view/registerPage.vue'

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
        path:'/aichat',
        name: 'AIChat',
        component: Dataset
    },
    {
        path:'/my-model',
        name: 'MyModel',
        component: Grading
    },
    {
        path: '/model/:id',
        name: 'ModelDetail',
        component: ModelDetail
    },
    {
        path:'/profile',
        name: 'Profile',
        component:Profile
    },
    {
        path: '/register',
        name: 'Register',
        component: RegisterPage
    }
]

const router = createRouter({
    history: createWebHistory(), // 使用 history 模式
    routes,
});

export default router



