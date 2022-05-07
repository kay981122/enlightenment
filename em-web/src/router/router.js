const routes = [
    {
        name:'LoginItem',
        path:"/login",
        component:() => import('@/components/base/Login')
    },
    {
        name:'LoginItem',
        path:"/",
        component:() => import('@/components/base/Login')
    },
];
export default routes