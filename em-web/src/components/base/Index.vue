<template>
    <el-container class="layout-container-demo">
      <el-header style="text-align: right; font-size: 12px">
        <div class="toolbar">
          <el-dropdown>
            <el-icon style="margin-right: 8px; margin-top: 1px"><setting/></el-icon>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>View</el-dropdown-item>
                <el-dropdown-item>Add</el-dropdown-item>
                <el-dropdown-item>Delete</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
          <span>Tom</span>
        </div>
      </el-header>
      <el-container>
     <el-aside :width="isCollapse ? '64px' : '15%'">
      <el-scrollbar>
        <el-button v-model="isCollapse" @click="changeCollapse">| | |</el-button>  
            <el-menu default-active="2" class="el-menu-vertical" :collapse="isCollapse" :collapse-transition = "false">
            <el-menu-item index="1">
            <el-icon><Burger /></el-icon>
            <template #title>首页</template>
            </el-menu-item>
            <el-sub-menu index="2">
            <template #title>
                <el-icon><setting /></el-icon>
                <span>系统管理</span>
            </template>
                <el-menu-item index="2-1">用户管理</el-menu-item>
                <el-menu-item index="2-2">权限管理</el-menu-item>
                <!-- <el-menu-item index="1-3">item three</el-menu-item> -->
            </el-sub-menu>
            <el-menu-item index="3">
            <el-icon><document /></el-icon>
            <template #title>项目管理</template>
            </el-menu-item>
        </el-menu>
      </el-scrollbar>
    </el-aside>
      <el-main>
        <el-scrollbar>
          <!-- <el-table :data="tableData">
            <el-table-column prop="date" label="Date" width="140" />
            <el-table-column prop="name" label="Name" width="120" />
            <el-table-column prop="address" label="Address" />
          </el-table> -->
        </el-scrollbar>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
    name:'IndexItem',
    data() {
        return {
            isCollapse:false
        }
    },
    methods: {
        changeCollapse(){
            this.isCollapse = !this.isCollapse;
        }
    }
}
</script>

<style lang="scss" scoped>
.layout-container-demo .el-header {
  position: relative;
  background-color: var(--el-color-primary-light-7);
  color: var(--el-text-color-primary);
}
.layout-container-demo .el-aside {
  color: var(--el-text-color-primary);
  background: var(--el-color-primary-light-8);
}
.layout-container-demo .el-menu {
  border-right: none;
}
.layout-container-demo .el-main {
  padding: 0;
}
.layout-container-demo .toolbar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  right: 20px;
}
.el-menu-vertical:not(.el-menu--collapse) {
      width: auto;
}

.layout-container-demo .el-aside {
    background-color: #fff;
}
</style>>
