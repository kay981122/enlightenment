import {
  createRouter,
  createWebHistory
} from 'vue-router';
import routes from './router';

const router = createRouter({
  history: createWebHistory(),
  routes
})
router.beforeEach((to, from, next) => {
  let token = localStorage.getItem('token');
  if (token) {
    next()
  } else {
    if (to.path === '/login' || to.path === '/register') {
      next()
    } else {
      next('/login')
    }
  }
})
export default router