import Login from './views/Login.vue'
import NotFound from './views/404.vue'
import Home from './views/Home.vue'
import Leviatom from './views/nav1/Leviatom.vue'
import StreamNet from "./views/nav1/StreamNet";
import NetCoin from "./views/nav1/NetCoin";
import DevOps from "./views/nav2/DevOps";
import Deployment from "./views/nav2/Deployment";
import Experiment from "./views/nav2/Experiment";
import Addition from "./views/Addition"
import Register from "./views/Register"
import UserManagement from "./views/management/UserManagement"

let routes = [
    {
        path: '/login',
        component: Login,
        name: '',
        hidden: true
    },
    {
        path: '/register',
        component: Register,
        name: '',
        hidden: true
    },
    {
        path: '/addition',
        component: Addition,
        name: '',
        hidden: true
    },
    {
        path: '/404',
        component: NotFound,
        name: '',
        hidden: true
    },
    {
        path: '/',
        component: Home,
        name: 'Home',
        hidden: true
    },
    {
        path: '/user',
        component: Home,
        name: 'Functioin List',
        iconCls: 'el-icon-share',
        children: [
            {path: "/user/leviatom", component: Leviatom, name: 'Leviatom'},
            {path: "/user/streamNet", component: StreamNet, name: "StreamNet"},
            // {path: "/user/netCoin", component: NetCoin, name: "NetCoin"},
        ],
        rootName: "user"
    },
    // {
    //     path: '/server',
    //     component: Home,
    //     name: 'Servers Platform',
    //     iconCls: 'el-icon-s-platform',
    //     children: [
    //         {path: "/server/devOps", component: DevOps, name: "DevOps"},
    //         {path: "/server/deployment", component: Deployment, name: "Deployment"},
    //         {path: "/server/experiment", component: Experiment, name: "Experiment"},
    //         {path: "/server/userManagement", component: UserManagement, name: "UserManagement"},
    //     ],
    //     rootName: "server"
    // },
    {
        path: '/management',
        component: Home,
        name: 'Management',
        iconCls: 'el-icon-s-platform',
        children: [
            {path: "/management/userManagement", component: UserManagement, name: "UserManagement"},
        ],
        rootName: "management"
    },
    {
        path: '*',
        hidden: true,
        redirect: {path: '/404'}
    }
];

export default routes;
