import {
    ListenerAction,
    ListenerContext,
    registerAction,
    RendererAction
} from 'amis-core';
import {RendererEvent} from 'amis-core';

// 动作定义
interface IMyAction extends ListenerAction {
    actionType: 'index-request';
    args: {
        method: string, // 请求方式
        url: string, // 请求地址
        data: any, // 请求的数据
        numberField: [], // 需要转换为int的字段
        numberListField: [] // 需要转换为intList的字段
    };
}

/**
 * 我的动作实现
 */
export class RequestAction implements RendererAction {
    // @ts-ignore
    run(action: IMyAction, renderer: ListenerContext, event: RendererEvent<any>) {
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
    }
}

export default function initAction() {
    // 注册自定义动作
    // @ts-ignore
    registerAction('index-request', new RequestAction());
}
