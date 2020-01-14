var api = require('./api')
import axios from 'axios'
import {fetch, post} from './http'

export default class UserInfo{
  constructor(props){
    this.uid = 0;
    this.wxid = "";
    this.vip = 0;
    this.phone = "";
    this.email = "";
    this.channel = 0;
    this.scene = 0;  /// come from
    this.Info = {};
    this.isLogined = false;
    this.page = null;
    Object.assign(this,props);

    return this;
  }
    login(email,pswd,callback=null,from=900000000){
      // 调用登录接口
      var that = this;
      //api.login(email,pswd,from)
      var data = { email: email,pswd:pswd, channel:from };
      var that = this;
      post('login', data)
      .then(function (resolve) {
          // 这里自然不用解释了，这是接口返回的参数
          if (resolve.code == 0) {
              // 成功
              that.uid = resolve.data.info.uid;
              that.email = resolve.data.info.email;
              var slts = resolve.data.slots;
              this.GLOBAL.slots.initSlots(slts);
              that.vip = parseInt(slts.length/5);
              that.isLogined = true;
          } else {
              // 失败
              //wx.showToast({title:"登陆失败",image:"/image/failed.png",duration:3000})
              console.log("login failed.");
              that.isLogined = false;
          }
      },(err) => {
        console.log(err)
      }).then(callback);
    }

    onLogin(resolve){
      var that = this
      if (resolve.code == 0) {
          // 成功
          that.uid = resolve.data.info.uid;
          that.email = resolve.data.info.email;
          var slts = resolve.data.slots;
          that.vip = parseInt(slts.length/5);
          that.isLogined = true;
          if (that.page != null)
            that.page.updateView(sucess);
      } else {
          // 失败
          //wx.showToast({title:"登陆失败",image:"/image/failed.png",duration:3000})
          console.log("login failed.");
          that.isLogined = false;
          if (that.page != null)
            that.page.updateView(false);
      }
    }

    onVip(){
      this.vip = parseInt(this.GLOBAL.slots.slots.length/5);
    }

    regist(email,pswd,page=null,from=900000000){
      this.page = page
      api.regist(email,pswd,from)
    }
    onRegist(resolve){
      this.onLogin(resolve)
  }
}
  export function CreateUser(){
    return new UserInfo()
  }
