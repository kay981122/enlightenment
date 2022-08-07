<template>
  <el-container class="layout-container-demo">
    <el-header style="margin-bottom: 2px">
      <el-row style="font-size: 20px; height: 100%">
        <el-col :span="6" style="text-align: left">
          <span
            class="toolbar"
            style="
              font-size: 25px;
              font-weight: bolder;
              text-shadow: 5px 5px 5px #ff0000;">Enlightenment后台系统</span>
        </el-col>
        <el-col :span="18" style="text-align: right">
          <div class="toolbar">
            <el-dropdown>
              <el-icon style="margin-right: 8px; margin-top: 1px"><setting/></el-icon>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click="loginOut">登出</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
            <span>{{userInfo.nickname}}</span>
          </div>
        </el-col>
      </el-row>
    </el-header>
    <el-container>
      <el-aside :width="isCollapse ? '64px' : '10%'">
        <el-scrollbar>
          <div class="btn-center">
            <el-button v-model="isCollapse" @click="changeCollapse">| | |</el-button>
          </div>
          <el-menu class="el-menu-vertical" :collapse="isCollapse" :collapse-transition="false" router :default-active="$route.path">
            <el-menu-item index="/home/main">
              <el-icon><Burger /></el-icon>
              <template #title>主页</template>
            </el-menu-item>
            <el-sub-menu index="2">
              <template #title>
                <el-icon><setting /></el-icon>
                <span>系统管理</span>
              </template>
              <el-menu-item index="/home/user">用户管理</el-menu-item>
              <el-menu-item index="/home/permission">权限管理</el-menu-item>
            </el-sub-menu>
            <el-sub-menu index="3">
                <template #title>
                <el-icon><document /></el-icon>
                <span>项目管理</span>
              </template>
              <el-menu-item index="/home/domain">域名列表</el-menu-item>
            </el-sub-menu>
          </el-menu>
        </el-scrollbar>
      </el-aside>
      <el-main>
        <!-- 内容区 -->
        <div class="app-wrap">
          <!-- 此处放置el-tabs代码 -->
          <div>
            <el-tabs v-model="$store.state.tabVuex.activeIndex" type="border-card" closable v-if="$store.state.tabVuex.openTab.length" @tab-click='tabClick' @tab-remove='tabRemove'>
              <!-- 获取vuex中的openTab数组数据，循环展示 -->
              <el-tab-pane :key="index" v-for="(item, index) in $store.state.tabVuex.openTab"
              :label="item.name" :name="item.route">
                <div class="content-wrap">
                    <router-view v-slot="{ Component }">
                      <transition name="fade-transform" mode="out-in">
                        <keep-alive>
                          <component :is="Component" />
                        </keep-alive>
                      </transition>
                    </router-view>
                </div>
              </el-tab-pane>
            </el-tabs>
          </div>
        </div>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  name: "IndexItem",
  data() {
    return {
      isCollapse: false,
      userInfo:{}
    };
  },
  methods: {
    changeCollapse() {
      this.isCollapse = !this.isCollapse;
    },
    // tab标签点击时，切换相应的路由
    tabClick() {
      // console.log("tab",tab);
      // console.log('active',this.$store.state.tabVuex.activeIndex);
      this.$router.push({path:this.$store.state.tabVuex.activeIndex});
    },
    // 移除tab标签
    tabRemove(targetName) {
      // console.log("tabRemove",targetName);
      //首页不删
      if(targetName == '/home/main') {
        return
      }
      this.$store.commit('deleteTabs',targetName);
      if(this.$store.state.tabVuex.openTab && this.$store.state.tabVuex.openTab.length >= 1) {
        // console.log('=============',this.$store.state.tabVuex.openTab[this.$store.state.tabVuex.openTab.length-1].route)
        // 设置路由展示，为索引前一个路由
        this.$store.commit('setActiveIndex',this.$store.state.tabVuex.openTab[this.$store.state.tabVuex.openTab.length-1].route);
        //跳转路由
        this.$router.push({path: this.$store.state.tabVuex.activeIndex});
      }else {
        // 否则跳转到首页
        this.$router.push({path:'/home/main'})
      }
    },
    loginOut() {
      localStorage.clear()
      this.$router.push({
        path: "/login",
      });
    }
  },
  mounted() {
      this.$store.commit('setActiveIndex', this.$route.path);
      let userInfo = this.$route.params.userInfo
      if(userInfo) {
        this.userInfo = JSON.parse(userInfo);
        localStorage.setItem("userInfo",userInfo)
        // 新登录初始化菜单
        this.$store.commit('initTabs', {});
        return;
      }
        userInfo = localStorage.getItem("userInfo")
        if(userInfo) {
           localStorage.setItem("userInfo",userInfo)
           this.userInfo = JSON.parse(userInfo)
           return;
        }
      this.$router.push({
        path: "/login",
      });
  },
  created() {
      // 在页面加载时读取localStorage里的状态信息
      if (localStorage.getItem("data") ) {
          //replaceState替换数据 Object.assign合并对象
          this.$store.replaceState(Object.assign({}, this.$store.state,JSON.parse(localStorage.getItem("data"))))
      } 
      //在页面刷新时将vuex里的信息保存到localStorage里
      window.addEventListener("beforeunload",()=>{
          localStorage.setItem("data",JSON.stringify(this.$store.state))
      })
  },
  watch:{
    '$route'(to){
        //判断路由是否已经打开
        //已经打开的 ，将其置为active
        //未打开的，将其放入队列里
        let flag = false;
        for(let item of this.$store.state.tabVuex.openTab) {
          // console.log("item.path",item.route)
          // console.log("to.path",to.path)
          if(item.route === to.path) {
            // console.log('to.path',to.path);
            this.$store.commit('setActiveIndex',to.path)
            flag = true;
            break;
          }
        }
        if(!flag) {
          // console.log('to.path',to.path);
          // 通过路由的判断，来加入标签页的名称
          if(this.$route.path == '/home/domain') {
          this.$store.commit('addTabs',{route:this.$route.path,name:'域名列表'})
          }
          if(this.$route.path == '/home/permission') {
            this.$store.commit('addTabs',{route:this.$route.path,name:'权限管理'})
          }
          if(this.$route.path == '/home/user') {
            this.$store.commit('addTabs',{route:this.$route.path,name:'用户管理'})
          }
          this.$store.commit('setActiveIndex', this.$route.path);
        }
    }
  }
};
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
  height: 100vh;
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
  border-right: 1px rgb(198, 226, 255) solid;
}
.page-header {
  margin: 20px 10px;
}
.btn-center {
  text-align: center;
}
.content-wrap{
  height:100vh;
}
</style>>
