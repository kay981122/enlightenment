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
        component:() => import('@/components/base/Index')
    }
];
export default routes