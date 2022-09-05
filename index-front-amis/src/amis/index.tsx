import initAction from "./action";
import initComponent from "./component";
import axios from "axios";
import {updateEnv} from "amis";
import copy from 'copy-to-clipboard';
import {message} from "antd";

const fetch = ({ url, method, data, responseType, config, headers }: any) => {
    config = config || {};
    // config.withCredentials = true;
    responseType && (config.responseType = responseType);

    if (config.cancelExecutor) {
        config.cancelToken = new (axios as any).CancelToken(
            config.cancelExecutor
        );
    }

    config.headers = headers || {};

    if (method !== 'post' && method !== 'put' && method !== 'patch') {
        if (data) {
            config.params = data;
        }
        return (axios as any)[method](url, config);
    } else if (data && data instanceof FormData) {
        config.headers = config.headers || {};
        config.headers['Content-Type'] = 'multipart/form-data';
    } else if (
        data &&
        typeof data !== 'string' &&
        !(data instanceof Blob) &&
        !(data instanceof ArrayBuffer)
    ) {
        data = JSON.stringify(data);
        config.headers = config.headers || {};
        config.headers['Content-Type'] = 'application/json';
    }
    return (axios as any)[method](url, data, config)
}
const isCancel = (value: any) => axios.isCancel(value)

export default function AmisRegister () {
    // 更新env，用于获取数据
    updateEnv({fetcher: fetch, isCancel: isCancel, copy: content => {copy(content);message.success('内容已复制到粘贴板')}})
    // 注册自定义事件
    initAction()
    // 注册自定义组件
    initComponent()
}
