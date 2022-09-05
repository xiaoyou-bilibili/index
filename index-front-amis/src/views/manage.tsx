import { useParams } from 'react-router-dom';
import React, {useCallback, useEffect, useState} from "react";
import {getManagesView, getView} from "../api/api";
import {render as renderAmis} from "amis-core";


export default function Manage(props:any) {
    const id = Number(useParams()["id"])
    let [view, setView] = useState('{}')

    // 我们监听路由参数变化，如果路由变更那么就重新渲染
    useEffect(() => {
        // 获取所有的子菜单
        getView(id).then((data:any)=>setView(data.view))
    }, [useParams()])

    const renderView = () => {return <div>{renderAmis(JSON.parse(view))}</div>}

    return <div>{renderView()}</div>
}
