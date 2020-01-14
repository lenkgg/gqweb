var MyHedge = require('./myhedge')
var api = require('../api')

export default class Scopes {
  constructor(props){
      this.hedges = [];
      this.myPage = null;
      Object.assign(this, props);
      return this;
    }

      initHedge(vip){
          /*if (this.getHedge("CN_HS300","sz399300") == null)
            api.getDefaultHedge();*/
          if (this.getHedge("CN_ALL","sz399300") == null) ///&& wepy.$instance.globalData.userInfo.vip >= 1
              api.hedge("CN_ALL","sz399300")
      }
      /////bind page view
      setPage(page){
        this.myPage = page;
      }
      getHedge(){
        //default hedge_all
        scp = "CN_ALL";
        bsym = "sz399300"
        ///////////////
        for (var i=0; i<this.hedges.length; i++){
          if (this.hedges[i].scope == scp && this.hedges[i].bsymbol == bsym) //.toLocaleUpperCase())
            return this.hedges[i]
        }
        return null
      }
      getRank(){
        //default hedge_all
        scp = "CN_ALL";
        bsym = "sz399300"
        /////
          var hg = this.getHedge(scp, bsym)
          if (hg!=null)
            return hg.rank;
          return null;
      }

      getRankItem(index){
        //default hedge_all
        scp = "CN_ALL";
        bsym = "sz399300"
        /////
        var hg = this.getHedge(scp, bsym)
        if (hg != null){
          if (index >=0 && index < hg.rank.length)
             return hg.rank[index]
        }
        return null
      }

      fetch_net(page){
        var that = this
        var param = {bsymbol:"sz399300", scope:"CN_ALL"};
        fetch("hedge", param)
        .then(function(resolve){
          console.log(resolve);
          if (resolve.code == 0){
            var hg = that.getHedge(data.scope, data.bsymbol)
            if (hg != null)
              hg.onUpdateHedge(data)
            else{
              var hg = new MyHedge();
              hg.init(data.scope, data.bsymbol)
              //console.log(data[i])
              hg.onUpdateHedge(data)
              that.hedges.push(hg)
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
}

export function CreateScopes(){
  return new Scopes()
}
