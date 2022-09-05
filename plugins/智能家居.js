const unique = "home"
let token = tools.GetConsulKV(tools.ContextBackground(), "index/app/ha_token")

// LED控制
gin.Handle("GET", "/server/:domain/:service/:entity", ctx => {
    let res = tools.HttpSendRequest(ctx, `http://ha.xiaoyou.host/api/services/${ctx.Param("domain")}/${ctx.Param("service")}`, "POST", {"Authorization": `Bearer ${token}`}, {"entity_id": ctx.Param("entity")})
    // 添加粉丝
    gin.Success(ctx, res)
})
// 全量数据同步
gin.Handle("GET", "/sync", ctx => {
    // 先把旧记录全部删除
    mqServer.DeleteAll(ctx, unique)
    mqServer.Add(ctx, -2, unique, "灯", "", [])
    // 添加粉丝
    gin.Success(ctx, "同步成功")
})
// B站粉丝数卡片
view.HandleSearch((ctx,id) => {
    let card = {
        "type": "service",
        "data": { "host": "http://index.xiaoyou.host" },
        "body": {
        "type": "flex",
        "justify": "flex-start",
        "items": [
            {"type":"image","width":"40px","height":"40px","imageMode":"original","src":"${host}/data/object/6309d715e891a61a9478303a"},
            {"className": "m-l","type": "button-group-select", "name": "type", "options": [{"label": "打开","value": "turn_on"},{"label": "关闭","value": "turn_off"}],
                "onEvent": {"change": {"actions": [{"actionType": "index-request", "args": {"url":"/app/home/server/light/${event.data.value}/light.mbulb3_cloud_128410","method":"get"}}]}}
            }
        ]}
    }
    // 返回内容
    return JSON.stringify(card)
})
