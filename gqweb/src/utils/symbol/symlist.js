/*用于管理指数列表，name标示指数、行业、概念*/

//var Symbol = require('./symbol')
import Symbol from '../symbol/symbol'
var util = require('../util')
var api = require('../api')
import {fetch} from '../http'
export default class Symlist {
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
    updateView(){
      if (this.myPage != null){
        console.log("notify to update view.")
        this.myPage.updateView();
      }

    }

    fetch_net(page){
      console.log("cnindex:enter fetch_net()")
      //api.getDefaultPredict(this.name);
      var that = this
      fetch(this.name, {}).then((resolve)=>{
        console.log(resolve.data)
        if (resolve.code == 0) {
            // 成功
            var data = resolve.data;
            //that.onUpdateData(resolve.data);
            console.log(data)
            that.mysymbols=[]

            for (var i=0; i<data.length; i++){
              var sym = new Symbol()
              sym.info = util.getSymbolInfo(data[i].symbol)
              sym.setKline(data[i].kline);
              that.mysymbols.push(sym)
            }
            return true;
      }else{
        console.log("error code return:"+resolve.code)
        return false;
      }},
    (err)=>{
      console.log(err);
      return false;
    }).then(()=>{
      console.log("fetch done.")
      page.updateView()
    });
    }

    updateItem(symbol){
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
            //  wx.showToast({title:"系统测算中",duration:3000})
            //else

      }else{
        console.log("error code return:"+resolve.code)
        //wx.showToast({title:"测算失败",image:"/image/failed.png",duration:3000})
      }}).then(()=>{
        that.updateView()
      });
    }
    ///////////on update net response////

}

export function CreateSymlist(name){
  return new Symlist(name);
}
