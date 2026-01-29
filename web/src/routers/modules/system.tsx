import { AppstoreOutlined, IdcardOutlined, MenuOutlined, SettingOutlined, UserOutlined, } from "@ant-design/icons";
import lazyLoad from "../lazyLoad";
import React, { lazy } from "react";
import { MenuRouteObject } from "../router";

const system: MenuRouteObject = {
    path: "system",
    label: "menu.system.name",
    icon: <SettingOutlined />,
    // element: lazyLoad(lazy(() => import("../../pages/Dashboard"))),
    children: [
        {
            path: "user",
            label: "menu.system.user.name",
            icon: <UserOutlined />,
            element: lazyLoad(lazy(() => import("../../pages/System/User"))),
        },
        {
            path: "role",
            label: "menu.system.role.name",
            icon: <IdcardOutlined />,
            element: lazyLoad(lazy(() => import("../../pages/System/Role"))),
        },
        {
            path: "permission",
            label: "menu.system.permission.name",
            icon: <MenuOutlined />,
            element: lazyLoad(lazy(() => import("../../pages/System/Permission"))),
        },
        {
            path: "application",
            label: "menu.system.application.name",
            icon: <AppstoreOutlined />,
            element: lazyLoad(lazy(() => import("../../pages/System/Application"))),
        },
    ],

}

export default system