<template>
  <div class="register">
    <form role="form" class="form-horizontal">
      <ul class="formarea">
        <li>
          <label class="lit">账号：</label>
          <input type="email" placeholder="邮箱" class="textbox" required v-model="email" autocomplete="email"/>
        </li>
        <li>
          <label class="lit">密码：</label>
          <input type="password" placeholder="设置密码" class="textbox" required v-model="pwd"/>
        </li>
        <li>
          <label class="lit">确认密码：</label>
          <input type="password" placeholder="确认密码" class="textbox" required @blur="canEdit" v-model="pwdAgain"/>
        </li>
        <li class="liLink">
          <router-link class="fr" to="/login">已有账号，登陆</router-link>
        </li>
        <li>
          <input type="button" ref="isSubmit" @click="register" value="立即注册" class="button"/>
        </li>
      </ul>
    </form>
  </div>
</template>

<script>
  import router from '../router/router'

  export default{
    data: function () {
      return {
        email:'',
        pwd:'',
        pwdAgain:'',
      }
    },
    created () {
      this.$store.commit('changeIndexConf', {
        isFooter: false,
        isSearch: false,
        isBack: true,
        title: '注册'
      })
    },
    methods: {
      register: function () {
        console.log("submit regist")
        if ((this.email.indexOf('@') < 1)||(this.email.indexOf('.') < 2))
          {
            this.$msgbox('邮箱地址错误.');
            return
          }
        if (this.pwd !== this.pwdAgain)
        {
          this.$msgbox('2次密码不一致');
          return
        }
        this.GLOBAL.userInfo.regist(this.email, this.pwd, this)
      },
      canEdit: function () {
        return
        if (this.pwd === this.pwdAgain) {
          this.$refs.isSubmit.removeAttribute('disabled');
        } else {
          console.log('2次密码不一致');
        }
      },
      updateView: function(result){
        if (result)
          history.go(-1);
        else {
          alert("regist failed.")
        }
      }
    }
  }
</script>
<style >
  @import "../../static/css/login.scss";
</style>
