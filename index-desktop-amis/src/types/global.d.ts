import {ipcRenderer} from "electron";

export interface IElectronAPI {
    openPage(info: WindowsSendMessage),
    download(url:string),
    onProgress(handle:(event, message)=>void),
}
declare global {
    // 定义自己的window对象，封装自己的api
    interface Window {
        electron: IElectronAPI
    }
    // 定义全局的通信格式
    interface WindowsSendMessage {
        url:        string // 页面url
        height?:     number // 页面高度
        width?:      number // 页面宽度
        frame?:  boolean // 是否显示frame
        title?: string // 标题
        logo?: string // 页面logo
        dev?: boolean // 是否显示开发者工具
        resize?: boolean // 是否可以改变大小
    }
}
