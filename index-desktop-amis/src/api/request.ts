import axios, {AxiosResponse} from 'axios'
import { message } from 'antd'


// 对axios函数进行封装，用来发api请求，post使用qs进行处理，避免自己把from数据转换为json字符串
export default async function request (url:string, data:any, type:string) {
  let req:any
  // 判断请求类型
  if (type === 'get') {
    req = axios.get(url, { params: data, timeout: 1000 * 60 * 10 })
  } else if (type === 'post') {
    req = axios.post(url, data)
  } else if (type === 'put') {
    req = axios.put(url, data)
  } else if (type === 'delete') {
    req = axios.delete(url, {params: data})
  } else if (type === 'patch') {
    req = axios.patch(url, data)
  }
  return new Promise((resolve, reject) => {
    req.then((res: AxiosResponse) => {
      if (res.status !== 200) {
        message.error('请求失败')
        reject('请求失败')
      }else if (res.data.code !== 200) {
        message.error(res.data.msg)
        reject(res.data.msg)
      } else {
        resolve(res.data.data)
      }
    })
  })
}
