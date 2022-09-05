import { render as renderAmis } from 'amis';
import React, {useCallback, useEffect, useState} from 'react';
import {getCustomView} from "@/api/api";
import {useParams} from "react-router-dom";


export default function Page() {
  let [view,setView] = useState({type:"page",body:{}})
  let param = useParams()
  // 初始化获取所有页面信息
  useEffect(() => {
    getCustomView(param.name as string, param.id as string, param.view as string).then((data:any) => {
      setView(data.page)
    })
  }, [])

  return <div>{renderAmis(view)}</div>;
}
