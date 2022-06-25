<template>
  <div>
    <div class="search-box">
      <el-form :inline="true" :model="searchData">
        <el-form-item label="域名">
          <el-input v-model="searchData.domain" placeholder="domain" />
        </el-form-item>
        <el-form-item label="标题">
          <el-input v-model="searchData.title" placeholder="title" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchData.status" placeholder="status" clearable>
            <el-option label="0" value="0" />
            <el-option label="1" value="1" />
            <el-option label="2" value="2" />
            <el-option label="3" value="3" />
            <el-option label="4" value="4" />
          </el-select>
        </el-form-item>
        <el-form-item label="注册时间">
          <el-date-picker
            v-model="searchData.registerDate"
            type="daterange"
            unlink-panels
            range-separator="To"
            start-placeholder="Start date"
            end-placeholder="End date"
            :shortcuts="shortcuts"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        <el-form-item label="更新时间">
          <el-date-picker
            v-model="searchData.updateDate"
            type="daterange"
            unlink-panels
            range-separator="To"
            start-placeholder="Start date"
            end-placeholder="End date"
            :shortcuts="shortcuts"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="query">查询</el-button>
          <el-button type="primary" @click="reset">重置</el-button>
          <el-button type="primary" @click="reset">导出</el-button>
          <el-button type="primary" @click="reset">导出结果</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="table-box">
      <el-scrollbar>
        <el-table :data="tableData" style="width: 100%;line-height: normal;" ref="tableData" v-loading="loading">
          <el-table-column prop="domain" label="域名" />
          <el-table-column prop="title" label="标题" />
          <el-table-column prop="registerDate" label="注册时间" />
          <el-table-column prop="updateDate" label="更新时间" />
          <el-table-column prop="status" label="状态" />
        </el-table>
      </el-scrollbar>
           <div class="table-box-bottom">
          <el-pagination
            background
            layout="total, sizes, prev, pager, next, jumper"
            :page-sizes="[10, 50, 100, 200, 400]"
            v-model:currentPage="page"
            v-model:page-size="pageSize"
            :total="total"
            @current-change="currentChange"/>
     </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "HomeItem",
  data() {
    return {
      tableData: [],
      searchData: {
        domain: "",
        title: "",
        status: "",
        registerDate: "",
        registerBeginTime: "",
        registerEndTime: "",
        updateDate: "",
        updatebeginTime: "",
        updateEndTime: "",
      },
      total: 0,
      pageSize: 10,
      page: 1,
      loading:false,
      shortcuts: [
        {
          text: "Last week",
          value: () => {
            const end = new Date();
            const start = new Date();
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
            return [start, end];
          },
        },
        {
          text: "Last month",
          value: () => {
            const end = new Date();
            const start = new Date();
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
            return [start, end];
          },
        },
        {
          text: "Last 3 months",
          value: () => {
            const end = new Date();
            const start = new Date();
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
            return [start, end];
          },
        },
      ],
    };
  },
  methods: {
    query() {
      if (this.searchData.registerDate != "") {
        this.searchData.registerBeginTime = this.searchData.registerDate[0];
        this.searchData.registerEndTime = this.searchData.registerDate[1];
      }
      if (this.searchData.updateDate != "") {
        this.searchData.updatebeginTime = this.searchData.updateDate[0];
        this.searchData.updateEndTime = this.searchData.updateDate[1];
      }
      this.loading = true;
      this.$http
        .post("/domain/getDomainList", {
          domain: this.searchData.domain,
          title: this.searchData.title,
          status: this.searchData.status,
          page: this.page,
          pageSize: this.pageSize,
          registerBeginTime: this.searchData.registerBeginTime,
          registerEndTime: this.searchData.registerEndTime,
          updatebeginTime: this.searchData.updatebeginTime,
          updateEndTime: this.searchData.updateEndTime,
        })
        .then((res) => {
          if (res.code == 200) {
            this.tableData = res.data.list;
            this.total = res.data.total;
          } else {
            this.$ElMessage({ type: "error", message: "请求失败！" });
          }
            this.loading = false;
        })
        .catch((err) => {
          console.log(err);
          this.loading = false;
        });
    },
    currentChange() {
      this.query();
    },
    reset() {
      for (let key in this.searchData) {
        this.searchData[key] = "";
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.search-box {
  padding: 24px 24px 2px;
  background-color: #fff;
  border-radius: 2px;
  margin-bottom: 12px;
}
.table-box {
  padding: 24px;
  background-color: #fff;
  border-radius: 2px;
  height:600px;
}
.table-box-bottom {
  display: flex;
  justify-content: flex-end;
}
.el-pagination {
  padding: 20px 0;
}
</style>