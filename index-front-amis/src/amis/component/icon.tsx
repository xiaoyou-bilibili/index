import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
// 引入lib库
import { library } from '@fortawesome/fontawesome-svg-core'
// 把所有的组件都引入进来
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'
import React, {useState} from "react";
// 添加组件
library.add(fas, far, fab)
// <FontAwesomeIcon icon="fa-solid fa-cloud-arrow-down" />
// {"type":"index-icon","icon":{"type":"fab","name":"google"}}
export const Icon = React.forwardRef((props:any, ref) => {
    let icon = props.icon
    let color = icon.color || '#999'
    let hoverColor = icon.hoverColor || '#999'
    let [current, setColor]= useState(color)
    return  <FontAwesomeIcon
        icon={icon.name}
        size={icon.size || '2x'}
        style={{
            cursor: icon.cursor || "default",
            color: current
        }}
        onMouseLeave={()=>{setColor(color)}}
        onMouseEnter={()=>{setColor(hoverColor)}}
    />;
})
