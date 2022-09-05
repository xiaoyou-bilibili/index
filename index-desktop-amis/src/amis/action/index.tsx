import {
    ListenerAction,
    ListenerContext,
    registerAction,
    RendererAction
} from 'amis-core';
import {RendererEvent} from 'amis-core';
import request from "@/api/request";
import { base } from "@/api/api";
import {message} from "antd";

// 动作定义
interface IRequestAction extends ListenerAction {
    actionType: 'index-request';
    args: {
        method: string, // 请求方式
        url: string, // 请求地址
        data: any, // 请求的数据
        numberField: [], // 需要转换为int的字段
        numberListField: [] // 需要转换为intList的字段
    };
}

// 发送请求
export class RequestAction implements RendererAction {
    // @ts-ignore
    run(action: IRequestAction, renderer: ListenerContext, event: RendererEvent<any>) {
        const {method, url, data, numberField, numberListField} = action.args;
        if (numberField != undefined) {
            numberField.forEach(((field)=>{
                data[field] = Number(data[field])
            }))
        }
        if (numberListField != undefined) {
            numberListField.forEach(((field)=>{
                data[field] = String(data[field]).split(",").map(value => Number(value))
            }))
        }
        console.log(method, url, data)
        request(`${base}${url}`, data, method).then(() => message.success('操作成功'))
    }
}

// 打开页面
interface IPageAction extends ListenerAction {
  actionType: 'index-page';
  args: {
    dev: boolean,
    frame: boolean,
    height: number,
    logo: string,
    resize: boolean,
    url: string,
    width: number,
    title: string
  };
}
export class PageAction implements RendererAction {
  // @ts-ignore
  run(action: IPageAction, renderer: ListenerContext, event: RendererEvent<any>) {
    const {dev, frame, height, logo, resize, url, width, title} = action.args;
    window.electron.openPage({dev, frame, height, logo, resize, url, width, title})
  }
}

// 下载事件
interface IDownloadAction extends ListenerAction {
  actionType: 'index-download';
  args: { url: string, };
}
export class DownloadAction implements RendererAction {
  // @ts-ignore
  run(action: IDownloadAction, renderer: ListenerContext, event: RendererEvent<any>) {
    const {url} = action.args;
    window.electron.download(url)
  }
}

export default function initAction() {
    // 注册自定义动作
    // @ts-ignore
    registerAction('index-request', new RequestAction());
    // @ts-ignore
    registerAction('index-page', new PageAction());
    // @ts-ignore
    registerAction('index-download', new DownloadAction());
}
