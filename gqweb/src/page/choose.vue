<template>
  <div class="com-home">
    <div>
      <el-table :data="list"
      highlight-current-row
      @current-change="handleRow"
      style="width: 100%"
      max-height="250"
      :show-header=false>
        <el-table-column
          prop="symbol"
          label="代码">
        </el-table-column>
        <el-table-column
          prop="name"
          label="名称">
        </el-table-column>
        <el-table-column
          prop="indu"
          label="行业">
        </el-table-column>
        <el-table-column
          fixed="right"
          label="操作"
          v-if="onOpAdd">
          <template slot-scope="scope">
            <el-button
              @click.native.prevent="addRow(scope.$index, tableData)"
              type="text"
              size="small">
              增加
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
  import axios from 'axios'
  import router from '../router/router'
  export default{
    data: function () {
      return {
        list:[],
      }
    },
    created: function () {
      this.isLogined = this.GLOBAL.userInfo.isLogined;
      this.GLOBAL.hedgeScopes.fetch_net(this);
      //this.GLOBAL.hedgeScopes.initHedge();
      this.$store.commit('changeIndexConf', {
        isFooter: true,
        isSearch: false,
        isBack: false,
        title: '系统选股'
      })

    },
    methods: {
      updateView(){
				console.log("updateView()")
				this.list = this.GLOBAL.hedgeScopes.getRank();
			},
      handleColumn(tab,val){
        console.log(tab.name)
        this.activeName = tab.name;

      },
      handleRow(row, event) {
        console.log(row)
        this.onRowClick(row);
      }
    },
    addRow(row, event){
      this.GLOBAL.mystocks.addStock(row.symbol)
    }
  }
</script>

<style lang="scss">
  @import "../../static/css/home.scss";
</style>
