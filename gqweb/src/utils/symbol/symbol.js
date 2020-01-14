var api = require("../api");
var util = require("../util")

export default class Symbol{
  constructor(props){
      this.info = {},
      /*{
      symbol:"",
      name:"",
      indu:"",
      desc:"",
      predict:"",
      color:"",
      sid:0,
    },*/
    this.type = "",  ///enum: index, stock, future
    this.kline = {},
    this.predict = 0,   //1:up, -1:down
    /*bar: [],
    ptime: unix timestamp, predict time
    */
    Object.assign(this, props);
    return this;
  }
   ////get the datetime of  kline[index]
  getBarofDay(index){
    if (index <0)
      index = this.kline.length+index;

    if (index<0 || index >= this.kline.length)
      return null

      var rlt = {}
      var d = this.kline[index].split(',');
      rlt.date = d[0].split(' ')[0];
      rlt.open = d[1]/1; //开盘价
      rlt.high = d[2]/1; //最高价
      rlt.low = d[3]/1; //最低价
      rlt.close = d[4]/1; //现价

      return rlt;

  }
  init(sym){
    if (typeof(sym)== "undefined" || sym == "" ){
      this.info.symbol = "";
      this.info.name = "";
      this.info.indu = "";
      this.info.area = "";
      this.info.markt ="";
      this.info.list_date = "";
      this.info.sid = 0;
    }else{
       //sym = sym.toLocaleUpperCase()
       var info = util.getSymbolInfo(sym);
       if (info == null){
         this.info.symbol = "";
         this.info.name = "";
         this.info.indu = "";
         this.info.area = "";
         this.info.markt ="";
         this.info.list_date = "";
         this.info.sid = 0;
       }else{
         this.info = info;
       }
    }
    this.kline = null;
    this.predict = 0
    this.info.predict = ""
    if(this.info.indu == "大盘指数"){
      this.type = "index";
    }else{
      this.type = "stock";
    }
  }
  isRealSymbol(){
    if (this.info.symbol == "") return false;
    return true;
  }
  setKline(kl){
    this.kline = kl;
    var predict = this.getBarofDay(-1)
    var today = this.getBarofDay(-2)
    if (predict.close - today.close > 0)
    {
      this.predict = 1
      this.info.predict = "看涨"
      this.info.color = '#ff2f2f'
    }
    else {
      this.predict = -1
      this.info.predict = "看跌"
      this.info.color = '#4cda64'
    }
  }
  predict(page){
    var that = this
    var param = {symbol:that.info.symbol};
    fetch("predict",param).then((resolve)=>{
      console.log(resolve.data)
      if (resolve.code == 0) {
          // 成功
          that.setKline(resolve.data);
    }else{
      console.log("error code return:"+resolve.code)
    }},
  (err)=>{
    console.log(err);
  }).then(()=>{
    if (page)
      page.updateView()
  });
  }
}
