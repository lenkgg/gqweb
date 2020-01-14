<template>
	<div class="index">
		<div class="center">
			<klcanvas class = 'wnd' :symbol='cvsSymbol' :kline='cvsKline' @click="onRefresh()"></klcanvas>
		</div>
		<input name = "sc" class = "input" type="text" placeholder="输入股票代码或名称查询"
			placeholder-class="edit-txt" autofocus="true" bindinput="onSearch"></input>
			<div class="symbol-head" v-show="isActive">
				<p style="height: 200rpx;width:100%;"></p>
				<p class="symbol-item">{{search.symbol}}</p>
				<p class="symbol-item">{{search.name}}</p>
				<p class="symbol-item">{{search.indu}}</p>
				<p class="symbol-item">{{search.markt}}</p>
				<p class="symbol-item">{{search.area}}</p>
				<p class="symbol-item">{{search.list_date}}</p>
			</div>
		<tablist :list="list" :onRowClick="onRow" :onTabClick="onTab" :onDelRow="onDelRow" ></tablist>
	</div>
</template>

<script>
	import Canvas from '../components/canvas'
	import Tablist from '../components/tablist.vue'
	import axios from 'axios'

	var cncode = require('../utils/stock_code_cn')

	export default{
		data: function () {
			return {
				data: 'index',
				num: 1,
				isActive: false,
				symbol: {},
				search:{},

				selcode: 'sh000001',
				selTab:'index',

				index:[],
				industry:[],
				concept:[],
				mine:[],
				advise:[],

				list:[]
			}
		},
		created () {
			this.GLOBAL.defaultIndexs.fetch_net(this);
			this.GLOBAL.mystocks.fetch_net(this);

			this.isActive = false
			this.calc();
			// head的修改
			this.$store.commit('changeIndexConf', {
				isFooter: true,
				isSearch: true,
				isBack: false,
				title: ''
			})
		},
		components: {
			klcanvas: Canvas,
			tablist: Tablist
		},
		computed: {
			cvsSymbol: function(){
				if (this.symbol && this.symbol.info)
					return this.symbol.info.symbol
				else {
					return ""
				}
			},
			cvsKline: function(){
				if (this.symbol)
					return this.symbol.kline;
				else {
					return []
				}
			},
		},
		watch: {
			'$route' (to, from) {
			}
		},
		methods: {
			updateView(){
				console.log("updateView()")
				this.calc();
			},
			onTab(activeName){
				//console.log(activeName)
				this.selTab = activeName
				this.calc();
			},
			onRow(row){
				console.log(row)
				this.selcode = row.symbol
				this.calc();
				//this.onRefresh();
			},
			onDelRow(row){
				this.GLOBAL.MyStocks.delStock(row.symbol);
				this.calc();
			},
			onRefresh(){
				//this.symbol = this.list.getSymbol(this.selcode)
				//this.$forceUpdate()
				console.log(this.selcode)
			},
			onSearch:function(e){
				if (e.detail.value == ""){
					isActive= false
					return
				}
	      var reg = new RegExp("/*" + e.detail.value +"/*","i");
	      for (let i of cncode){
	        if (i.symbol.match(reg) ||i.name.match(reg))
	          {
	            this.search = i;
							isActive = true
	            break;
	          }
	      }
			},
			calc: function(){
				console.log("selTab:"+this.selTab)
				if (this.selTab == "index")
					{
						this.index = this.GLOBAL.defaultIndexs.getSymbolCodeList();
						this.symbol = this.GLOBAL.defaultIndexs.getSymbol(this.selcode);
						this.list = this.index;
					}
				else if (this.selTab == "industry")
					this.list = this.listIndustry;
				else if (this.selTab == "concept")
					this.list = this.listConcept;
				else if (this.selTab == "mine")
					{
						this.mine = this.GLOBAL.mystocks.getSymbolCodeList();
						this.symbol = this.GLOBAL.mystocks.getSymbol(this.selcode)
						this.list = this.mine;
					}
			}
		}
	}

</script>
<style lang="scss">
	@import "../../static/css/index.scss";
	.wnd {
		width: 500px;
		height: 256px;
	}
</style>
