<template>
  <div class="com-homeuserinfo">
    <ul class="user-link">
      <li class="user-icon">
        <router-link to="#" class="link-box ">
          <span class="txt">头像</span>
          <img :src="logoimage">
          <span class="iconfont icon-xiangyoujiantou icon-right"></span>
        </router-link>
      </li>
      <li>
        <router-link to="#" class="link-box">
          <span class="txt">昵称</span>
          <p>{{memberName}}</p>
          <span class="iconfont icon-xiangyoujiantou icon-right"></span>
        </router-link>
      </li>
      <li>
        <router-link to="#" class="link-box">
          <span class="txt">性别</span>
          <p>{{sex}}</p>
          <span class="iconfont icon-xiangyoujiantou icon-right"></span>
        </router-link>
      </li>
    </ul>
  </div>
</template>
<script>
  export default{
    data: function () {
      return {
        memberAvatar:'',
        memberName:'',
        baseUrl:'https://www.lilishare.com:8443/',
        sex:''
      }
    },
    created: function () {
      this.$store.commit('changeIndexConf', {
        isFooter: false,
        isSearch: false,
        isBack: true,
        isShare: false,
        title: '个人信息'
      })

      let userMsg = this.GLOBAL.userInfo

      this.memberAvatar = userMsg.Info.Avata
      if(userMsg.Info.gender === 1){
        this.sex = '女'
      }else{
        this.sex = '男'
      }
      this.memberName = userMsg.Info.nick
    },
    computed:{
      logoimage:function(){
        if (this.GLOBAL.userInfo.isLogined){
          return 'static/assets/logo.jpg'
        }else{
          return 'static/assets/logo_wb.jpg'
        }
      }
    }
  }
</script>
<style lang="scss">
  @import "../../static/css/homeuserinfo.scss";
</style>
