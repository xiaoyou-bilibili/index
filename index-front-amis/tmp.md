// 图标按钮
```json
{
    "type": "index-icon",
    "icon": {
    "type": "fab",
        "name": "google",
        "hoverColor": "red",
        "cursor": "pointer"
    }
}
```

// 音乐播放器组件
```json
{
  "type": "page",
  "body": {
    "type": "grid",
    "columns": [
      {"lg":0,"body":{"type":"index-music","audio":{"audio":"https://index.xiaoyou.host/data/object/62ee8d7a0d999bc677e8753c","lrc":"[ti:晚安喵]\n[ar:艾索]\n[al:罗小黑战记系列歌曲]\n[by:]\n[offset:0]\n[00:00.00]晚安喵 - 艾索 (Iso)\n[00:00.20]词：薄荷映像\n[00:00.41]曲：薄荷映像\n[00:00.62]编曲：薄荷映像\n[00:00.83]早安喵 午安喵 晚安喵 喵喵\n[00:06.06]早安喵 午安喵 晚安喵 喵喵\n[00:20.88]喜欢你的微笑和调皮的嘴角\n[00:25.85]那午后的阳光穿过你的发梢\n[00:30.93]想让全世界停在这一秒\n[00:36.34]看着你把世界都忘掉\n[00:41.38]早安喵 午安喵 晚安喵 喵喵\n[00:46.37]早安喵 午安喵 晚安喵 喵喵\n[00:55.40]Yo Yo\n[01:00.05]Check Check it out\n[01:11.36]喜欢你的微笑和调皮的嘴角\n[01:16.33]喜欢你的拥抱和黄色外套\n[01:21.06]这甜蜜的感觉只有我知道\n[01:27.14]就是喜欢你的味道 奥 奥 奥\n[01:31.94]早安喵 午安喵 晚安喵 喵喵\n[01:36.84]早安喵 午安喵 晚安喵 喵喵\n[01:45.46]嘿咻嘿咻","name":"晚安喵1","pic":"https://index.xiaoyou.host/data/object/62ee07e00d999bc677e86381","qq_mid":"001TJiEe2IT9a7","singer":"艾索1111123"}}},
      {
        "lg": 8,
        "body": {
          "type": "button-group",
          "className": "m-t-lg",
          "buttons": [
            {
              "icon": "fa fa-star",
              "type": "button",
              "label": "收藏",
              "actionType": "dialog",
              "dialog": {"title":"提示","actions":[{"label":"提交","primary":true,"type":"button","close":true,"onEvent":{"click":{"actions":[{"actionType":"index-request","args":{"url":"http://index.xiaoyou.host/app/music/album","method":"post","data":{"music_id":"123","album_id":"${album_id}"}}}]}}}],"body":{"type":"form","body":[{"type":"input-text","name":"album_id","required":true,"placeholder":"请输入专辑ID","label":"专辑ID"}]}}
            },
            {
              "icon": "fa fa-cloud-download",
              "type": "button",
              "label": "下载"
            }
          ]
        }
      }
    ]
  }
}
```

## 番剧
```json
{
  "type": "page",
  "body": [
    {
      "type": "flex",
      "items": [
        {"type":"image","width":"80px","height":"150px","imageMode":"original","src":"https://index.xiaoyou.host/data/object/62e62ca8a18ac345cff5fb10"},
        {"type": "flex", "direction": "column", "alignItems": "flex-start", "className": "m-l", "items": [
            {"type":"flex","justify":"flex-start","items":[{"type":"tag","label":"番剧","displayMode":"normal","color":"active"},{"type":"html","html":"<div style=\"font-weight: 600;font-size:18px\">NEW GAME!</div>"}]},
            {"type":"html","html":"<div style=\"font-size: 13px;line-height:18px;margin-bottom: 4px;color: #9499A0;height: 90px;overflow: scroll;overflow-x: hidden;\">高中毕业后，进入了曾制作自己自幼便入迷的游戏的制作公司“Eagle Jump”的青叶，在那里与担任这款游戏角色设计师的八神光相遇了。 开始在憧憬的人手下工作的青叶，虽然对于第一次的工作感到困惑，但在以光为首的充满个性的前辈社员的帮助下一点一点地成长着。 描绘在游戏公司工作的女孩子们的日常的工作女孩喜剧，现在开幕！</div>"}
          ]
        }
      ]
    }
  ]
}
```

## 文件下载
```json
{
  "type": "page",
  "body": [
    {
      "type": "grid",
      "columns": [
        {"md":1,"body":[{"type":"image","width":"65px","height":"65px","imageMode":"original","src":"https://index.xiaoyou.host/data/object/630077a48668adf7ba1b9c24"}]},
        {"md":8,"body":[{"type":"html","html":"<div style=\"font-weight:800\">毕业证书.pdf</div><div style=\"color: #333;font-size: 13px;\">大小：100k</div><div style=\"font-size: 13px;color: #9195A3;\">修改日期：2020-2-1</div>"}]},
        {"md":2,"body":[{"type":"button-group","vertical":true,"buttons":[{"icon":"fa fa-clipboard","type":"button","label":"复制链接"},{"icon":"fa fa-cloud-download","type":"button","label":"下载文件"}]}]}
      ]
    }
  ]
}
```

## 博客
```json
{
  "type": "page",
  "body": [
    {
      "type": "grid",
      "columns": [
        {
          "lg": 4,
          "body": [{"type":"image","width":"60px","height":"60px","imageMode":"original","src":"https://index.xiaoyou.host/data/object/630077a48668adf7ba1b9c24"}]
        },
        {
          "lg": 4,
          "body": [
            {
              "type": "html",
              "html": "<h2 style=\"margin-top:5px;border-bottom:1px solid #bbb;font-size:16px;font-weight:600;\">你好</h2><p style=\"font-size: 13px;color: #86909c\">啊啊啊啊啊</p>"
            }
          ]
        },
        {
          "lg": 4,
          "body": [{
            "type": "button-group",
            "className": "m-t-lg",
            "buttons": [
              {
                "icon": "fa fa-star",
                "type": "button",
                "label": "收藏"
              },
              {
                "icon": "fa fa-cloud-download",
                "type": "button",
                "label": "下载"
              }
            ]
          }]
        }
      ]
    }
  ]
}
```

## 音乐管理
```json
{
  "type": "page",
  "data": {
    "host":"http://index.xiaoyou.host",
    "upload-url": "http://index.xiaoyou.host/data/object"
  },
  "body": [
    {
      "type": "crud",
      "api": "${host}/app/music?search_type=${search_type || 'name'}&search_keyword=${search_keyword}&size=${perPage}&current=${page}",
      "syncLocation": false,
      "source": "list",
      "filter": {
        "title": "条件查询",
        "body": [
          {
            "type": "flex",
            "justify": "flex-start",
            "alignItems": "flex-start",
            "items": [
              {
                "label": "选项",
                "type": "select",
                "name": "search_type",
                "options": [
                  {
                    "label": "名字",
                    "value": "name"
                  },
                  {
                    "label": "歌手",
                    "value": "singer"
                  }
                ]
              },
              {
                "label": "关键词",
                "type": "input-text",
                "name": "search_keyword",
                "placeholder": "请输入关键词"
              }
            ]
          }
        ]
      },
      "headerToolbar": [
        {
          "type": "button",
          "label": "新增音乐",
          "level": "primary",
          "icon": "fa fa-music",
          "actionType": "dialog",
          "dialog": {
            "title": "新增音乐",
            "body": {
              "type": "form",
              "api": "post:${host}/app/music",
              "body": [
                {
                  "type": "input-text",
                  "name": "name",
                  "label": "名字"
                },
                {
                  "type": "input-text",
                  "name": "singer",
                  "label": "歌手"
                },
                {
                  "type": "input-file",
                  "name": "pic",
                  "label": "封面",
                  "accept": "*",
                  "receiver": "${upload-url}",
                  "urlField": "object_id",
                  "valueField": "object_id"
                },
                {
                  "type": "input-file",
                  "name": "audio",
                  "label": "音频",
                  "accept": "*",
                  "receiver": "${upload-url}",
                  "urlField": "object_id",
                  "valueField": "object_id"
                },
                {
                  "type": "input-file",
                  "name": "lrc",
                  "label": "歌词",
                  "accept": "*",
                  "receiver": "${upload-url}",
                  "urlField": "object_id",
                  "valueField": "object_id"
                }
              ]
            }
          }
        },
        "bulkActions"
      ],
      "bulkActions": [
        {
          "label": "添加到歌单",
          "level": "success",
          "icon": "fa fa-plus",
          "actionType": "dialog",
          "dialog": {
            "title": "提示",
            "actions": [
              {
                "label": "提交",
                "primary": true,
                "type": "button",
                "close": true,
                "onEvent": {
                  "click": {
                    "actions": [
                      {
                        "actionType": "index-request",
                        "args": {
                          "url": "${host}/app/music/album",
                          "method": "post",
                          "numberField": [
                            "album_id"
                          ],
                          "numberListField": [
                            "music_id"
                          ],
                          "data": {
                            "music_id": "${ids|raw}",
                            "album_id": "${album_id}"
                          }
                        }
                      }
                    ]
                  }
                }
              }
            ],
            "body": {
              "type": "form",
              "body": [
                {
                  "type": "input-text",
                  "name": "album_id",
                  "required": true,
                  "placeholder": "请输入专辑ID",
                  "label": "专辑ID"
                }
              ]
            }
          }
        },
        {
          "label": "从歌单中删除",
          "icon": "fa fa-chain-broken",
          "level": "warning",
          "actionType": "dialog",
          "dialog": {
            "title": "提示",
            "actions": [
              {
                "label": "提交",
                "primary": true,
                "type": "button",
                "close": true,
                "onEvent": {
                  "click": {
                    "actions": [
                      {
                        "actionType": "index-request",
                        "args": {
                          "url": "${host}/app/music/album",
                          "method": "post",
                          "numberField": [
                            "album_id"
                          ],
                          "numberListField": [
                            "music_id"
                          ],
                          "data": {
                            "music_id": "${ids|raw}",
                            "album_id": "${album_id}"
                          }
                        }
                      }
                    ]
                  }
                }
              }
            ],
            "body": {
              "type": "form",
              "body": [
                {
                  "type": "input-text",
                  "name": "album_id",
                  "required": true,
                  "placeholder": "请输入专辑ID",
                  "label": "专辑ID"
                }
              ]
            }
          }
        },
        {
          "label": "批量删除",
          "icon": "fa fa-trash",
          "level": "danger",
          "actionType": "ajax",
          "api": "delete:${host}/app/music/${ids|raw}",
          "confirmText": "确定要批量删除?"
        }
      ],
      "columns": [
        {
          "name": "id",
          "label": "ID"
        },
        {
          "name": "pic",
          "label": "封面",
          "quickEdit": {
            "mode": "inline",
            "type": "html",
            "html": "<img style=\"width: 50px;border-radius: 5px;\"  src=\"${pic}\"/>"
          }
        },
        {
          "name": "name",
          "label": "名称"
        },
        {
          "name": "singer",
          "label": "歌手"
        }
      ]
    }
  ]
}
```

## 专辑管理
```json

```
