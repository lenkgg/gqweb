<template>
  <div class="com-home">
    <div class="user-info clearfix">
      <router-link v-if="isLogined" to="/homeUserInfo" class="link-box clearfix">
        <img :src="logoimage">
        <label>{{name}}</label>
        <span class="iconfont icon-xiangyoujiantou icon-right"></span>
      </router-link>
      <router-link v-else="isLogined" to="/login" class="link-box clearfix">
        <img :src="logoimage">
        <label>未登录</label>
      </router-link>
    </div>
    <div class="user-link">
      <ul>
        <li>
          <router-link to="/search/own" class="link-box">
            <span class="iconfont Hui-iconfont-feedback1 icon-detail"></span>
            <span class="txt">我的圈子</span>
            <span class="iconfont icon-xiangyoujiantou icon-right"></span>
          </router-link>
        </li>
        <li>
          <router-link to="/modif" class="link-box">
            <span class="icon iconfont icon-yijianfankui"></span>
            <span class="txt">修改密码</span>
            <span class="iconfont icon-xiangyoujiantou icon-right "></span>
          </router-link>
        </li>
      </ul>
    </div>
    <button v-if="isLogined" class="logout" @click="logout">退出登陆</button>
  </div>
</template>
<script>
  import axios from 'axios'
  import router from '../router/router'
  //import GLOBAL from '../components/global'

  export default{
    data: function () {
      return {
        uid: 0,
        vip: 0,
        memberAvatar:'',
        memberName:'',
        role: '',
        baseUrl: this.$store.state.comm.imgUrl,
      }
    },
    created: function () {
      this.isLogined = this.GLOBAL.userInfo.isLogined;
      this.uid = this.GLOBAL.userInfo.uid;
      this.vip = this.GLOBAL.userInfo.vip;

      let userMsg = this.GLOBAL.userInfo.Info
      this.role = userMsg.gender
      this.memberAvatar = userMsg.avata
      if (userMsg.nick == ""){
        this.memberName = "用户"+this.uid
      }else{
        this.memberName = userMsg.nick
      }


      this.$store.commit('changeIndexConf', {
        isFooter: true,
        isSearch: false,
        isBack: false,
        title: '个人首页'
      })
    },
    methods: {
      logout:function () {
        let vm = this
        let url = vm.$store.state.comm.apiUrl + '/logout'
        axios.get(url).then(function (res) {
          console.log(res.data.result)
          if(res.data.result === 1){
            vm.$store.commit('logout')
            router.push('index/fresh')
          }
        }).catch(function (error) {
          console.log(error)
        })
      }
    },
    computed:{
      isLogined:function(){
        return this.GLOBAL.userInfo.isLogined;
      },
      logoimage:function(){
        if (this.GLOBAL.userInfo.isLogined){
          return 'static/assets/logo.jpg'
        }else{
          return 'static/assets/logo_wb.jpg'
        }
      },
      name:function(){
        this.uid = this.GLOBAL.userInfo.uid
        this.vip = this.GLOBAL.userInfo.vip
        if ((this.GLOBAL.userInfo.Info.nick == "")||(this.GLOBAL.userInfo.Info.nick == "undefined")){
          this.memberName = "用户"+this.GLOBAL.userInfo.uid
        }else{
          this.memberName = this.GLOBAL.userInfo.Info.nick
        }
        return this.memberName;
      }
    }
  }
</script>
<style lang="scss">
  @import "../../static/css/home.scss";
</style>
