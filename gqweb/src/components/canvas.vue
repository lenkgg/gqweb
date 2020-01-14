<template>
  <div>
  <div>
    <canvas id="kkline" class="stage" @click="onSymbol('{{symbol}}')"></canvas>
  </div>
  <div style="height: 20px;"></div>
  <div class = "no_wrap">
      <label class="text_nm">预测 {{t_nexday}}  </label>
      <label class="text_bold">{{t_symbol}}    </label>
      <label class='up_down' :style="c_change">{{t_change}}</label>
  </div>
</div>
</template>

<script>
var kl = require('../utils/wxChart/kline')
var common = require('../utils/common')
var util = require('../utils/util')

export default{
  props: {
    'symbol': {
      type: String,
      default: ''
    },
  'kline': {
    type: Array,
    default: ()=>{return []}
  }
},
  data: function(){
    return {
      predict:{},
      today:{},

      kLine: null, // 价格走势图对象
    }
  },
  watch: {
      kline(val,old) {
        console.log("enter watch kline.");
        this.kline = val
        this.today = this.getBarofDay(-2)
        this.predict = this.getBarofDay(-1)
        this.paintDataK()
      }
  },
  computed:{
    c_change(){
      if (!this.predict || !this.today)
        return ''
      var change = this.predict.close - this.today.close;
      //console.log(change)
      return  change >= 0 ? 'color: #ff2f2f;': 'color: #4cda64;';
    },
    t_change(){
      if (!this.predict || !this.today)
        return "-"
      var change = this.predict.close - this.today.close;
      return  change >= 0 ? '看涨':'看跌';
    },
    t_nexday(){
      var reg = /\-/g
      if ((this.predict) && (this.predict.date != null))
        return this.predict.date.replace(reg,".");
      return ""
    },
    t_symbol(){
        var info = util.getSymbolInfo(this.symbol)
        if (info != null)
          return info.name
        else
          return '';
    },
  },

  created () {
    this.today = this.getBarofDay(-2)
    this.predict = this.getBarofDay(-1)
    this.paintDataK()
  },
  methods:{
    getBarofDay: function(index){
      if (!this.kline)
        return null
      if (index < 0)
        index = this.kline.length + index;

      if (index < 0 || index >= this.kline.length)
        return null

        var rlt = {}
        var d = this.kline[index].split(',');
        rlt.date = d[0].split(' ')[0];
        rlt.open = d[1]/1; //开盘价
        rlt.high = d[2]/1; //最高价
        rlt.low = d[3]/1; //最低价
        rlt.close = d[4]/1; //现价
        return rlt;
    },
    // 绘制日K线图
    paintDataK: function() {
      //window.οnlοad=function(){
        console.log('enter paintDatak()')
        let canvas = document.getElementById('kkline');
        if (!canvas){
          console.log("canvas is null, exit.")
          return
        }else{
          console.log("canvas is ok.")
        }
        //let canvas = this.$refs.kkline;
        let ctx = canvas.getContext('2d');

        this.kLine = kl('kkline', ctx).init(common.getOptionKline1('ddk'))
        this.kLine.metaData1(this.kline, common.getOptionKline1('ddk'))
        this.kLine.draw()
      }
  }
}
</script>
<style lang="scss">
  .stage {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 160px;
  }
  .up_down {
    font-size: 28rpx;
    font-weight: bold;
  }
  .text_nm {
    size:18;
  }
  .text_bold {
    size:18;
    font-weight:bold;
  }
  .no_wrap {
    width:100%;
    overflow:hidden;
    white-space:nowrap;
    text-overflow:ellipsis;
  }
</style>
