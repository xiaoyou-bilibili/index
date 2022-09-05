import * as React from "react"
import {useEffect, useState} from "react"
import { marked } from 'marked'
import hljs from 'highlight.js'
import 'highlight.js/styles/foundation.css'
import {Renderer, ScopedContext, setVariable} from 'amis';


// markdown编辑器
import Vditor from 'vditor'
import "vditor/dist/index.css"
import {message} from "antd";


const render = new marked.Renderer()
marked.setOptions({
  renderer: render, // 这是必填项
  gfm: true,	// 启动类似于Github样式的Markdown语法
  pedantic: false, // 只解析符合Markdwon定义的，不修正Markdown的错误
  sanitize: false, // 原始输出，忽略HTML标签（关闭后，可直接渲染HTML标签）
  // 高亮的语法规范
  highlight: (code:string, lang:string) => hljs.highlight(code, { language: lang }).value,
})

// markdown显示
export const Markdown = React.forwardRef((props:any, ref) => {
  let [content,setContent] = useState('')

  useEffect(() => {
    let { content } = props
    setContent(props.content || '')
  }, [props])

  return  <div dangerouslySetInnerHTML={{__html: marked.parse(content)}}></div>;
})


// markdown编辑器
export const MarkdownEdit = React.forwardRef((props:any, ref) => {
  let id = (new Date()).valueOf().toString()
  let editor:Vditor

  useEffect(() => {
    let { content, upload, onBulkChange } = props
    if (content === undefined) {
      return
    }
    if (editor != null) {
      editor.setValue(content)
    } else {
      editor = new Vditor(id, {
        height: 600,
        toolbarConfig: {pin: true},
        cache: {enable: true},
        // 图片上传处理事件
        upload: { url: upload,  fieldName: "file",  success(_, msg: string) {
            const data = JSON.parse(msg)
            data.code !== 200? message.error(String(data.msg)) : editor.insertValue(`![](${data.data.url})`)
          }
        },
        after: () => {
          editor.setValue(content)
          onBulkChange({"editor": content})
        },
        input: (value: string) => onBulkChange({"editor": editor.getValue()})
      });
    }
  }, [props.content])

  return  <div><div id={id}/></div>;
})
