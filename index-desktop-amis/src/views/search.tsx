import { Input } from 'antd';
import React, { useState } from "react";
import { search } from "@/api/api";
import { render as RenderUI } from "amis-core";
import './search.css'

export default function Search() {
  // 输入框内容
  let [input,setInput] = useState('')
  // 自定义渲染
  let [ui,setUI] = useState({ "type": "page", "body": {}})
  const inputOnchange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { value } = e.target
    setInput(value)
    search(value).then((res:any) => {
      setUI({ "type": "page", "body": {}})
      // 延迟触发，避免重叠
      setTimeout(() => {setUI(res.view)}, 100)
    })
  }
  return (
    <div>
      <Input style={{margin: '5px', fontSize: '25px', width: '99%'}} placeholder="请输入关键词" value={input} onChange={inputOnchange} />
      <div style={{maxHeight: "535px", marginTop: "-10px", overflowY: "scroll"}}>{RenderUI(ui)}</div>
    </div>
  );
}
