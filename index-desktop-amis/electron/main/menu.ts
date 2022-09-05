// 快速打开窗口
import {app, BrowserWindow, Menu, Tray} from "electron";
import {join} from "path";
import {ROOT_PATH} from "./common";

// 三方应用窗口
let appWin:Map<string,BrowserWindow> = new Map()

// 快速对窗口进行操作
const windowOption = (name:string, width:number, height:number, icon: string, url: string) => {
    // 先判断map是否存在
    if (!appWin.has(name) || appWin.get(name) == null || appWin.get(name)?.isDestroyed()) {
        let win:BrowserWindow  = new BrowserWindow({ width: width, height: height, icon: join(ROOT_PATH.public, 'static/icon/'+icon) })
        win.loadURL(url).then(()=> win.show())
        appWin.set(name, win)
    } else {
        // 如果存在就隐藏，否则显示
        appWin.get(name)?.isVisible()? appWin.get(name)?.hide():appWin.get(name)?.show()
    }
}

// 菜单栏设置
export function setupMenu(handle:()=>void){
    // 显示系统托盘
    let icon = new Tray(join(ROOT_PATH.public, 'static/image/logo.png'))
    // 右键菜单
    const contextMenu = Menu.buildFromTemplate([
        { label: '主界面', type: 'normal', click: handle},
        { label: '音乐', type: 'normal', click: ()=> windowOption('music', 1000,700,"music.png", "http://music.xiaoyou.host")},
        { label: '网盘', type: 'normal', click: ()=> windowOption('cloud', 1200, 700 ,"cloud.png", "http://cloud.xiaoyou.host")},
        { label: '后台', type: 'normal', click: ()=> windowOption('cloud', 1200, 700 ,"admin.png", "http://cloud.xiaoyou.host")},
    ])
    // 给我们的托盘图标设置系统菜单
    icon.setContextMenu(contextMenu)
    icon.setToolTip("index")
    // 设置点击事件
    icon.addListener("click", ()=> contextMenu.popup() )
}
