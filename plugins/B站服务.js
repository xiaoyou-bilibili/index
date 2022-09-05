const unique = "bilibili"
const uid = "343147393"
const getHeader = () => {
    return {
        "cookie": tools.GetConsulKV(tools.ContextBackground(), "index/app/bili_cookie")
    }
}
// 获取B站粉丝数
const getBilibiliFans = (ctx) => {
    let res = tools.HttpSendRequest(ctx,`https://api.bilibili.com/x/relation/stat?vmid=${uid}&jsonp=jsonp`,"GET",getHeader(),null)
    if(res.err != null) { return 0}
    return res.data.data.follower
}
// 获取播放数
const getBilibiliViews = (ctx) => {
    let res = tools.HttpSendRequest(ctx,`https://api.bilibili.com/x/space/upstat?mid=${uid}&jsonp=jsonp`,"GET",getHeader(),null)
    if(res.err != null) { return 0}
    return res.data.data.archive.view
}
gin.Handle("GET", "/fans", ctx => {
    gin.Success(ctx, {
        fans: getBilibiliFans(ctx),
        views: getBilibiliViews(ctx),
    })
})
// 全量数据同步
gin.Handle("GET", "/sync", ctx => {
    // 先把旧记录全部删除
    mqServer.DeleteAll(ctx, unique)
    mqServer.Add(ctx, -1, unique, "B站粉丝", "", [])
    // 添加粉丝
    gin.Success(ctx, "同步成功")
})
// B站粉丝数卡片
view.HandleSearch((ctx,id) => {
    let card ={
        "type": "service",
        "data": { "object": tools.GetObjectLink(), "fans": getBilibiliFans(), "plays": getBilibiliViews() },
        "body": {
            "type": "grid","className": "index-align-item index-text-align-center",
            "columns": [
                {"md":1,"body":[{"type":"image","width":"60px","height":"60px","imageMode":"original","src":"${object}6309bff2e891a61a94783035"}]},
                {"md":5,"body":[{"type":"html","html":"<span style=\"font-size:50px;font-weight:600;color:#20b0e3\">${fans}</span>"}]},
                {"md":1,"body":[{"type":"image","width":"60px","height":"60px","imageMode":"original","src":"${object}6309c76ae891a61a94783038"}]},
                {"md":5,"body":[{"type":"html","html":"<span style=\"font-size:50px;font-weight:600;color:#fb7299\">${plays}</span>"}]}
        ]}
    }
    // 返回内容
    return JSON.stringify(card)
})
