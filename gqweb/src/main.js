// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Element from 'element-ui'
Vue.use(Element)
//import App from './App'
import router from './router/router'
import global from './components/global'
import store from './store/store'
import 'element-ui/lib/theme-chalk/index.css';
import {fetch, post} from './utils/http'

//Vue.config.productionTip = false
Vue.prototype.$loading = Element.Loading.service
Vue.prototype.$msgbox = Element.MessageBox
Vue.prototype.$alert = Element.MessageBox.alert
Vue.prototype.$confirm = Element.MessageBox.confirm
Vue.prototype.$prompt = Element.MessageBox.prompt
// Vue.prototype.$notify = Notification
Vue.prototype.$message = Element.Message

//开启debug模式
Vue.config.debug = true
window.log = console.log

let data = {
  el:'#app',
  router,
  store
}

Vue.prototype.GLOBAL = global

Vue.prototype.post = post
Vue.prototype.fetch = fetch

//创建一个app实例，并且挂载到选择符#app匹配的元素上
new Vue(data).$mount('#app')
