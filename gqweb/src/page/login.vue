<template>
  <div class="login">
    <form role="form" class="form-horizontal" name="RegisterForm">
      <ul class="formarea">
        <li>
          <label class="lit">账号：</label>
          <input type="email" placeholder="邮箱" class="textbox" required v-model="email" autocomplete/>
        </li>
        <li>
          <label class="lit">密码：</label>
          <input type="password" placeholder="登陆密码" class="textbox" required v-model="psw" autocomplete/>
        </li>
        <li class="liLink lg_liLink">
          <!-- span><label><input type="checkbox"/>记住密码</label></span -->
          <span><router-link to="/register">新用户注册</router-link></span>
          <span><router-link to="/find_pwd">忘记密码?</router-link></span>
        </li>
        <li>
          <input type="button" @click="loginAction" value="立即登陆" class="button"/>
        </li>
      </ul>
    </form>
  </div>

</template>

<script>
  import router from '../router/router'

  export default{
    created () {
      this.$store.commit('changeIndexConf', {
        isFooter: false,
        isSearch: false,
        isBack: true,
        isShare: false,
        title: '登陆页面'
      })
    },
    data: function(){
      return {
        email:'',
        psw:'',
        logined:false,
      }
    },
    watch:{
      /*logined(nv,ov){
        this.logined = nv
        console.log('back to prepage.')
        if ((ov != nv) &&(nv))
        {

          history.go(-1);
        }
      }*/
    },
    methods: {
      loginAction: function () {
        this.login(this.email, this.psw)
      },
      updateView: function(){
        console.log("login page onupdateview()")
        if (this.GLOBAL.userInfo.isLogined){
          console.log('back to prepage.')
          history.go(-1);
        }
        else {
          alert("login failed.")
        }
      },
      login:function(email,pswd,from=900000000){
        // 调用登录接口
        //api.login(email,pswd,from)
        var that = this;
        var data = { email: email,pswd:pswd, channel:from };
        that.post('login', data)
        .then(function (resolve) {
            // 这里自然不用解释了，这是接口返回的参数
            if (resolve.code == 0) {
                // 成功
                that.GLOBAL.userInfo.uid = resolve.data.info.uid;
                that.GLOBAL.userInfo.email = resolve.data.info.email;
                that.GLOBAL.userInfo.isLogined = true;
                that.GLOBAL.userInfo.Info = resolve.data.info
                //that.logined = true;
                //that.GLOBAL.slots.initSlots(slts);
                return true;
            } else {
                // 失败
                //wx.showToast({title:"登陆失败",image:"/image/failed.png",duration:3000})
                console.log("login failed.");
                that.GLOBAL.userInfo.isLogined = false;
                //that.logined = false;
                //alert("login failed.")
                return false;
            }
        },(err) => {
          console.log(err)
        }).then((ret)=>{
          that.updateView()
        });
      }
    }
  }
</script>

<style style="scss">
  @import "../../static/css/login.scss";
</style>
