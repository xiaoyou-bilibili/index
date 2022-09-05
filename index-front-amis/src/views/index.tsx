import React, {useCallback, useEffect, useState} from 'react';
import {render as renderAmis,runAction} from 'amis';
import {ListenerAction, ListenerContext, registerAction, RendererAction, RendererEvent} from "amis-core";
import {Button, Form, Select, Modal, Input, message} from "antd";
import {DeleteOutlined, SaveOutlined, FileAddOutlined} from "@ant-design/icons";
import {addView, deleteView, getAllView, getPluginInterface, getPluginList, getView, updateView} from "../api/api";
import * as monaco from "monaco-editor";
const {Option} = Select;
let editor: monaco.editor.IStandaloneCodeEditor;

export default function Index() {
    // 输入的json数据
    const [json, setJson] = useState('{}')
    let [viewList, setViewList] = useState([])
    const [addViewVisible, setAddViewVisible] = useState(false);
    const [addViewName, setAddViewName] = useState('');
    const [viewName, setViewName] = useState('');
    // 当前选择的页面信息
    const [viewID, setViewID] = useState(0)

    const editInit = () => {
        editor = monaco.editor.create(document.getElementById("viewEditBox") as HTMLElement, {
            value: '', // 编辑器初始显示文字
            language: 'json', // 使用JavaScript语言
            automaticLayout: true, // 自适应布局
            theme: 'vs-dark', // 官方自带三种主题vs, hc-black, or vs-dark
            foldingStrategy: 'indentation', // 代码折叠
            renderLineHighlight: 'all', // 行亮
            selectOnLineNumbers: true, // 显示行号
            minimap: {enabled: false}, // 是否开启小地图（侧边栏的那个全览图）
            readOnly: false, // 只读
            fontSize: 18, // 字体大小
            scrollBeyondLastLine: false, // 取消代码后面一大段空白
            overviewRulerBorder: false, // 不要滚动条的边框
        })
        editor.onDidBlurEditorText(() => {
            let input = editor.getValue()
            try {
                JSON.parse(input)
                setJson(input)
            } catch (e) {}
        })
    }

    // 渲染页面
    const renderView = () => {return <div>{renderAmis(JSON.parse(json))}</div>}

    // 渲染options
    const renderOption = (data:any) => {
        let plugins: JSX.Element[] = []
        data.forEach((item:any)=>plugins.push(<Option key={item.id} value={item.id}>{`${item.name}(${item.id})`}</Option>))
        return plugins
    }

    // 初始化获取所有页面信息
    const viewInit = useCallback(() => {
        monaco.editor.getModels().forEach(model => model.dispose());
        getAllViewAction().then()
        editInit()
    }, [])
    useEffect(() => viewInit(), [viewInit])

    // 获取所有页面
    const getAllViewAction = () => getAllView().then((data:any) => setViewList(data))

    // 添加页面
    const addViewAction = () => addView({name: addViewName, view: "{\"type\":\"page\",\"body\":[]}"}).then(_=>{
        getAllViewAction().then()
        setAddViewVisible(false)
        message.success('添加成功')
    })

    const handleKeyDown = (event: React.KeyboardEvent) => {
        if (event.key === "s" && (event.ctrlKey || event.metaKey)) {
            viewName === ''? message.error('请选择页面'):savePage()
            // 阻止默认事件
            event.preventDefault();
        }
    };


    // 多选框选择事件
    const selectAction = (value:number) => getView(value).then((data:any) => {
        setViewName(data.name)
        setViewID(value)
        editor.setValue(data.view)
    })

    // 保存页面
    const savePage = () => updateView(viewID, {name: viewName, view: editor.getValue()}).then(() => message.success('保存成功'))
    // 删除页面
    const deletePage = () =>deleteView(viewID).then(() => {
        message.success('删除成功')
        getAllViewAction().then()
    })

    // 输入框渲染
    const renderViewNameInput = (name:string) => {
        return <Input placeholder={"请输入页面名称"} defaultValue={name} value={viewName} onChange={(e:any)=>{setViewName(e.target.value)}} />
    }

    return (
        <div>
            <Form style={{padding: "10px"}} layout="inline">
                <Form.Item><Button type="primary" onClick={()=>{setAddViewVisible(true)}} icon={<FileAddOutlined />}>新增页面</Button></Form.Item>
                <Form.Item label="页面"><Select onSelect={selectAction} style={{width: "150px"}}>{renderOption(viewList)}</Select></Form.Item>
                <Form.Item label="名称">{renderViewNameInput(viewName)}</Form.Item>
                <Form.Item><Button type="primary" onClick={savePage} icon={<SaveOutlined />}>保存</Button></Form.Item>
                <Form.Item><Button type="primary" danger onClick={deletePage} icon={<DeleteOutlined />}>删除页面</Button></Form.Item>
            </Form>
            <div><div onKeyDown={handleKeyDown} style={{height: "400px"}} id="viewEditBox" /></div>
            {renderView()}
            <Modal title="新增页面" visible={addViewVisible} onOk={addViewAction} onCancel={()=>{setAddViewVisible(false)}}>
                <Input addonBefore={"页面名称"} placeholder={"请输入页面名称"} value={addViewName} onChange={(e:any)=>setAddViewName(e.target.value)}  />
            </Modal>
        </div>
    );
}
