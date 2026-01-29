import React from 'react';
import welcome from "../../assets/images/welcome.png";



const Dashboard: React.FC = () => {
    return (
        <div style={{textAlign:"center"}}>
            <img src={welcome} style={{maxWidth:"60%"}}/>
            <p>欢迎使用 react-develop-template</p>
        </div>
    );
};

export default Dashboard;
