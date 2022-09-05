import axios from '@/utils/axios'

axios.defaults.baseURL = process.env.VUE_APP_BASE_API_URL

// 排行榜列表
export function getToplistDetail() {
  return axios.get('/toplist/detail')
}

// 推荐歌单
export function getPersonalized() {
  return axios.get('/personalized')
}

// 获取某个歌单详情
export function getPlaylistDetail(id, current) {
  return new Promise((resolve, reject) => {
    axios
      .get(`/app/music/album/${id}?current=${current}&size=30`)
      .then(res => {
        if (res.code !== 200) {
          reject(res.msg)
        }
        resolve(res.data)
      })
  })
}

// 搜索音乐
// eslint-disable-next-line camelcase
export function search(search_type, search_keyword, current) {
  return axios.get('/app/music', { params: { search_type, search_keyword, current, size: 300 } })
}

// 获取所有的专辑
export function getUserPlaylist(q) {
  return axios.get('/app/music/album', { params: { search_type: 'name' } })
}

// 获取音乐信息
export function getLMusic(id) {
  const url = `/app/music/${id}`
  return axios.get(url)
}

// 专辑添加音乐
export function AlbumAddMusic(data) {
  const url = `/app/music/album/link`
  return axios.post(url, JSON.stringify(data), {
    headers: { 'Content-Type': 'application/json' }
  })
}

// 专辑删除音乐
export function AlbumDeleteMusic(data) {
  const url = `/app/music/album/link`
  return axios.delete(url, {
    data: JSON.stringify(data),
    headers: { 'Content-Type': 'application/json' }
  })
}
