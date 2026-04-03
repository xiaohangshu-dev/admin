import React from 'react';
import welcome from "../../../assets/images/welcome.png";
import { Button, Flex, Input, Space, Table } from 'antd';
import { useTranslation } from 'react-i18next';
import { DeleteOutlined, PlusOutlined, ReloadOutlined, SearchOutlined, TableOutlined, UploadOutlined } from '@ant-design/icons';
import styles from "./user.module.scss";


const User: React.FC = () => {
    const { t } = useTranslation();
    const addText = t("menu.common.add");
    const importText = t("menu.common.import");
    const deleteText = t("menu.common.delete");
    const search_placeholder = t("menu.system.user.search_placeholder");
    const refreshText = t("menu.common.refresh");



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
            <Space vertical size={10} style={{ width: "100%" }}>
                <Flex justify="space-between">
                    <Space size={'small'}>
                        <Button type='default' icon={<ReloadOutlined />} style={{ background: "#40485b", color: "#fff" }} > {refreshText}</Button>
                        <Button type="primary" icon={<PlusOutlined />}>{addText}</Button>
                        <Button type="default" icon={<UploadOutlined />}>{importText}</Button>
                        <Button type="primary" danger icon={<DeleteOutlined />}>{deleteText}</Button>
                    </Space>
                    <Space size={'small'}>
                        <Input placeholder={search_placeholder} />
                        <Button type='default' icon={<TableOutlined />} />
                        <Button type='default' icon={<SearchOutlined />} />
                    </Space>
                </Flex>
                <Table
                    dataSource={dataSource}
                    columns={columns}
                    bordered
                />
            </Space>
        </>
    );
};

export default User;
