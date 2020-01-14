// initial state
// shape: [{ id, quantity }]
import {AuthUser} from "../../api/api";
import * as Cookies from "js-cookie"

const state = {
    loginState: {
        loginIn: false,
        user: {
            userName: ""
        }
    }
}

// getters
const getters = {
    userName: (state, getters, rootState) => {
        if (state.loginState.loginIn) {
            return state.loginState.user.userName
        }
    },
    offLine: (state, getters, rootState) => {
        return !state.loginState.loginIn;
    }
}

//actions
const actions = {
    //从服务器端校验本地登录 Cookie 有效性
    authUser({state, commit}) {
      var data = { email: email,pswd:pswd, channel:from };
      var that = this;
      post('login', data)
      .then(res => {
            debugger;
            if (res.Success) {
                commit('loginIn', {userName: res.Data.UserName});
                return true;
            } else {
                commit('loginOut');
                return false;
            }
        });
    }
}


// mutations
const mutations = {
    //登入状态
    loginIn(state, user) {
        state.loginState.loginIn = true;
        state.loginState.user = user;
        debugger;
        Cookies.set('loginState', state.loginState, {expires: 1});
    },
    //登出状态
    loginOut(state) {
        state.loginState.loginIn = false;
        state.loginState.user = {};
        Cookies.remove('loginState');
    },
    syncLoginState(state) {
        debugger;
        let cookieState = Cookies.getJSON('loginState');
        if (cookieState) {
            state.loginState = cookieState;
        }
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
