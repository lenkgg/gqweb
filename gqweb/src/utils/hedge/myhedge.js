var api = require('../api')
var util = require('../util')
module.exports = function(){
  return{
    scope:"",
    bsymbol:"",
    rank:[],
    htime: "",

    init: function(scp, bsym){
      this.scope = scp;
      this.bsymbol = bsym;   //.toLocaleUpperCase();
      this.rank = [];
    },
    clear: function(){
        this.rank = []
    },
    fetch_net: function(){
      if ((this.htime != "" &&
        Math.round(new Date().getTime()/1000) - Date.parse(new Date(this.htime))/1000 > 300)
        || this.rank.length == 0){
          api.hedge(this.scope, this.bsymbol);
        }
    },
    ///////on update data from svr//////
    onUpdateHedge: function(data){
      this.clear()

      this.htime = data.htime;
      var colors = ["#333333", "#4cd264", "#ff2f2f", "#f257e4", "#FFA500"]
      for (var i=0; i< data.rank.length; i++){
        var ll = data.rank[i].split(',')
        var info = util.getSymbolInfo(ll[0]);
        if(!info)
          continue;

        info.e_change = ll[1];
        var id = parseInt(info.e_change/2)
        if (id >4)
          id = 4
        info.color = colors[id]
        this.rank.push(info)
    }
  },


}}
