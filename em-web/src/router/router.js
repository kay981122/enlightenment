const routes = [{
    path: "/",
    redirect: "login"
  },
  {
    name: 'Login',
    path: "/login",
    component: () => import('@/components/modules/base/Login')
  },
  {
    name: 'Home',
    path: "/home",
    component: () => import('@/components/modules/base/Home'),
    children: [{
        name: 'Main',
        path: 'main',
        component: () => import('@/components/modules/base/Index')
      }, {
        name: 'User',
        path: 'user',
        component: () => import('@/components/modules/system/UserManager')
      },
      {
        name: 'Permission',
        path: 'permission',
        component: () => import('@/components/modules/system/PermissionManager')
      },
      {
        name: 'Domain',
        path: 'domain',
        component: () => import('@/components/modules/item/DomainManager')
      },
    ]
  },
];
export default routes