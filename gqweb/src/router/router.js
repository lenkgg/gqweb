import Vue from 'vue'
import Router from 'vue-router'
import App from '../app'
import Index from '../page/index'
import Login from '../page/login'
import Register from '../page/register'
import Home from '../page/home'
import HomeUserInfo from '../page/homeUserInfo'
import Choose from '../page/choose'

Vue.use(Router)

const routes =[
    {
      path: '/',
      component: App,
      children: [
        // {path: '/index/:type', name: 'index', component: Index, meta:{requireAuth: true }},
        {path: '/index', name: 'index', component: Index, meta: { notRequiredLogin: true}},
        {path: '/choose', name: 'choose', component: Choose, meta: { notRequiredLogin: true}},
        {path: '/login', name: 'login', component: Login},
        {path: '/home', name: 'home', component: Home},
        {path: '/HomeUserInfo', name: 'homeUserInfo', component: HomeUserInfo},
        {path: '/register', name: 'register', component: Register},
      ]
    }
  ]

  const router = new Router({
    routes: routes, // short for routes: routes
    linkActiveClass: 'active',  // router-link的选中状态的class，也有一个默认的值
    mode: 'history'
  });

  //这个是请求页面路由的时候会验证token存不存在，不存在的话会到登录页
  /*router.beforeEach(function (to,from,next) {
    var gqkey = localStorage.getItem('gqkey')
    // if(to.path === '/home'){
      if(!gqkey){
        next({ path: '/login' })
    //  }
    }
    next()
  })*/

  export default router;
