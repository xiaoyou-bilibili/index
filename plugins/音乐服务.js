const nodeMusic = "music"
const nodeAlbum = "album"

const musicInfo = {name:"", pic:"", audio:"", lrc:"", singer:"", qq_mid:""}
const albumInfo = {name: "", pic:"", desc:""}
const albumMusicReq = {album_id:0, music_id:[]}
const qqMusicInfo = { name: '', song_mid: '', album_id: '', media_id: '', singer: ''}

const api = "http://192.168.1.40:32226"
const getHeader = (ctx) => { return { "Cookie": tools.GetConsulKV(ctx, "index/music/qq_api_cookie") } }

// 获取音乐信息
const getMusicInfo = (ctx, id) => {
    let info = relationServer.GetNode(ctx, id, musicInfo)
    if(info.err !== null){ return {data: null, err: info.err} }
    let data = info.info
    // 手动给内容加上链接
    data.pic = tools.GetObjectLink(data.pic)
    data.audio = tools.GetObjectLink(data.audio)
    // 获取歌词信息
    if(data.lrc !== ""){
        let res = dataServer.GetText(ctx, data.lrc)
        data.lrc=res.err==null?res.content:data.lrc
    }
    return {data, err: null} 
}


// 添加音乐
gin.Handle("POST", "", function (ctx) {
    gin.BindData(ctx, musicInfo,data => {
        relationServer.GinAddNode(ctx, nodeMusic, data, id => mqServer.Add(ctx, id, nodeMusic, data.name, dataServer.GetText(ctx, data.lrc).content, [data.singer]))
    })
})
// 获取音乐信息
gin.Handle("GET", "/:id",function (ctx) {
    gin.GetPathInt(ctx, "id", id => {
        let res = getMusicInfo(ctx, id)
        if (res.err != null) { return gin.Fail(ctx, "get music error %v", res.err)}
        gin.Success(ctx, res.data)
    })
})
// 查找音乐
gin.Handle("GET", "",function (ctx) {
    let info = gin.GetFindField(ctx)
    let data = relationServer.FindNode(ctx, nodeMusic, info.field, info.keyword, info.current, info.size, musicInfo)
    if(data.err!=null){ return  gin.Fail(ctx, "find data error %v", data.err) }
    let res = relationServer.NodeRange(data.nodes, musicInfo)
    res.map(value => value.pic = tools.GetObjectLink(value.pic))
    gin.ReturnPageInfo(ctx, info.current, data.total, res)
})
// 删除音乐
gin.Handle("DELETE", "/:id",function (ctx) {
    let ids = gin.GetPathIntList(ctx, "id")
    ids.forEach(function (value) {
        let info = relationServer.GetNode(ctx, value, musicInfo)
        if(info.err != null) {
            tools.LogCtxError(ctx, "get music error %v", info.err)
        } else {
            dataServer.DeleteObject(ctx, [info.info.pic,info.info.lrc,info.info.audio])
            mqServer.Delete(ctx, value)
        }
    })
    relationServer.DeleteNode(ctx, ids)
    gin.Success(ctx, null)
})
// 添加专辑
gin.Handle("POST", "/album",function (ctx) {
    gin.BindData(ctx, albumInfo, data => {
        relationServer.GinAddNode(ctx, nodeAlbum, data, id => {})
    })
})
// 查找专辑
gin.Handle("GET", "/album",function (ctx) {
    let info = gin.GetFindField(ctx)
    let data = relationServer.FindNode(ctx, nodeAlbum, info.field, info.keyword, info.current, info.size, albumInfo)
    if(data.err!=null){ return gin.Fail(ctx, "find data error %v", data.err) }
    let res = relationServer.NodeRange(data.nodes, albumInfo)
    res.map(value => value.pic = tools.GetObjectLink(value.pic))
    gin.ReturnPageInfo(ctx, info.current, data.total, res)
})
// 删除专辑
gin.Handle("DELETE", "/album/:id",function (ctx) {
    let ids = gin.GetPathIntList(ctx, "id")
    ids.forEach(value => {
        let info = relationServer.GetNode(ctx, value, albumInfo)
        if(info.err != null) {
            tools.LogCtxError(ctx, "get album error %v", info.err)
        } else {
            dataServer.DeleteObject(ctx, [info.info.pic])
        }
    })
    relationServer.DeleteNode(ctx, ids)
    gin.Success(ctx, null)
})
// 专辑添加音乐
gin.Handle("POST", "/album/link",function (ctx) {
    gin.BindData(ctx, albumMusicReq,data => {
        data.music_id.forEach(value => relationServer.AddRelation(ctx, data.album_id, value, nodeMusic, {}))
        gin.Success(ctx, data)
    })
})
// 专辑删除音乐
gin.Handle("DELETE", "/album/link",function (ctx) {
    gin.BindData(ctx, albumMusicReq, data => {
        data.music_id.forEach(value => relationServer.DeleteRelationWithNode(ctx, data.album_id, value))
        gin.Success(ctx, data)
    })
})
// 获取专辑下所有音乐
gin.Handle("GET", "/album/:id",function (ctx) {
    gin.GetPathInt(ctx, "id", id=>{
        let page = gin.GetPageField(ctx)
        let data = relationServer.GetNodeChild(ctx, id, nodeMusic, page.current, page.size, musicInfo)
        if(data.err !== null) { return gin.Fail(ctx, "get node error %v", data.err) }
        gin.ReturnPageInfo(ctx, page.current, data.total ,relationServer.NodeRange(data.nodes, musicInfo))
    })
})
// 获取QQ音乐专辑
gin.Handle("GET", "/qq/album/:id",function (ctx) {
    gin.GetPathInt(ctx, "id", id => {
        let data = tools.HttpSendRequest(ctx, `${api}/songlist?id=${id}`, "GET", getHeader(ctx), null)
        if(data.err != null) { return gin.Fail(ctx, "%v", data.err) }
        let res = []
        data.data.data.songlist.forEach(value=>{
            let singer = value.singer.map(value => {return value.name}).join(",")
            res.push({ name: value.songname, singer: singer, info: { name: value.songname, song_mid: value.songmid, album_id: value.albumid, media_id: value.strMediaMid, singer: singer}})
        })
        gin.ReturnPageInfo(ctx,1, res.length,res)
    })
})
// 获取专辑封面
const getPic = (ctx, id) => tools.HttpSendRequest(ctx, `${api}/album?albummid=${id}`, "GET", getHeader(ctx), null).data.data.picurl
// 获取下载链接
const getDownloadLink = (ctx, id, mid) => {
    let data = tools.HttpSendRequest(ctx, `${api}/song/url?id=${id}&mediaId=${mid}&type=320`, "GET", getHeader(ctx), null)
    return data.data.data === ""?tools.HttpSendRequest(ctx, `${api}/song/url?id=${id}&mediaId=${mid}&type=128`, "GET", getHeader(ctx), null).data.data:data.data.data
}
// 获取歌词
const getLrc = (ctx, id) => tools.HttpSendRequest(ctx, `${api}/lyric?songmid=${id}`, "GET", getHeader(ctx), null).data.data.lyric
// 同步QQ音乐
gin.Handle("POST", "/qq/sync",function (ctx) {
    gin.BindData(ctx, qqMusicInfo, data => {
        let res = relationServer.FindNode(ctx, nodeMusic, "qq_mid", data.song_mid, 0, 0, musicInfo)
        if (res.total > 0) { return gin.Success(ctx, {id: res.nodes[0].Id}) }
        let lrc = getLrc(ctx, data.song_mid)
        let node = relationServer.AddNode(ctx, nodeMusic, { name: data.name, singer: data.singer, qq_mid: data.song_mid,
            pic: dataServer.DownloadLinkAndUpload(ctx, getPic(ctx, data.album_id), `${data.name}.jpg`, {}).id,
            audio: dataServer.DownloadLinkAndUpload(ctx, getDownloadLink(ctx, data.song_mid, data.media_id), `${data.name}.mp3`, {}).id,
            lrc: lrc!==""?dataServer.AddText(ctx,`${data.name}.lrc`,lrc).id:""
        })
        if(node.err != null) {return gin.Fail(ctx, "add node err %v", node.err)}
        mqServer.Add(ctx, node.id, nodeMusic, data.name, lrc, [])
        gin.Success(ctx, {id: node.id})
    })
})

// 处理搜索请求
view.HandleSearch((ctx, id) => {
    let res = getMusicInfo(ctx, id)
    if (res.err != null) { return ""}
    let music = {
        "type": "grid",
        "columns": [
            {"lg":0,"body":{"type":"index-music","audio": res.data}},
            {"lg":8,"body":{"type":"button-group","className":"m-t-lg","buttons":[
                {"icon":"fa fa-star","type":"button","label":"收藏",
                    "actionType": "dialog",
                    "dialog": {"title":"提示","actions":[{"label":"提交","primary":true,"type":"button","close":true,"onEvent":{"click":{"actions":[{"actionType":"index-request","args":{"url":"/app/music/album/link","method":"post","numberField": ["album_id"],"data":{"music_id":[id], "album_id":"${album_id}"}}}]}}}],"body":{"type":"form","body":[{"type":"input-text","name":"album_id","required":true,"placeholder":"请输入专辑ID","label":"专辑ID"}]}}
                },
                {"icon":"fa fa-cloud-download","type":"button","label":"下载","onEvent":{"click":{"actions":[{"actionType":"index-download","args":{ "url": res.data.audio}}]}}}
            ]}}
        ]
    }
    return JSON.stringify(music)
})
// 注册管理界面
view.RegisterManage("音乐管理", 6732, ()=>{})
view.RegisterManage("专辑管理", 6740, ()=>{})