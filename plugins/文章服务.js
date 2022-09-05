const nodeArticle = "article"
const nodeTag = "tag"
const articleInfo = {name:"", content:""}
const articleWitchTag = {name:"",content:"",tags:""}

const getArticleContent = (ctx, id) => {
    let node = relationServer.GetNode(ctx, id, articleInfo)
    if (node.err != null) {return {data:null, err:tools.ReturnError("get node err %v", node.err)}}
    let res = {name:node.info.name, content:node.info.content}
    let content = dataServer.GetText(ctx, node.info.content)
    if (content.err != null) {return {data:null, err:tools.ReturnError("get content err %v", content.err)}}
    res.content = content.content
    let tags = relationServer.GetNodeParent(ctx, id, nodeArticle, 0, 0, nodeArticle)
    if (tags.err != null) {return {data:null, err:tools.ReturnError("get tags err %v", tags.err)}}
    res.tags = tags.nodes.map(node=>{ return node.Info.Attribute["name"]}).join(",")
    return {data: res, err: null}
}

// 新增文章
gin.Handle("POST", "/article", ctx => {
    gin.BindData(ctx, articleWitchTag, data=>{
        let content = dataServer.AddText(ctx, data.name, data.content)
        if (content.err != null) { return gin.Fail(ctx, "add text err %v", content.err) }
        data.content = content.id
        let node = relationServer.AddNode(ctx, nodeArticle, data)
        if (node.err != null) { return gin.Fail(ctx, "add node err %v", content.err) }
        data.tags.split(",").forEach(item=>{
            let tag = relationServer.AddTag(ctx, item)
            if (tag.err == null) {
                relationServer.AddRelation(ctx, tag.id, node.id, nodeArticle, {})
            }
        })
        mqServer.Add(ctx, node.id, "article", data.name, data.content, data.tags.split(","))
        gin.Success(ctx, {id: node.id})
    })
})
// 获取文章内容
gin.Handle("GET","/article/:id",ctx=>{
    gin.GetPathInt(ctx, "id", id=>{
        let data = getArticleContent(ctx, id)
        if (data.err != null) { return gin.Fail(ctx, "%v", data.err) }
        gin.Success(ctx, data.data)
    })
})
// 更新文章
gin.Handle("PUT","/article/:id",ctx=>{
    gin.GetPathInt(ctx, "id", id => {
        gin.BindData(ctx, articleWitchTag, data=>{
            let article = relationServer.GetNode(ctx, id, articleInfo)
            if (article.err != null) { return gin.Fail(ctx, "get node err %v", article.err) }
            let err = dataServer.UpdateText(ctx, article.info.content, data.name, data.content)
            if (err != null) { return gin.Fail(ctx, "get node err %v", err) }
            if (article.info.name != data.name) { relationServer.UpdateNode(ctx, id, {"name": data.name}) }
            let tags = relationServer.GetNodeParent(ctx, id, nodeArticle, 0, 0, articleInfo)
            if  (tags.err == null) {
                let tagMap = new Map()
                tags.nodes.forEach(node => { tagMap.set(node.Info.Attribute["name"], node) })
                data.tags.split(",").forEach(tag => {
                    if(tagMap.has(tag)) { tagMap.delete(tag) } else {
                        let node = relationServer.AddTag(ctx, tag)
                        if (node.id != 0) {
                            relationServer.AddRelation(ctx, node.id, id, nodeArticle, {})
                        }
                    }
                })
                tagMap.forEach((node) => relationServer.DeleteRelationWithNode(ctx, node.Id, id))
            }
            mqServer.Update(ctx, id, "article", data.name, data.content, data.tags.split(","))
            gin.Success(ctx, null)
        })
    })
})
// 获取文章列表
gin.Handle("GET","/article",ctx=>{
    let page = gin.GetFindField(ctx)
    let data = relationServer.FindNode(ctx, nodeArticle, page.field, page.keyword, page.current, page.size)
    if (data.err != null) { return gin.Fail(ctx, "find node err %v", data.err) }
    let res = relationServer.NodeRange(data.nodes, articleInfo)
    gin.ReturnPageInfo(ctx, page.current, data.total, res)
})
// 删除文章
gin.Handle("DELETE","/article/:id",ctx=>{
    let idList = gin.GetPathIntList(ctx, "id")
    let contents = []
    idList.forEach(id => {
        let node = relationServer.GetNode(ctx, id, articleInfo)
        if (node.err == null) {
            contents.push(node.info.content)
            mqServer.Delete(ctx, id)
        }
    })
    dataServer.DeleteObject(ctx, contents)
    let err = relationServer.DeleteNode(ctx, idList)
    if (err != null) { gin.Fail(ctx, "delete relation %v", err) }
    gin.Success(ctx, null)
})
// 上传图片
gin.Handle("POST","/img",ctx=>{
    let data = dataServer.UploadObjectFromFile(ctx, "file")
    if (data.err != null) { return gin.Fail(ctx, "update data err %v", data.err) }
    gin.Success(ctx, {url: tools.GetObjectLink(data.id)})
})
// 自动同步
gin.Handle("GET","/article/sync",ctx=>{

})
// 文章卡片
view.HandleSearch((ctx,id) => {
    // 获取文章内容
    let data = getArticleContent(ctx, id)
    if (data.err != null) { return "" }
    let content = String(data.data.content)
    if (content.length > 100) {
        content = content.substring(0, 100)
    }
    let card = {
        "type": "grid",
        "columns": [
            {"md":1,"body":[{"type":"image","width":"60px","height":"60px","imageMode":"original","src":tools.GetObjectLink("630077a48668adf7ba1b9c24")}]},
            {"md":8,"body":[{"type":"html","html":`<h2 style=\"margin-top:5px;border-bottom:1px solid #bbb;font-size:16px;font-weight:600;\">${data.data.name}</h2><p style=\"font-size: 13px;color: #86909c;margin-top:-6px;max-height:38px;overflow:hidden;\">${content}</p>`}]},
            {"md":1,"body":[{"type":"button-group","className":"m-t-xs","buttons":[
                        {"icon":"fa fa-pencil-square","type":"button","label":"编辑",
                            "onEvent":{"click":{"actions":[{"actionType": "index-page","args": {"url": `/view/article/${id}/edit`,"width": 1200,"height": 700,"frame":"true","title":"文章显示","dev":true}}]}}
                        },
                        {"icon":"fa fa-eye","type":"button","label":"查看",
                            "onEvent":{"click":{"actions":[{"actionType": "index-page","args": {"url": `/view/article/${id}/view`,"width": 600,"height": 700,"frame":"true","title":"文章显示","dev":true}}]}}
                        }
                    ]}]}
        ]
    }
    return JSON.stringify(card)
})
view.RegisterManage("文章管理", 6741, ()=>{})
// 自定义文章显示界面
view.HandleView((ctx, id, action) => {
    // 如果是显示，那么就把文章界面给显示出来
    if(action == "view") {
        // 获取文章内容
        let data = getArticleContent(ctx, id)
        if (data.err != null) { return "" }
        let article = data.data
        article.content = `## ${article.name}\n${article.content}`
        // 获取页面信息
        return { "page": view.GetView(ctx, 6736, article) }
    } else {
        let data = {id: 0, name: "", content: "", tags: ""}
        if (id != 0) {
            // 如果id不为0才去获取相关信息
            let res = getArticleContent(ctx, id, articleInfo)
            if (res.err!=null) { return "" }
            data = res.data
            data.id = id
        }
        // 否则为编辑
        return { "page": view.GetView(ctx, 6743, data) }
    }

})