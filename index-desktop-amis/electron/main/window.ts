import {app, BrowserWindow, shell, dialog, ipcMain} from "electron";
// @ts-ignore
import path, {join} from "path";
import {ROOT_PATH} from "./common";

const preload = join(__dirname, '../preload/index.js')
// 设置开发环境和正式环境的页面，
const url = `http://127.0.0.1:${process.env['VITE_DEV_SERVER_PORT']}`
const indexHtml = join(ROOT_PATH.dist, 'index.html')

// @ts-ignore
export function openWindow(info:WindowsSendMessage):BrowserWindow {
  // 默认就加载我们的脚本
  let option:Electron.BrowserWindowConstructorOptions = {
      webPreferences: {preload},
      width: info.width,
      height: info.height,
      frame: info.frame,
      icon: info.logo,
      resizable: info.resize,
      title: info.title,
  }
  let saveUrl =  ''

  const childWindow = new BrowserWindow(option)
  if (app.isPackaged) {
      childWindow.loadFile(indexHtml, { hash:  info.url}).then(() => childWindow.show())
  } else {
      childWindow.loadURL(`${url}/#${info.url}`).then(() => childWindow.show())
      if (info.dev !== undefined && info.dev) {
          childWindow.webContents.openDevTools({ mode: "undocked", activate: true })
      }
  }

  // 监听渲染进程发出的download事件
  ipcMain.on('download', async (evt, args) => {
    // 打开系统弹窗 选择文件下载位置
    // @ts-ignore
    dialog.showOpenDialog({properties: ['openFile', 'openDirectory']}, (files) => {}).then(res => {
      if (res.canceled) return; // 如果用户没有选择路径,则不再向下进行
      saveUrl = res.filePaths[0];  // 保存文件路径
      if (!saveUrl) return; // 如果用户没有选择路径,则不再向下进行
      childWindow.webContents.downloadURL(args); // 触发 will-download 事件
    })
  });

  // 监听文件下载事件
  childWindow.webContents.session.on('will-download', (e, item) => {
    const filePath = path.join(saveUrl, item.getFilename());
    item.setSavePath(filePath);
    //监听下载过程，计算并设置进度条进度
    item.on('updated', (evt, state) => {
      if ('progressing' === state) {
        let value = 0
        //此处  用接收到的字节数和总字节数求一个比例  就是进度百分比
        if (item.getReceivedBytes() && item.getTotalBytes()) {
          value = parseInt(String(100 * (item.getReceivedBytes() / item.getTotalBytes())))
        }
        // 把百分比发给渲染进程进行展示
        childWindow.webContents.send('updateProgressing', value);
        // mac 程序坞、windows 任务栏显示进度
        childWindow.setProgressBar(value);
      }
    });
    //监听下载结束事件
    item.on('done', (e, state) => {
      //如果窗口还在的话，去掉进度条
      if (!childWindow.isDestroyed()) {
        childWindow.setProgressBar(-1);
      }
      //下载被取消或中断了
      if (state === 'interrupted') {
        dialog.showErrorBox('下载失败', `文件 ${item.getFilename()} 因为某些原因被中断下载`);
      }
      // 下载成功后打开文件所在文件夹
      if (state === 'completed') {
        setTimeout(() => {
          shell.showItemInFolder(filePath)
        }, 1000);
      }
    });
  });

  // 把当前窗口返回
  return childWindow
}
