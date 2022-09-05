import * as React from 'react';
// @ts-ignore
import Aplayer from 'aplayer';
import DPlayer  from 'dplayer'
import {useEffect, useState} from "react";
import 'aplayer/dist/APlayer.min.css';

// {"type":"index-music","audio":{"audio":"https://index.xiaoyou.host/data/object/62ee8d7a0d999bc677e8753c","lrc":"[ti:晚安喵]\n[ar:艾索]\n[al:罗小黑战记系列歌曲]\n[by:]\n[offset:0]\n[00:00.00]晚安喵 - 艾索 (Iso)\n[00:00.20]词：薄荷映像\n[00:00.41]曲：薄荷映像\n[00:00.62]编曲：薄荷映像\n[00:00.83]早安喵 午安喵 晚安喵 喵喵\n[00:06.06]早安喵 午安喵 晚安喵 喵喵\n[00:20.88]喜欢你的微笑和调皮的嘴角\n[00:25.85]那午后的阳光穿过你的发梢\n[00:30.93]想让全世界停在这一秒\n[00:36.34]看着你把世界都忘掉\n[00:41.38]早安喵 午安喵 晚安喵 喵喵\n[00:46.37]早安喵 午安喵 晚安喵 喵喵\n[00:55.40]Yo Yo\n[01:00.05]Check Check it out\n[01:11.36]喜欢你的微笑和调皮的嘴角\n[01:16.33]喜欢你的拥抱和黄色外套\n[01:21.06]这甜蜜的感觉只有我知道\n[01:27.14]就是喜欢你的味道 奥 奥 奥\n[01:31.94]早安喵 午安喵 晚安喵 喵喵\n[01:36.84]早安喵 午安喵 晚安喵 喵喵\n[01:45.46]嘿咻嘿咻","name":"晚安喵1","pic":"https://index.xiaoyou.host/data/object/62ee07e00d999bc677e86381","qq_mid":"001TJiEe2IT9a7","singer":"艾索1111123"}}
// {"type":"index-video","media":{"name":"乐器！","no":"02","thumb":"https://index.xiaoyou.host/data/object/62db4d33d70806f28ecc17c4","mp4":"http://192.168.1.40:32561/video/62db5680cddbe2d9534a63a8?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=D8E4K4KAKX71QO2X79IG%2F20220818%2Findex-data-obj%2Fs3%2Faws4_request&X-Amz-Date=20220818T140304Z&X-Amz-Expires=3600&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3B%20filename%3D%22%E8%BD%BB%E9%9F%B3%E5%B0%91%E5%A5%B3%20%282009%29%2FSeason%201%2FS01E02-%E4%B9%90%E5%99%A8%EF%BC%81.mp4%22&X-Amz-Signature=237493e95dd84bf892c0140f7ceb92d22a44aa124eebf8f2b01d0356a4288b8c","barrage":"https://index.xiaoyou.host/data/object/62db5528cddbe2d9534a6332","subtitle":"https://index.xiaoyou.host/data/object/","nfo":"62db5529cddbe2d9534a6333"}}

export const Music = React.forwardRef((props:any, ref) => {
  let id = (new Date()).valueOf().toString()
  let [option,setOption] = useState({})

  useEffect(() => {
    let { audio } = props
    let ap = new Aplayer({
      container: document.getElementById(id),
      lrcType: 1,
      audio: [{name: audio.name, artist: audio.singer, url: audio.audio, cover: audio.pic, lrc: audio.lrc, theme: '#ebd0c2'}]
    })
  }, [props])

  return  <div id={id} />;
})



export const Video = React.forwardRef((props:any, ref) => {
  let id = (new Date()).valueOf().toString()
  // 只要参数变化就重新渲染
  useEffect(() => {
    let dp = new DPlayer({
      container: document.getElementById(id),
      video: { url: props.video, pic: props.thumb}
    });
    dp.play()
  }, [props])

  return  <div id={id} />;
})
