// tools工具类
declare class tools {
    // 打印日志
    static LogInfo(format:string,data:any)
    static LogCtxInfo(ctx:context, format:string, ...data:any)
    static LogCtxError(ctx:context, format:string, ...data:any)
    // 获取Object的链接
    static GetObjectLink(id:string):string
    // 发送HTTP请求
    static HttpSendRequest(ctx:context, url:string, method:string, header:any, data:any):{data:any, err:any}
    // 获取consul的配置
    static GetConsulKV(ctx:context, key:string):string
    // 获取一个默认的context
    static ContextBackground():context
    // 获取Unix时间戳
    static GetUnix():string
    // 解析unix时间戳
    static ParseUnix(data:string):string
    // 获取文件后缀
    static GetFileSuffix(filename:string):string
    // 返回错误
    static ReturnError(format:string, ...data:any):any
}

// gin context对象
declare interface context {
    // 获取路径参数
    Param(filed:string):string
    Redirect(code:number, link:string)
}

// gin接口
declare class gin {
    // 处理函数
    static Handle(method:string, url:string, fun: (ctx: context) => void)
    // 成功的回调函数
    static Success(ctx:context, data:any)
    // 失败的返回
    static Fail(ctx:context,format:string,...data:any)
    // 快速获取path路径上的字段
    static GetPathInt(ctx:context,field:string,callback:(data:number)=>void)
    // 快速获取query参数上的int字段
    static GetQueryInt(ctx:context,field:string):number
    // gin快速获取路径上的int数组
    static GetPathIntList(ctx:context,field:string):number[]
    // 绑定数据
    static BindData<T>(ctx:context,req:T, callback:(data:T)=>void)
    // 快速获取Find接口的四个字段
    static GetFindField(ctx:context):{keyword:string, field:string, current:number, size:number}
    // 获取分页的两字段
    static GetPageField(ctx:context):{current:number, size:number}
    // 返回分页信息
    static ReturnPageInfo(ctx:context,current:number,total:number,list:any)
}

// data服务接口
declare class dataServer {
    // 下载链接并上传
    static DownloadLinkAndUpload(ctx:context, url:string, name:string, header:any):{id:string,err:any}
    // 从gin的字段中自动上传对象
    static UploadObjectFromFile(ctx:context, field:string):{id:string,err:any}
    //添加一个文本对象
    static AddText(ctx:context, name:string, content:string):{id,err}
    //更新文本对象
    static UpdateText(ctx:context, id:string, name:string, content:string):any
    //获取文本对象
    static GetText(ctx:context, id:string):{name:string,content:string,err:any}
    //删除对象
    static DeleteObject(ctx:context, id:string[]):any
    //获取对象的下载链接
    static GetDownloadLink(ctx:context, id:string):{link:string,err:any}
}

// node结构定义
interface node<T> {
    Id:number
    Info:{NodeLabel:string[], Attribute:T}
}

// relation服务接口
declare class relationServer {
    // 添加一个节点
    static AddNode(ctx:context, name:string, info:any):{id:number, err:any}
    // Gin路由快速添加节点
    static GinAddNode(ctx:context, name:string, data:any, successCallback:(id:number)=>void)
    // 获取节点信息
    static GetNode<T>(ctx:context, id:number, data:T):{info:T,err}
    // 删除节点
    static DeleteNode(ctx:context, id:number[]):any
    // 添加联系
    static AddRelation(ctx:context,start:number, end:number, relationInfo:string,info:any): {id: number, err: any}
    // 获取某个节点下所有子节点
    static GetNodeChild<T>(ctx:context, id :number, relationType:string, current:number, size:number, data:T):{nodes:node<T>[],total:number,err:any}
    // 获取某个节点下所有的父节点
    static GetNodeParent<T>(ctx:context, id :number, relationType:string, current:number, size:number, data:T):{nodes:node<T>[],total:number,err:any}
    //根据两个节点的关系删除联系
    static DeleteRelationWithNode(ctx:context, start :number, end :number):any
    //查找节点
    static FindNode<T>(ctx:context, label:string, field:string, keyword:string, current:number, size:number, data:T):{nodes:node<T>[],total:number,err:any}
    // 更新节点
    static UpdateNode(ctx:context, id :number, attribute :Map<string,string>):any
    // 节点快速遍历
    static NodeRange<T>(nodes:node<T>[],data:T):T[]
    // 新建一个tag
    static AddTag(ctx:context, name:string):{id:number,err:any}
    // 遍历所有节点
    static RangeAllNode<T>(ctx:context,name:string,data:T,handle:(nodes:node<T>[])=>void)
}

declare class mqServer {
    // 新增对象
    static Add(ctx:context, id:number, app:string, name:string, content:string, tags:string[])
    // 删除对象
    static Delete(ctx:context, id:number)
    // 更新对象
    static Update(ctx:context,id:number, app:string, name:string, content:string, tags:string[])
    // 删除所有对象
    static DeleteAll(ctx:context, app:string)
}

declare class view {
    // 展示搜索卡片
    static HandleSearch(handle:(ctx:context, id:number)=>string)
    // 获取自定义页面
    static HandleView(handle:(ctx:context,id:number,view:string)=>any)
    // 注册管理界面
    static RegisterManage(name:string, id:number, handle:()=>Map<string, string>)
    // 获取页面数据
    static GetView(ctx:context,id:number, data:any,exception:boolean):any
}