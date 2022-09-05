import request from "./request"

const base = 'https://index.xiaoyou.host'

// 搜索内容
export function search(keyword:string) {return request(`${base}/search?q=${keyword}`, null, 'get')}

interface pluginInfo {
    name:string,
    code:string,
    unique:string
}

// 新建插件
export function addPlugin(data:pluginInfo) {return request(`${base}/app/core/plugin`, data, 'post')}
// 获取插件列表
export function getPluginList() {return request(`${base}/app/core/plugins`, {}, 'get')}
// 获取插件信息
export function getPluginInfo(name:string) {return request(`${base}/app/core/plugin/${name}`, {}, 'get')}
// 获取插件定义
export function getPluginInterface() {return request(`${base}/app/core/interface/plugin`, {}, 'get')}
// 更新插件信息
export function updatePluginInfo(name:string,data:pluginInfo) {return request(`${base}/app/core/plugin/${name}`, data, 'put')}
// 重载加载
export function pluginReload(name:string) {return request(`${base}/app/core/plugin/${name}/reload`, {}, 'get')}
// 删除插件
export function deletePlugin(name:string) {return request(`${base}/app/core/plugin/${name}`, {}, 'delete')}

// 新建页面
export function addView(data:{name:string,view:string}) { return request(`${base}/app/core/views`, data, 'post') }
// 获取所有页面
export function getAllView() { return request(`${base}/app/core/views`, {}, 'get') }
// 更新页面
export function updateView(id:number,data:{name:string,view:string}) { return request(`${base}/app/core/views/${id}`, data, 'put') }
// 删除页面
export function deleteView(id:number) { return request(`${base}/app/core/views/${id}`, {}, 'delete') }
// 获取某个页面
export function getView(id:number) { return request(`${base}/app/core/views/${id}`, {}, 'get') }
// 获取所有管理界面
export function getManagesView() { return request(`${base}/app/core/view/manage`, {}, 'get') }

