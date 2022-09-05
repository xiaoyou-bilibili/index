import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css';

// antd
import 'antd/dist/antd.css';
import AmisRegister from "./amis";
// 引入amis的css样式
import 'amis/lib/themes/cxd.css';
import 'amis/lib/helper.css';
import 'amis/sdk/iconfont.css';
// 注册自定义的组件和各种事件
AmisRegister()

ReactDOM.createRoot(document.getElementById('root')!).render(
  <App />
)

postMessage({ payload: 'removeLoading' }, '*')
