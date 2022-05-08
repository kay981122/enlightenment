const routes = [
    {
        path:"/",
        redirect:"login"
    },
    {
        name:'LoginItem',
        path:"/login",
        component:() => import('@/components/base/Login')
    },
    {
        name:'indexItem',
        path:"/index",
        component:() => import('@/components/base/Index'),
        children:[
            {   
                name:'HomeItem',
                path:'/home',
                component:() => import('@/components/base/Home')
            },
            {   
                name:'PermissionManagerItem',
                path:'/permission',
                component:() => import('@/components/sys/PermissionManager')
            },
            {   
                name:'UserManagerItem',
                path:'/user',
                component:() => import('@/components/sys/UserManager')
            }
        ]
    },

];
export default routes