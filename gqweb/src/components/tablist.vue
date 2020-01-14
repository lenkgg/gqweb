<template>
  <div>
    <span style="width:80%;">
      <el-card>
        <el-tabs :active-name="activeName" @tab-click="handleColumn">
          <el-tab-pane name="index" label="指数">指数</el-tab-pane>
          <el-tab-pane name="industry" label="行业">行业</el-tab-pane>
          <el-tab-pane name="concept" label="概念">概念</el-tab-pane>
          <el-tab-pane name="mine" label="自选">自选</el-tab-pane>
        </el-tabs>
      </el-card>
    </span>
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
        prop="predict"
        label="测算">
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        v-if="onOpDel">
        <template slot-scope="scope">
          <el-button
            @click.native.prevent="deleteRow(scope.$index, tableData)"
            type="text"
            size="small">
            移除
          </el-button>
        </template>
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
</template>

<script>

export default{
  props:{
    list:{
      type: Array,
      default: ()=>{return []}
    },
    onRowClick:{
      type: Function,
      default: null
    },
    onTabClick:{
      type:Function,
      default:null
    },
    onOpDel:{
      type:Function,
      default:null
    },
    onOpAdd:{
      type:Function,
      default:null
    }
  },
  data: function () {
    return {
      activeName: 'index',
  }},
  methods:{
    deleteRow(index, rows) {
      //rows.splice(index, 1);
      console.log(rows)
      this.onOpDel(rows[index])
    },
    addRow(index, rows) {
      //rows.splice(index, 1);
      console.log(rows)
      this.onOpAdd(rows[index])
    },
    handleColumn(tab,val){
      console.log(tab.name)
      this.activeName = tab.name;
      this.onTabClick(tab.name)
    },
    handleRow(tab, event) {
      console.log(tab)
      this.onRowClick(tab);
    }
  },
  watch:{

  }
}
</script>

<style lang='scss'>
.box-card{
  width: 100%;
  display:flex;
  white-space: nowrap;
}
.box-card .el-tabs__item.is-top:last-child{
  color: #ff4949;
}
.edit-button{
  width: 18%;
  align-content: flex-end;
  align-items: flex-end;
}
</style>
