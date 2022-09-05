import {makeSorter, render as renderAmis, updateEnv} from 'amis';
import * as monaco from 'monaco-editor';
import React, {useCallback, useEffect, useState} from 'react'
import {Select, Form, Button, message, Input, Modal} from 'antd';
import {
    addPlugin, deletePlugin,
    deleteView,
    getPluginInfo,
    getPluginInterface,
    getPluginList,
    pluginReload,
    updatePluginInfo
} from "../api/api";
import {SaveOutlined, ReloadOutlined, FileAddOutlined, DeleteOutlined} from '@ant-design/icons';
const {Option} = Select;
let editor: monaco.editor.IStandaloneCodeEditor;
// 开启外部ts定义功能
monaco.languages.typescript.javascriptDefaults.setCompilerOptions({
    target: monaco.languages.typescript.ScriptTarget.ES2016,
    allowNonTsExtensions: true,
})

export default function Code() {
    let [windowHeight, setWindowHeight] = useState(700)
    let [pluginList, setPluginList] = useState([])
    // 插件名字和插件标识
    const [pluginName, setPluginName] = useState('');
    const [pluginUnique, setPluginUnique] = useState('');
    // 添加插件是否可见
    const [addPluginVisible, setAddPluginVisible] = useState(false);
    //  新建插件的名字
    const [addPluginName, setAddPluginName] = useState('')
    const [addPluginUnique, setAddPluginUnique] = useState('')

    const editInit = () => {
        editor = monaco.editor.create(document.getElementById("codeEditBox") as HTMLElement, {
            value: '', // 编辑器初始显示文字
            language: 'javascript', // 使用JavaScript语言
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
    }

    const handleKeyDown = (event: React.KeyboardEvent) => {
        if (event.key === "s" && (event.ctrlKey || event.metaKey)) {
            console.log(pluginUnique)
            pluginUnique === ''? message.error('请选择插件'):savePlugin()
            // 阻止默认事件
            event.preventDefault();
        }
    };

    // useEffect只渲染一次
    const viewInit = useCallback(() => {
        // 清除所有插件
        monaco.editor.getModels().forEach(model => model.dispose());
        setWindowHeight(window.innerHeight - 150)
        // 调用接口获取代码提示
        getPluginInterface().then((res: any) => {
            const libSource = res.data
            const libUri = 'ts:filename/index.d.ts';
            monaco.languages.typescript.javascriptDefaults.addExtraLib(libSource, libUri);
            monaco.editor.createModel(libSource, 'typescript', monaco.Uri.parse(libUri));
            editInit()
        })
        // 获取插件列表
        getAllPlugin()
    }, [])

    useEffect(() => viewInit(), [viewInit])

    // 渲染options
    const renderOption = (data:any) => {
        if (data == null) {return}
        console.log(data)
        let plugins: JSX.Element[] = []
        data.forEach((item:any)=>plugins.push(<Option key={item.id} value={item.unique}>{`${item.name}`}</Option>))
        return plugins
    }

    // 插件选择事件
    const optionChoose = (value: string) => {
        setPluginUnique(value)
        getPluginInfo(value).then((data:any) => {
            editor?.setValue(data.code)
            setPluginName(data.name)
            setPluginUnique(data.unique)
        })
    }

    // 获取所有插件
    const getAllPlugin = () => getPluginList().then((data: any) => setPluginList(data))

    // 保存插件
    const savePlugin = () => updatePluginInfo(pluginUnique, {code: editor?.getValue(), name: pluginName, unique: pluginUnique}).then(()=>message.success('保存成功'))
    // 插件重载
    const reloadPlugin = () => { pluginReload(pluginUnique).then(_ => message.success('重载成功'))}

    // 输入框渲染
    const renderPluginNameInput = (name:string) => {
        return <Input placeholder={"请输入插件名称"} defaultValue={name} value={pluginName} onChange={(e:any)=>{setPluginName(e.target.value)}} />
    }
    const renderPluginUniqueInput = (name:string) => {
        return <Input placeholder={"请输入插件标识"} defaultValue={name} value={pluginUnique} onChange={(e:any)=>{setPluginUnique(e.target.value)}} />
    }
    // 删除所有插件
    const deletePluginAction = () => deletePlugin(pluginUnique).then(()=>{
        message.success("删除成功")
        getAllPlugin()
    })
    // 新建插件
    const addPluginAction = () => addPlugin({name: addPluginName, unique: addPluginUnique, code: "// 输入代码"}).then(_=>{
        setAddPluginVisible(false)
        getAllPlugin().then()
        message.success('添加成功')
    })

    return (
        <div style={{padding: "10px"}}>
            <Form style={{padding: "5px", marginBottom: "10px"}} layout="inline">
                <Form.Item><Button type="primary" onClick={()=>{setAddPluginVisible(true)}} icon={<FileAddOutlined />}>新增插件</Button></Form.Item>
                <Form.Item label="插件" name="plugin"><Select onSelect={optionChoose} style={{width: "150px"}}>{renderOption(pluginList)}</Select></Form.Item>
                <Form.Item label="名称">{renderPluginNameInput(pluginName)}</Form.Item>
                <Form.Item label="标识">{renderPluginUniqueInput(pluginUnique)}</Form.Item>
                <Form.Item><Button type="primary" onClick={savePlugin} icon={<SaveOutlined />}>保存</Button></Form.Item>
                <Form.Item><Button type="primary" onClick={reloadPlugin} icon={<ReloadOutlined />} >重载</Button></Form.Item>
                <Form.Item><Button type="primary" danger onClick={deletePluginAction} icon={<DeleteOutlined />}>删除插件</Button></Form.Item>
            </Form>
            <div onKeyDown={handleKeyDown}><div style={{height: windowHeight}} id="codeEditBox" /></div>
            <Modal title="新增插件" visible={addPluginVisible} onOk={addPluginAction} onCancel={()=>{setAddPluginVisible(false)}}>
                <Input addonBefore={"插件名称"} placeholder={"请输入插件名称"} value={addPluginName} onChange={(e:any)=>setAddPluginName(e.target.value)}  />
                <Input style={{marginTop: "5px"}} addonBefore={"插件标识"} placeholder={"请输入插件标识"} value={addPluginUnique} onChange={(e:any)=>setAddPluginUnique(e.target.value)}  />
            </Modal>
        </div>
    )
}
