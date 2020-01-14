var Symbol = require('../symbol/symbol')
var util = require('../util')
var api = require('../api')
import {fetch} from '../http'
export default class MyStocks {
  constructor(props){
    this.mysymbols = [];  ////4 index data: sh000001,sh000016, sz399001, ...
    this.myPage = null;
    this.name = props;    ////enum: index, indu, concept
    Object.assign(this, props);
    return this;
  }
  setPage(page){
    this.myPage = page;
  }

  getSymbol(sym){
    for (var i=0; i< this.mysymbols.length; i++){
      if (this.mysymbols[i].info.symbol == sym)
        return this.mysymbols[i];
    }
    return null;
  }
  getSymbolCodeList(){
    var ret = []
    for (var i=0; i< this.mysymbols.length; i++){
      ret.push(this.mysymbols[i].info);
    }
    return ret
  }

  addStock(code){
    var info = util.getSymbolInfo(code)
    if (!info){
      alert("error code.")
      return;
    }
    var param = {stk:code};
    fetch('addstock', param).then((resolve)=>{
      console.log(resolve.data)
      if (resolve.code == 0) {
        var sym = new Symbol()
        sym.info = info;
        this.mysymbols.push(sym)
      }
    });
  }
  delStock(code){
    for (var i=0; i< this.mysymbols.length; i++){
      if (code == this.mysymbols[i].info.symbol){
        var param = {stk:code};
        fetch('delstock', param).then((resolve)=>{
          console.log(resolve.data)
          if (resolve.code == 0) {
            this.mysymbols.splice(i,1)
          }
        });
      }
    }
  }
  fetch_net(page){
    var that = this
    fetch('getstocks', {}).then((resolve)=>{
      console.log(resolve.data)
      if (resolve.code == 0) {
          // 成功
          var data = resolve.data;
          //that.onUpdateData(resolve.data);
          console.log(data)
          that.mysymbols=[]

          for (var i=0; i<data.length; i++){
            var sym = new Symbol()
            var info = util.getSymbolInfo(data[i])
            sym.info = info;
            that.mysymbols.push(sym)
          }
    }else{
      console.log("error code return:"+resolve.code)
    }},
  (err)=>{
    console.log(err);
  }).then(()=>{
    page.updateView()
  });
  }

  predict(symbol){
    if (!this.getSymbol(symbol))
      return
    //api.predict(symbol)
    var that = this;
    var param = {symbol:symbol};
    fetch("predict",param)
    .then(function(resolve){
      if (resolve.code == 0) {
          // 成功
          console.log(resolve.data)
          if (resolve.data.length > 0)
            {
              that.getSymbol(data.symbol).kline = data.kline
            }
          else{
            console.log("return data length is 0.")
            //wx.showToast({title:"系统测算中",duration:3000});
          }
    }else{
      console.log("error code return:"+resolve.code)
      //wx.showToast({title:"测算失败",image:"/image/failed.png",duration:3000})
    }}).then(()=>{
      that.updateView()
    });
  }
}

export function CreateMystocks(){
  return new MyStocks();
}
