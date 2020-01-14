import axios from 'axios' //引用axios
import QS from "qs"
import {Promise} from 'es6-promise' //引入Promise
var util = require('./util')//引用刚才我们创建的util.js文件，并使用getCookie方法

// axios 配置
axios.defaults.timeout = 5000;
axios.defaults.withCredentials =true;
//axios.defaults.baseURL = 'https://www.lilishare.com:8443/'; //这是调用数据接口
//axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
//axios.defaults.headers.get['Content-Type'] = 'application/json; charset=utf-8';
axios.defaults.headers = {
  "Content-Type": "application/x-www-form-urlencoded"
}

// http request 拦截器，通过这个，我们就可以把Cookie传到后台
axios.interceptors.request.use(
  config => {
    //const uid = util.getCookie('uid'); //获取Cookie
    //const gqkey = util.getCookie('gqkey');
    var vk =localStorage.getItem("vk")
    //config.data = JSON.stringify(config.data);
    config.headers = {
      "Content-Type": "application/x-www-form-urlencoded",
      "Access-Control-Allow-Origin": " *",
      'Authorization': vk
    };
    return config;
  },
  err => {
    return Promise.reject(err);
  }
);


// http response 拦截器
axios.interceptors.response.use(
  response => {
//response.data.errCode是我接口返回的值，如果值为7 gqkey校验失败，然后跳转到登录页，这里根据大家自己的情况来设定
    if(response.code == 7) {
      router.push({
        path: '/login',
        query: {redirect: router.currentRoute.fullPath} //从哪个页面跳转
      })
    }
    return response;
  },
  error => {
    return Promise.reject(error.response)
  });

export default axios;

/**
 * fetch 请求方法
 * @param url
 * @param params
 * @returns {Promise}
 */
export function fetch(url, params = {}) {

  return new Promise((resolve, reject) => {
    //axios.defaults.headers.get['Content-Type'] = 'application/json; charset=utf-8';
    axios.get(url, {
      params: params
    })
    .then(response => {
      resolve(response.data);
    })
    .catch(err => {
      reject(err)
    })
  })
}

/**
 * post 请求方法
 * @param url
 * @param data
 * @returns {Promise}
 */
export function post(url, data = {}) {
  return new Promise((resolve, reject) => {
    //axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
    axios.post(url, QS.stringify(data))
      .then(response => {
        resolve(response.data);
      }, err => {
        reject(err);
      })
  })
}

/**
 * patch 方法封装
 * @param url
 * @param data
 * @returns {Promise}
 */
export function patch(url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.patch(url, data)
      .then(response => {
        resolve(response.data);
      }, err => {
        reject(err);
      })
  })
}

/**
 * put 方法封装
 * @param url
 * @param data
 * @returns {Promise}
 */
export function put(url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.put(url, data)
      .then(response => {
        resolve(response.data);
      }, err => {
        reject(err);
      })
  })
}
