import request from "./request"

export const base = 'https://index.xiaoyou.host'

// 搜索内容
export function search(keyword:string) {return request(`${base}/search?q=${keyword}`, null, 'get')}
// 获取自定义界面
export function getCustomView(name:string,id:string,view:string) {return request(`${base}/app/core/view/app/${name}/${id}/${view}`, null, 'get')}

