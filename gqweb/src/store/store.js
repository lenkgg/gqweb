import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
const store = new Vuex.Store({
  state: {
    // 公共
    comm: {
      loading: false,
      login: {
        uid: 0,
        gqkey: ''
      },
      indexConf: {
        isFooter: true, // 是否显示底部
        isSearch: true, // 是否显示搜索
        isBack: false,  // 是否显示返回
        isShare: false, // 是否显示分享
        title: '' // 标题
      }
    }
  },
  mutations: {
    /*
     * loading的显示
     * */
    isLoading: (state, status) => {
      state.comm.loading = status
    },
    /*
     * 修改header的信息
     *
     * */
    changeIndexConf: (state, data) => {
      Object.assign(state.comm.indexConf, data)
    },
    isLogin: (state,data) => {
      localStorage.setItem('uid',data.uid)
      localStorage.setItem('gqkey',data.gqkey)
      state.comm.login.uid = localStorage.getItem('uid')
      state.comm.login.gqkey = localStorage.getItem('gqkey')
    },
    logout: (state,data) => {
      localStorage.removeItem('uid')
      localStorage.removeItem('gqkey')
      state.comm.login.memberId = ''
      state.comm.login.userData = ''
    }
  },
  actions: {

  },
  getter: {

  }
})
export default store
