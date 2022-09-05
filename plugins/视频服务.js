const movieInfo = {name:"",desc:"",poster:"",movie_type:""}
const seasonInfo = {name:"",poster:"",nfo:""}
const videoInfo = {name:"",no:"",thumb:"",mp4:"",barrage:"",subtitle:"",nfo:""}
const nodeSeason = "season"
const nodeMovie = "movie"
const nodeVideo = "video"
// 获取电影播放链接
// 获取影片详情
gin.Handle("GET", "/video/:id/mp4", ctx=>{
    gin.GetPathInt(ctx, "id", id => {
        let node = relationServer.GetNode(ctx, id, videoInfo)
        if (node.err != null) { return gin.Fail(ctx, "get node err %v", node.err) }
        let link = dataServer.GetDownloadLink(ctx, node.info.mp4)
        if (link.err != null) { return gin.Fail(ctx, "get dowload err %v", link.err) }
        ctx.Redirect(302, link.link)
    })
})
// 获取影片详情
gin.Handle("GET", "/movie/:id", ctx=>{})
// 删除影片
gin.Handle("DELETE", "/movie/:id", ctx=>{})
// 补充影片信息
gin.Handle("POST", "/movie", ctx=>{})
// 影片添加视频
gin.Handle("POST", "/movie/video", ctx=>{})
// 判断某集是否已经上传
gin.Handle("GET", "/movie/no", ctx=>{})
// 获取视频详情
gin.Handle("GET", "/movie/video/:id", ctx=>{})
// 获取影片列表
gin.Handle("DELETE", "/movie", ctx=>{})
// 查找视频
gin.Handle("DELETE", "", ctx=>{})
view.HandleSearch((ctx,id) => {
    let data = relationServer.GetNode(ctx, id, movieInfo)
    if (data.err != null) { return ""}
    let info = data.info
    let card = {"type": "flex", "justify": "flex-start", "className": "div-max",
      "items": [
        {"type":"image","width":"80px","height":"130px","imageMode":"original","src":tools.GetObjectLink(info.poster)},
        {"type": "flex", "direction": "column", "alignItems": "flex-start", "className": "m-l", "items": [
            {"type":"flex","justify":"flex-start","items":[
                {"type":"tag","label":"番剧","displayMode":"normal","color":"active"},
                {"icon":"fa fa-play-circle","type":"button","label":"播放","size": "sm", "className": "m-r",
                    "onEvent":{"click":{"actions":[{"actionType": "index-page","args": {"url": `/view/video/${id}/play`,"width": 1175,"resize": false,"height": 500,"frame":"true","title":"视频播放器","dev":true}}]}}
                },
                {"type":"html","html":`<div style=\"font-weight: 600;font-size:18px\">${info.name}</div>`},
            ]},
            {"type":"html","html":`<div style=\"font-size: 13px;line-height:18px;margin-bottom: 4px;color: #9499A0;height: 90px;overflow: scroll;overflow-x: hidden;\">${info.desc}</div>`}
          ]
        }
      ]
    }
    return JSON.stringify(card)
})
// 自定义界面
view.HandleView((ctx,id,action)=>{
    // 获取视频描述和标题
    let data = relationServer.GetNode(ctx, id, movieInfo)
    if (data.err != null) { return ""}
    let info = data.info
    // 获取所有的季度信息
    let seasons = []
    let res = relationServer.GetNodeChild(ctx, id, nodeSeason, 0, 0, seasonInfo)
    if (res.err != null) { return "" }
    let season = relationServer.NodeRange(res.nodes, seasonInfo)
    season.forEach(data => {
        let videos = []
        // 获取视频信息系
        let res2 = relationServer.GetNodeChild(ctx, data.id, nodeVideo, 0, 0, videoInfo)
        if (res2.err != null) {return}
        relationServer.NodeRange(res2.nodes).forEach(data => {
            data.mp4 = `http://index.xiaoyou.host/app/video/video/${data.id}/mp4`
            videos.push(data)
        })
        // 对视频进行自定义排序
        videos.sort((a,b) => (Number(a.no) - Number(b.no)))
        seasons.push({"id": data.id,"season": data.name,"videos": videos})
    })
    // 获取页面信息
    let page = view.GetView(ctx, 6733, {"name": info.name,"desc": info.desc,"tag": "番剧","thumb": tools.GetObjectLink(info.poster),"detail": seasons})
    return { "page": page }
})