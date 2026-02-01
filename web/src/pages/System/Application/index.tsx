import React from 'react';
import welcome from "../../../assets/images/welcome.png";
import { Button, Input, Table } from 'antd';
import { useTranslation } from 'react-i18next';
import { AppstoreAddOutlined, DeleteOutlined, PlusOutlined, ReloadOutlined, SearchOutlined, TableOutlined, UploadOutlined } from '@ant-design/icons';
import styles from "./user.module.scss";


const Application: React.FC = () => {
    const { t } = useTranslation();
    const addText = t("menu.common.add");
    const importText = t("menu.common.import");
    const deleteText = t("menu.common.delete");
    const refreshText = t("menu.common.refresh");
    const search_placeholder = t("menu.system.application.search_placeholder");


    const dataSource = [
        {
            key: '1',
            name: '胡彦斌',
            age: 32,
            address: '西湖区湖底公园1号',
        },
        {
            key: '2',
            name: '胡彦祖',
            age: 42,
            address: '西湖区湖底公园1号',
        },
    ];

    const columns = [
        {
            title: '姓名',
            dataIndex: 'name',
            key: 'name',
        },
        {
            title: '年龄',
            dataIndex: 'age',
            key: 'age',
        },
        {
            title: '住址',
            dataIndex: 'address',
            key: 'address',
        },
    ];

    return (
        <>
            <div className={styles.operate}>

                <div className={styles.user}>
                    <Button type='default' icon={<ReloadOutlined />} style={{ background: "#40485b", color: "#fff" }} >{refreshText}</Button>
                    <Button type="primary" icon={<PlusOutlined />}>{addText}</Button>
                    <Button type="primary" danger icon={<DeleteOutlined />}>{deleteText}</Button>
                </div>
                <div className={styles.page}>
                    <Input placeholder={search_placeholder} />
                    <Button type='default' icon={<TableOutlined />} />
                    <Button type='default' icon={<SearchOutlined />} />
                </div>
            </div>
            <Table
                dataSource={dataSource}
                columns={columns}
                bordered
            />

        </>
    );
};

export default Application;
