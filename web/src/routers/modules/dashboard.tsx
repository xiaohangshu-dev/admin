import { HomeOutlined,DesktopOutlined ,FundViewOutlined  ,MonitorOutlined } from "@ant-design/icons";
import lazyLoad from "../lazyLoad";
import React, { lazy } from "react";
import { MenuRouteObject } from "../router";

const dashboard: MenuRouteObject = {
    path: "index",
    label: "menu.dashboard.name",
    icon: <HomeOutlined />,
    element: lazyLoad(lazy(() => import("../../pages/Dashboard"))),
    // children: [
    //     {
    //         path: "index",
    //         label: "menu.dashboard.workplace",
    //         icon: <DesktopOutlined />,
    //         element: lazyLoad(lazy(() => import("../../pages/Dashboard"))),
    //     },
    //     {
    //         path: "index",
    //         label: "menu.dashboard.analysis",
    //         icon: <FundViewOutlined />,
    //         element: lazyLoad(lazy(() => import("../../pages/Dashboard"))),
    //     },
    //     {
    //         path: "index",
    //         label: "menu.dashboard.monitor",
    //         icon: <MonitorOutlined />,
    //         element: lazyLoad(lazy(() => import("../../pages/Dashboard"))),
    //     },
    // ],

}

export default dashboard