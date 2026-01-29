import React, {useMemo} from 'react';
import {Breadcrumb, Menu, theme} from 'antd';
import {getBreadcrumbs, MenuItem, useChildMenuItems, useMenuItems} from "../../routers/router";
import {useTranslation} from "react-i18next";
import {useLocation, useNavigate} from "react-router-dom";
import Sider from "antd/es/layout/Sider";

/**React
 * 侧边栏
 * @constructor
 */
const MySider: React.FC = () => {

    let {pathname} = useLocation();
    console.log("MySider pathname: ", pathname);

    const {
        token: {colorBgContainer},
    } = theme.useToken();

    const navigate = useNavigate();
    const menuItems: MenuItem[] | null = useMenuItems(false);

    console.log("menuItems: ", menuItems);

    let selectMenuKeys: string[] = [];
    if (menuItems) {
        let items = menuItems;
        while (items) {
            let needMatchChildren = false;
            for (let menuItem of items  ) {
                if (pathname.startsWith(menuItem!.key + "/") || pathname === menuItem!.key) {
                    selectMenuKeys.push(menuItem!.key as string);
                    if (menuItem.children) {
                        needMatchChildren = true;
                        items = menuItem.children;
                        break;
                    }
                }
            }
            if (!needMatchChildren) {
                break;
            }
        }
    }

    function handlerItemClick(item: { key: string }) {
        navigate(item.key);
    }

    return (
        <>
            {menuItems && menuItems.length > 0 ? <Sider width={240} style={{background: colorBgContainer}}>
                <Menu
                    mode="inline"
                    defaultSelectedKeys={selectMenuKeys}
                    defaultOpenKeys={selectMenuKeys}
                    style={{height: '100%', color:"#777"}}
                    items={menuItems}
                    onClick={handlerItemClick}
                />
            </Sider> : <div/>}
        </>

    );
};


export default MySider;