import React, { useState } from 'react';
import welcome from "../../../assets/images/welcome.png";
import { Button, Input, Table, TableColumnsType, TableProps } from 'antd';
import { useTranslation } from 'react-i18next';
import { CaretDownOutlined, DeleteOutlined, PlusOutlined, ReloadOutlined, TableOutlined } from '@ant-design/icons';
import styles from "./role.module.scss";


const Role: React.FC = () => {
    const { t } = useTranslation();
    const addText = t("menu.common.add");
    const deleteText = t("menu.common.delete");
    const search_placeholder = t("menu.system.role.search_placeholder");
    const expandText = t("menu.common.expand");
    const collapseText = t("menu.common.collapse");
    const refreshText = t("menu.common.refresh");


    type TableRowSelection<T extends object = object> = TableProps<T>['rowSelection'];

    interface DataType {
        key: React.ReactNode;
        name: string;
        age: number;
        address: string;
        children?: DataType[];
    }

    const columns: TableColumnsType<DataType> = [
        {
            title: 'Name',
            dataIndex: 'name',
            key: 'name',
        },
        {
            title: 'Age',
            dataIndex: 'age',
            key: 'age',
            width: '12%',
        },
        {
            title: 'Address',
            dataIndex: 'address',
            width: '30%',
            key: 'address',
        },
    ];

    const data: DataType[] = [
        {
            key: 1,
            name: 'John Brown sr.',
            age: 60,
            address: 'New York No. 1 Lake Park',
            children: [
                {
                    key: 11,
                    name: 'John Brown',
                    age: 42,
                    address: 'New York No. 2 Lake Park',
                },
                {
                    key: 12,
                    name: 'John Brown jr.',
                    age: 30,
                    address: 'New York No. 3 Lake Park',
                    children: [
                        {
                            key: 121,
                            name: 'Jimmy Brown',
                            age: 16,
                            address: 'New York No. 3 Lake Park',
                        },
                    ],
                },
                {
                    key: 13,
                    name: 'Jim Green sr.',
                    age: 72,
                    address: 'London No. 1 Lake Park',
                    children: [
                        {
                            key: 131,
                            name: 'Jim Green',
                            age: 42,
                            address: 'London No. 2 Lake Park',
                            children: [
                                {
                                    key: 1311,
                                    name: 'Jim Green jr.',
                                    age: 25,
                                    address: 'London No. 3 Lake Park',
                                },
                                {
                                    key: 1312,
                                    name: 'Jimmy Green sr.',
                                    age: 18,
                                    address: 'London No. 4 Lake Park',
                                },
                            ],
                        },
                    ],
                },
            ],
        },
        {
            key: 2,
            name: 'Joe Black',
            age: 32,
            address: 'Sydney No. 1 Lake Park',
        },
    ];

    // rowSelection objects indicates the need for row selection
    const rowSelection: TableRowSelection<DataType> = {
        onChange: (selectedRowKeys, selectedRows) => {
            console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
        },
        onSelect: (record, selected, selectedRows) => {
            console.log(record, selected, selectedRows);
        },
        onSelectAll: (selected, selectedRows, changeRows) => {
            console.log(selected, selectedRows, changeRows);
        },
    };

    const [checkStrictly, setCheckStrictly] = useState(false);

    return (
        <>
            <div className={styles.operate}>

                <div className={styles.role}>
                    <Button type='default' icon={<ReloadOutlined />} style={{ background: "#40485b", color: "#fff" }} >{refreshText}</Button>
                    <Button type="primary" icon={<PlusOutlined />}>{addText}</Button>
                    <Button type="primary" danger icon={<DeleteOutlined />}>{deleteText}</Button>
                    <Button type="primary" danger icon={<CaretDownOutlined />} style={{ background: "purple" }}>{expandText}</Button>
                </div>

                <div className={styles.page}>
                    <Input placeholder={search_placeholder} />
                    <Button type='default' icon={<TableOutlined />} />
                </div>

            </div>

            <Table<DataType>
                columns={columns}
                rowSelection={{ ...rowSelection, checkStrictly }}
                dataSource={data}
            />


        </>
    );
};

export default Role;
