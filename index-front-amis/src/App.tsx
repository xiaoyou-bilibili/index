import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import './App.css'
import Index from "./views";
import Code from "./views/code";
import View from "./views/view";
import {
    FormOutlined,
    FileTextOutlined,
    ControlOutlined
} from '@ant-design/icons';
import { Layout, Menu } from 'antd';
import React, {useCallback, useEffect, useState} from 'react';
import {getManagesView} from "./api/api";
import Manage from "./views/manage";
const { Sider, Content } = Layout;

function App() {

    const getMenu = (data:{id:number,name:string}[]) => {
        let child:any[] = []
        data.forEach(item=>{
            child.push({key: item.id.toString(), label: <Link to={`/manage/${item.id}`}>{item.name}</Link>})
        })
        return [
            {
                key: 'index',
                icon: <FileTextOutlined />,
                label: <Link to="/">页面编辑器</Link>,
            },
            {
                key: 'code',
                icon: <FormOutlined />,
                label:  <Link to="/code">代码编辑器</Link>,
            },
            {
                key: 'manage',
                icon: <ControlOutlined />,
                label: "管理",
                children: child,
            },
        ]
    }

    let [items,setItems] = useState(getMenu([]))
    // 初始化获取所有页面信息
    const viewInit = useCallback(() => {
        // 获取所有的子菜单
        getManagesView().then((data:any )=>setItems(getMenu(data)))
    }, [])
    useEffect(() => viewInit(), [viewInit])


    return (
    <Router>
      <Layout>
          <Sider collapsible theme="light">
              <Menu
                  theme="light"
                  mode="inline"
                  defaultSelectedKeys={['1']}
                  items={items}
              />
          </Sider>
          <Layout className="site-layout">
              <Content
                  className="site-layout-background"
                  style={{margin: '12px 12px', minHeight: 700, background: "#ffffff"}}
              >
                  <Routes>
                      <Route path="/" element={<Index />} />
                      <Route path="/code" element={<Code />} />
                      <Route path="/view" element={<View />} />
                      <Route path="/manage/:id" element={<Manage />} />
                  </Routes>
              </Content>
          </Layout>
      </Layout>
    </Router>
  )
}

export default App
