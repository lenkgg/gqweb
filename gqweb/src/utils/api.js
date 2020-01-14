import axios from 'axios'
import {fetch, post} from './http'
import GLOBAL from '../components/global'
export function login(email, pswd, from){
  var data = { email: email,pswd:pswd, channel:from };
  var that = this;
  post('login', data)
  .then(function (resolve) {
      console.log("api:" + resolve)
      // 这里自然不用解释了，这是接口返回的参数
      GLOBAL.userInfo.onLogin(resolve);
  },(err) => {
    console.log(err)
  });
}
export function regist(email,pswd,from){
  var data = {email:email, pswd:pswd, channel:from};
  post('regist',data).then(
    (resolve) => {
      GLOBAL.userInfo.onRegist(resolve);
    },
    (err) => {
      console.log(err)
    }
  );
}
export function isLogined(){
  return GLOBAL.userInfo.isLogined
}

export function predict(symbol){
  if (symbol == "") return null;
  if (!isLogined()) return null;
  var param = {symbol:symbol};
  fetch("predict",param)
  .then(function(resolve){
    if (resolve.code == 0) {
        // 成功
        console.log(resolve.data)
        if (resolve.data.length > 0)
          {
            GLOBAL.slots.onUpdateKline(resolve.data)
          }
        else{
          //wx.showToast({title:"系统测算中",duration:3000});
        }
        //  wx.showToast({title:"系统测算中",duration:3000})
        //else

  }else{
    //wx.showToast({title:"测算失败",image:"/image/failed.png",duration:3000})
  }});
}

export function setSymbol(sid, symbol){
  if (!isLogined())return false;
  var param= {sid: sid, symbol: symbol};
  fetch("setsymbol",param)
  .then(function(resolve){
    if (resolve.code == 0){
      console.log(resolve.data)
      GLOBAL.slots.updateSlot(resolve.data)
      //wx.showToast({title:"设置成功",icon:"success",duration:3000})
    }else{
      //wx.showToast({title:"设置失败",image:"/image/failed.png",duration:3000})
    }
  });
}

export function getDefaultPredict(name){
  fetch(name, {}).then((resolve)=>{
    console.log(resolve.data)
    if (resolve.code == 0) {
        // 成功
        GLOBAL.defaultIndexs.onUpdateData(resolve.data);
        //console.log(GLOBAL.defaultIndexs.myindexs)
  }else{
    //wx.showToast({title:"获取数据失败",image:"/image/failed.png",duration:3000})
  }},
(err)=>{
  console.log(err);
});
}

export function setSlot(slot){
  if (!isLogined()) return null;
  var param = {sid:slot.sid, name:slot.name, desc:slot.desc};
  fetch("setslot", param)
  .then(function(resolve){
    if(resolve.code == 0)
      return true;
    else {
      //wx.showToast({title:"设置失败",image:"/image/failed.png",duration:3000})
    }
  });
}

export function hedge(scope,bsymbol){
  if (!isLogined()) return null;
  var param = {bsymbol:bsymbol, scope:scope};
  fetch("hedge", param)
  .then(function(resolve){
    console.log(resolve);
    if (resolve.code == 0)
      GLOBAL.hedgeScopes.onUpdateHedge(resolve.data)
    else {
      //wx.showToast({title:"获取数据失败",image:"/image/failed.png",duration:3000})
    }
  });
}

export function getDefaultHedge(){
  ////no parameter for default
  hedge("CN_HS300","sz399300");
}

export function paySlot(sid,months){
  if (!isLogined()) return null;
  var data = {"sid":sid,"months":months}
  fetch('payslot',data)
  .then(function (resolve) {
    console.log(resolve.data)
    if (resolve.code == 0) {
      var myslots = resolve.data;
      GLOBAL.slots.initSlots(myslots)
      GLOBAL.userInfo.onVip()
      //wx.showToast({title:"购买成功",icon:"success",duration:3000})
    }
  else {
      // 失败
      //wx.showToast({title:"购买失败",image:"/image/failed.png",duration:3000})
      console.log("pay slot failed.");
  }
})
}

export function getSlots(){
  if (!isLogined()) return null;
  var data = {}
  fetch('getslot',data)
  .then(function (resolve) {
    if (resolve.code == 0) {
      var myslots = resolve.data;
      GLOBAL.slots.initSlots(myslots)
    }
  else {
      // 失败
      //wx.showToast({title:"获取数据失败",image:"/image/failed.png",duration:3000})
      console.log("fetch slot failed.");
  }
})
}

export function saveUserInfo(){
  if (!isLogined()) return null;
  var data = wepy.$instance.globalData.userInfo.wxInfo
  post('updateinfo',data)
  .then(function (resolve) {
      console.log(resolve.data)
})
}

export function loadShareReport(pview,start,end){
  if (!isLogined()) return null;
  var pPage = pview
  var data = {"start":start,"end":end}
  fetch('listshare', data)
  .then(function(resolve){
    if (resolve.code == 0) {
      var share = resolve.data;
      pPage.onUpdateShare(share)
    }
  else {
      // 失败
      //wx.showToast({title:"获取数据失败",image:"/image/failed.png",duration:3000})
      console.log("list share report failed.");
  }
  })
}

export function loadNotice(pView){
  var data = {}
  fetch('notice',data)
  .then(function(resolve){
    if (resolve.code == 0) {
      var nt = resolve.data;
      pView.onNotice(nt)
    }
  else {
      // 失败
      console.log("load notice failed.");
  }
  })
}

export function LoadCNCode(){
  var data = {}
  fetch('loadcncode',data)
  .then(function(resolve){
    var code = resolve.data
    console.log(code.length);
    GLOBAL.g_cn_code = code;
    localStorage.setItem("CN_CODE",code);
    console.log("init cn_code done.");
  })
  //wx.setStorageSync("CN_CODE",code.cn_code);
}

export function PreBill(months, slots, phone){
  if (months <=0 || slots <=0)
    return
  var data = {"months":months, "slots":slots, "phone":phone}
  fetch('prebill', data)
  .then(function(resolve){
    console.log(resolve)
  })
}
/*function LoadCNCodeEX(){
  wx.downloadFile(
    url:"https://www.lilishare.com/config/stock_code_cn.json",
    success:function(res){
      var filePath = res.tempFilePath;
      wx.openDocument({
          filePath: filePath,
          success: function (res) {
            wepy.$instance.globalData.g_cn_code = res
            wx.setStorageSync("CN_CODE",res);
          },
          fail: function (res) {
            console.log("open cn_code file failed.");
          }
        })
  })
}*/
export default {
  login,
  regist,
  predict,
  setSymbol,
  getDefaultPredict,
  setSlot,
  hedge,
  getDefaultHedge,
  getSlots,
  paySlot,
  saveUserInfo,
  loadShareReport,
  loadNotice,
  LoadCNCode,
  PreBill,
};
