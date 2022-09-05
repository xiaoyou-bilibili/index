const nodeFile = "file"
const nodeFolder = "folder"
const rootFolder = 891
const folderInfo = {name:'',create_time:''}
const fileInfo = {name:'',object_id:'',file_size:'',file_type:'',create_time:''}

const preview_folder = "https://index.xiaoyou.host/data/object/62f6efb8e60c06635d176de1"

const get_preview = (name, id) => {
    switch(tools.GetFileSuffix(name)) {
        case '': return "62f6efe0e60c06635d176de2"
        case "mp3": return "62f6f016e60c06635d176de3"
        case "jpg":
        case "png": return id
        case "mp4": return "62f6f021e60c06635d176de4"
        case "pdf": return "62f6f02ee60c06635d176de5"
        case "zip": return "62f6f04be60c06635d176de7"
        case "rar": return "62f6f03ee60c06635d176de6"
        case "txt": return "62f6f058e60c06635d176de8"
        case "exe": return "62f6f063e60c06635d176de9"
        case "md": return  "62f6f072e60c06635d176dea"
        default: return  "62f6efe0e60c06635d176de2"
    }
}

const get_file_type = (name) => {
    switch(tools.GetFileSuffix(name)) {
        case "png":
        case "jpg": return "img"
        case "mp3": return "music"
        case "txt":
        case "pdf":
        case "doc":
        case "md": return "doc"
        case "mp4":
        case "mkv": return "video"
        default: return "other"
    }
}

// 新建文件夹
gin.Handle("POST", "/folder", ctx => {
    const req = {parent_id:0, name:''}
    gin.BindData(ctx, req, data=>{
        let info = relationServer.AddNode(ctx, nodeFolder, {name: data.name, create_time: tools.GetUnix()})
        if(info.err != null) { return gin.Fail(ctx, "add node err %v", info.err) }
        // parentID为0说明是根文件夹
        if(data.parent_id === 0) { data.parent_id = rootFolder }
        let res = relationServer.AddRelation(ctx, data.parent_id, info.id, nodeFile, {})
        if(res.err != null) { return gin.Fail(ctx, "add relation err %v", res.err) }
        gin.Success(ctx, {id: res.id})
    })
})
// 获取文件夹下所有文件
gin.Handle("GET", "/folder/:id", ctx => {
    gin.GetPathInt(ctx, "id", id=>{
        let page = gin.GetPageField(ctx)
        let res = {err:null,nodes:[],total:0}
        if(id < 0 ){
            let keyword = "other"
            switch(id) {
                case -1: keyword="img";break
                case -2: keyword="doc";break
                case -3: keyword="video";break
                case -4: keyword="music";break
            }
            res = relationServer.FindNode(ctx, nodeFile, "file_type", keyword, page.current, page.size, fileInfo)
        } else {
            res = relationServer.GetNodeChild(ctx, id, nodeFile, page.current, page.size, fileInfo)
        }
        if (res.err != null) { return gin.Fail(ctx, "get data err %V", res.err) }
        let fileList = []
        res.nodes.forEach(value=>{
            let attribute = value.Info.Attribute
            let info = {id: value.Id, name: attribute.name, create_time: tools.ParseUnix(attribute.create_time)}
            switch(value.Info.NodeLabel[0]){
                case nodeFolder:
                    info.file_type = 0
                    info.preview = preview_folder
                    break
                case nodeFile:
                    info.file_type = 1
                    info.preview = tools.GetObjectLink(get_preview(info.name, attribute.object_id))
                    info.info = attribute
            }
            fileList.push(info)
        })
        gin.ReturnPageInfo(ctx, page.current, res.total, fileList)
    })
})
// 新建文件
gin.Handle("POST", "", ctx => {
    const req = {name:'',object_id:'',file_size:'', parent:0}
    gin.BindData(ctx, req, data => {
        let node = {name:data.name, object_id:data.object_id, file_size:data.file_size, file_type:get_file_type(data.name), create_time:tools.GetUnix()}
        let res = relationServer.AddNode(ctx, nodeFile, node)
        if (res.err != null || res.id == 0) { return gin.Fail(ctx, "add relation err  %v", info.err) }
        if(data.parent == 0) {data.parent = rootFolder}
        let r = relationServer.AddRelation(ctx, data.parent, res.id, nodeFile, {})
        if (r.err != null) { return gin.Fail(ctx, "get relation error %v", r.err) }
        mqServer.Add(ctx, res.id, "file", data.name, "", [])
        gin.Success(ctx, {id: r.id})
    })
})
// 删除文件
gin.Handle("DELETE", "/:id", ctx => {
    let idList = gin.GetPathIntList(ctx, "id")
    idList.forEach(id => {
        // 获取文件
        let node = relationServer.GetNode(ctx, id, fileInfo)
        if(node.err != null) { return gin.Fail(ctx, "get node error %v", node.err) }
        if(node.info.object_id !== null) { dataServer.DeleteObject(ctx, [node.info.object_id]) }
        mqServer.Delete(ctx, id)
    })
    let err = relationServer.DeleteNode(ctx, idList)
    if(err != null) { return gin.Fail(ctx, "delete data err %v", err) }
    gin.Success(ctx, null)
})
// 修改文件信息
gin.Handle("PUT", "/:id", ctx => {
    const req = {name: ''}
    gin.GetPathInt(ctx, "id", id => {
        gin.BindData(ctx, req, data => {
            let err = relationServer.UpdateNode(ctx, id, data)
            if (err != null) { return gin.Fail("update node err %v", err) }
            mqServer.Update(ctx, id, "file", data.name, "", [])
            gin.Success(ctx, null)
        })
    })
})
// 查找文件
gin.Handle("GET", "", ctx => {
    let page = gin.GetFindField(ctx)
    let res = relationServer.FindNode(ctx, nodeFile, page.field, page.keyword, page.current, page.size)
    if(res.err != null) { return gin.Fail(ctx, "fin node err %v", res.err) }
    let data = relationServer.NodeRange(res.nodes, fileInfo)
    data = data.map(info => {
        return { id: info.id, name: info.name, create_time: tools.ParseUnix(info.create_time), file_type: 1, preview: tools.GetObjectLink(get_preview(info.name, info.id)), info: info}
    })
    gin.ReturnPageInfo(ctx, page.current, res.total, data)
})
// 获取文件信息
gin.Handle("GET", "/:id", ctx => {})
// 获取下载链接
gin.Handle("GET", "/:id/data", ctx => {
    gin.GetPathInt(ctx, "id", id => {
        let data = relationServer.GetNode(ctx, id, fileInfo)
        if (data.err != null) { return gin.Fail(ctx, "get node err %v", data.err) }
        if (data.info.object_id == null || data.info.object_id == "") { return gin.Fail(ctx, "node file found") }
        let link = dataServer.GetDownloadLink(ctx, data.info.object_id)
        if (link.err != null) { return web.Fail(ctx, "get link err %v", link.err)  }
        ctx.Redirect(302, link.link)
    })
})
// 全量数据同步
gin.Handle("GET", "/sync", ctx => {
    // 先把旧记录全部删除
    mqServer.DeleteAll(ctx, "file")
    // 遍历所有节点
    relationServer.RangeAllNode(ctx, "file", fileInfo, nodes => {
        nodes.forEach(item => {
            let info = item.Info.Attribute
            mqServer.Add(ctx, item.Id, "file", info.name, "", [])
        })
    })
    gin.Success(ctx, "同步成功")
})
// 文件卡片
view.HandleSearch((ctx, id) => {
    let data = relationServer.GetNode(ctx, id, fileInfo)
    tools.LogInfo("data is %v", data.info)
    if (data.err != null || data.info.object_id == null || data.info.object_id == "") { return "" }
    let link = dataServer.GetDownloadLink(ctx, data.info.object_id)
    let info = data.info

    let card = {
      "type": "grid",
      "columns": [
        {"md":1,"body":[{"type":"image","width":"60px","height":"60px","imageMode":"original","src": tools.GetObjectLink(get_preview(info.name, info.object_id))}]},
        {"md":8,"body":[{"type":"html","html":`<div style=\"font-weight:800\">${info.name}</div><div style=\"color: #333;font-size: 13px;\">大小: ${info.size}</div><div style=\"font-size: 13px;color: #9195A3;\">修改日期：${tools.ParseUnix(info.create_time)}</div>`}]},
        {"md":2,"body":[{"type":"button-group","vertical":true,"buttons":[
            {"icon":"fa fa-clipboard","type":"button","label":"复制链接","actionType": "copy","content": link.link},
            {"icon":"fa fa-cloud-download","type":"button","label":"下载文件", "onEvent":{"click":{"actions":[{"actionType":"index-download","args":{ "url": link.link}}]}}}
        ]}]}
      ]
    }
    return JSON.stringify(card)
})
view.RegisterManage("文件管理", 6742, ()=>{})