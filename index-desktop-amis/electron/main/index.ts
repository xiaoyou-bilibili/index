import { app, BrowserWindow, Menu, ipcMain,globalShortcut  } from 'electron'
import { release } from 'os'
import {setupMenu} from "./menu";
import {openWindow} from "./window";

// 关闭报警
process.env['ELECTRON_DISABLE_SECURITY_WARNINGS']=String(true)
// Disable GPU Acceleration for Windows 7
if (release().startsWith('6.1')) app.disableHardwareAcceleration()

// Set application name for Windows 10+ notifications
if (process.platform === 'win32') app.setAppUserModelId(app.getName())

if (!app.requestSingleInstanceLock()) {
  app.quit()
  process.exit(0)
}

let win: BrowserWindow | null = null

async function createWindow() {
  win = openWindow({url:"/search", width:800,height:600,dev:true,resize:false})
  // 默认隐藏当前界面
  win.hide()
  // 打开主界面
  // openWindow({url:"", logo: join(ROOT_PATH.public, 'static/image/logo.png'), frame: true})
  // 隐藏默认的菜单栏
  Menu.setApplicationMenu(null);
  // 设置底部菜单栏
  setupMenu(() => {win?.isVisible()?win.hide():win?.show()})
  // 设置全局快捷键
  globalShortcut.register('alt+Space', () => {
    win?.isVisible()?win.hide():win?.show()
  })
}

app.whenReady().then(createWindow)

app.on('window-all-closed', () => {
  win = null
  if (process.platform !== 'darwin') app.quit()
})

app.on('second-instance', () => {
  if (win) {
    // Focus on the main window if the user tried to open another
    if (win.isMinimized()) win.restore()
    win.focus()
  }
})

app.on('activate', () => {
  const allWindows = BrowserWindow.getAllWindows()
  if (allWindows.length) {
    allWindows[0].focus()
  } else {
    createWindow()
  }
})

// 打开一个新窗口
// @ts-ignore
ipcMain.on('open-win', (event, info:WindowsSendMessage) => openWindow(info))
