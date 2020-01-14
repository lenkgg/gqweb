module.exports = function KDict(){
/*字典 Dictionary类*/
return {
    datastore : [
       {
      symbol: "",
      data: [],
    }
    ],

 add:function(key, value) {
   var it = [];
   for (var i in value){
     it.push(value[i])
   }
   for (var k in this.datastore){
      if (this.datastore[k].symbol == key){
          this.datastore[k].data = it;
          return ;
        }
      }
    var item  = {};
    item.symbol = key;
    item.data = it;
    this.datastore.push(item);
  },

find:function(key) {
  for (var k in this.datastore){
     if (this.datastore[k].symbol == key){
         return this.datastore[k].data;
       }
     }
},

remove:function(key) {
  for (var k in this.datastore){
    if (this.datastore[k].symbol == key){
      delete this.datastore[k];
    }
  }
},

showAll:function() {
    var str = "";
    for(var key in this.datastore) {
        str += this.datastore[key].symbol + " -> " + this.datastore[key].data + ";  "
    }
    console.log(str);
},

count:function() {
    /*var ss = Object.keys(this.datastore).length;
    console.log("ssss   "+ss);
    return Object.keys(this.datastore).length;*/
    /**/
    var n = 0;
    for(var key in Object.keys(this.datastore)) {
        ++n;
    }
    //console.log(n);
    return n;
},

clear:function() {
    for(var key in this.datastore) {
        delete this.datastore[key];
    }
}
}
}
