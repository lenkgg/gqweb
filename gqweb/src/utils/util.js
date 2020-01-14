// 下面是util.js的东西
var code = require('./stock_code_cn');
//获取cookie、
export function getCookie(name) {
 var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
 if (arr = document.cookie.match(reg))
  return (arr[2]);
 else
  return null;
}

//设置cookie,增加到vue实例方便全局调用
export function setCookie (c_name, value, expiredays) {
 var exdate = new Date();
 exdate.setDate(exdate.getDate() + expiredays);
 document.cookie = c_name + "=" + escape(value) + ((expiredays == null) ? "" : ";expires=" + exdate.toGMTString());
};

//删除cookie
export function delCookie (name) {
 var exp = new Date();
 exp.setTime(exp.getTime() - 1);
 var cval = getCookie(name);
 if (cval != null)
  document.cookie = name + "=" + cval + ";expires=" + exp.toGMTString();
};

/**
 * request请求封装
 * url   传递方法名
 * types 传递方式(1,GET,2,POST)
 * data  传递数据对象
 */
const URL= 'https://www.lilishare.com:8443/'
export function commonAjax(url, method, data) {

    // 获取公共配置
    // var app = getApp()

    // 公共参数（一般写接口的时候都会有些公共参数，你可以事先把这些参数都封装起来，就不用每次调用方法的时候再去写，）
    var d = {
        token: '123456789',// 例如：这是我们自己的验证规则
    }

    // 合并对象(公共参数加传入参数合并对象) mergeObj对象在下面
    var datas = mergeObj(d, data)
    //console.log(wx.getStorageSync("vk"));
    var vk =localStorage.getItem("vk")
    console.log("cookie:"+vk)
    console.log("net request:" + method +" " + url + " " +data )
    // 这是es6的promise版本库大概在1.1.0开始支持的，大家可以去历史细节点去看一下，一些es6的机制已经可以使用了
    var promise = new Promise(function (resolve, reject, defaults) {
    // 封装reuqest
    wx.request({
      url: URL + url,
      data: datas,
      method: method,
      header: {
        'content-type': (method === 'GET') ? 'application/json; charset=utf-8':'application/x-www-form-urlencoded',
        'Cookie': vk
      },
      success: resolve,
      fail: reject,
      complete: defaults,
    })
  });
  return promise;
}

/**
 * object 对象合并
 * o1     对象一
 * o2     对象二
 */
export function mergeObj(o1, o2) {
  for (var key in o2) {
    o1[key] = o2[key]
  }
  return o1;
}

export function formatTime(date) {
  var year = date.getFullYear()
  var month = date.getMonth() + 1
  var day = date.getDate()

  var hour = date.getHours()
  var minute = date.getMinutes()
  var second = date.getSeconds()


  return [year, month, day].map(formatNumber).join('/') + ' ' + [hour, minute, second].map(formatNumber).join(':')
}

export function formatNumber(n) {
  n = n.toString()
  return n[1] ? n : '0' + n
}

export function getSymbolInfo (sym) {
  if ( sym == "" || typeof(sym) == "undefined"){
    console.log("symbol is null,getSymbolInfo() failed.")
    return null;
  }
  var info = {}
  if (sym == "sh000001"){
    info["symbol"] = "sh000001";
    info["name"] = "上证指数";
    info["indu"] = "大盘指数";
    return info;
    }
    if (sym == "sh000016"){
    info["symbol"] = "sh000016";
    info["name"] = "上证50";
    info["indu"] = "大盘指数";
    return info;
    }
    if (sym == "sz399001"){
    info["symbol"] = "sz399001";
    info["name"] = "深证综指";
    info["indu"] = "大盘指数";
    return info;
    }
    if (sym == "sz399006"){
    info["symbol"] = "sz399006";
    info["name"] = "创业板指";
    info["indu"] = "大盘指数";
    return info;
    }
    if (sym == "sz399300"){
    info["symbol"] = "sz399300";
    info["name"] = "沪深300";
    info["indu"] = "大盘指数";
    return info;
  }
  //var code = {}
  /*
  if ((this.GLOBAL.g_cn_code) && (this.GLOBAL.g_cn_code.length > 0))
    code = this.GLOBAL.g_cn_code
  else
    code = localStorage.getItem("CN_CODE");*/
  for (var i = 0; i < code.length; i++){
    if ((code[i].symbol == sym) ||(code[i].name == sym)){
      info = code[i]
      if (info.indu.length > 16){
          info.indu = info.indu.substring(0, 16) + "...";
        }
      return info;
    }
  }
  return null;
}
export function dateFormat(date, fmt) { // author: meizz
    var o = {
        "M+": date.getMonth() + 1, // 月份
        "d+": date.getDate(), // 日
        "h+": date.getHours(), // 小时
        "m+": date.getMinutes(), // 分
        "s+": date.getSeconds(), // 秒
        "S": date.getMilliseconds() // 毫秒
    };
    if (/(y+)/.test(fmt))
        fmt = fmt.replace(RegExp.$1, (date.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
        if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
            return fmt
}
export default {
  getCookie,
  setCookie,
  delCookie,
  formatTime,
  commonAjax,
  getSymbolInfo,
  dateFormat
}
