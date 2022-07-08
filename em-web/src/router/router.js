const routes = [{
    path: "/",
    redirect: "login"
  },
  {
    name: 'login',
    path: "/login",
    component: () => import('@/components/modules/base/Login')
  },
  {
    name: 'home',
    path: "/home",
    component: () => import('@/components/modules/base/Home'),
    children: [{
        name: 'main',
        path: 'main',
        component: () => import('@/components/modules/base/Index')
      }, {
        name: 'user',
        path: 'user',
        component: () => import('@/components/modules/system/UserManager')
      },
      {
        name: 'permission',
        path: 'permission',
        component: () => import('@/components/modules/system/PermissionManager')
      },
      {
        name: 'domain',
        path: 'domain',
        component: () => import('@/components/modules/item/DomainManager')
      },
    ]
  },
];
export default routes