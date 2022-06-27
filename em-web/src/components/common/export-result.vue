<template>
  <div>
      <el-dialog v-model="show" title="导出结果" width="40%" :before-close="handleClose">
            <el-table :data="tableData" style="width: 100%">
              <el-table-column prop="index" label="序号" />
                <el-table-column prop="status" label="状态">
                  <template #default="scope">
                      <el-tag  v-if="scope.row.status == '0'" type="info">初始化中</el-tag>
                      <el-tag  v-else-if="scope.row.status == '1'" type="warning">生成中</el-tag>
                      <el-tag  v-else-if="scope.row.status == '2'"  type="success">成功</el-tag>
                      <el-tag  v-else type="danger">失败</el-tag>
                  </template>
                </el-table-column>
                <el-table-column prop="createTime" label="创建时间" />
                <el-table-column prop="updateTime" label="更新时间" />
                <el-table-column prop="" label="操作" >
                    <template #default="scope">
                        <el-button @click="download(scope.row.id)" :disabled="scope.row.status !== '2'" type="primary">下载</el-button>
                    </template>
                </el-table-column>
            </el-table>
      </el-dialog>
  </div>
</template>

<script>
export default {
  components: {

  },
  data() {
    return {
      show:false,
      tableData : [],
    }
  },
  methods: {
    open() {
      this.show = true;
      this.queryResult();
    },
    handleClose(){
      this.show = false
    },
    queryResult(){
        this.$http.post("/domain/downloadDomainCSV",{
          "module":"domain",
          "userId":"00000"
        }).then((res)=>{
        if (res.code != 200) {
            this.$ElMessage({ type: "error", message: "查询失败！" });
            return;
        }
          this.tableData = res.data;
          let i = 1;
          this.tableData.forEach(e =>{
            e['index'] = i++;
          })
      })
    },
    download(id){
      console.log(id);
    }
  },
  mounted () {

  },
}

</script>
<style lang='scss' scoped>

</style>