export let tabVuex = {
  state: {
    // 所有打开的路由
    openTab: [{
      route: '/home/main',
      name: '主页'
    }],
    // 默认激活路径；激活状态
    activeIndex: '/home/main'
  },
  mutations: {
    // 添加tabs
    addTabs(state, data) {
      state.openTab.push(data);
    },
    // 删除tabs
    deleteTabs(state, route) {
      let index = 0;
      for (let option of state.openTab) {
        if (option.route === route) {
          break;
        }
        index++;
      }
      state.openTab.splice(index, 1);
    },
    // 设置当前激活的tab
    setActiveIndex(state, index) {
      state.activeIndex = index;
    },
    initTabs(state) {
      state.openTab = []
      state.openTab.push({
        route: '/home/main',
        name: '主页'
      })
    }
  },
  actions: {

  },
  modules: {

  }
}