<template>
  <!--歌单详情-->
  <div class="details">
    <mm-loading v-model="mmLoadShow" />
    <music-list list-type="pullup" :list="list" @select="selectItem" @pullUp="getNext" />
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import { getPlaylistDetail } from 'api'
import MmLoading from 'base/mm-loading/mm-loading'
import MusicList from 'components/music-list/music-list'
import { loadMixin } from '@/utils/mixin'

export default {
  name: 'Detail',
  components: {
    MmLoading,
    MusicList
  },
  mixins: [loadMixin],
  data() {
    return {
      list: [], // 列表
      current: 1
    }
  },
  created() {
    this.getMusicList(this.current)
  },
  methods: {
    // 播放暂停事件
    selectItem(item, index) {
      this.selectPlay({
        list: this.list,
        index
      })
    },
    // 获取音乐列表
    getMusicList(current) {
      // 获取歌单详情
      getPlaylistDetail(this.$route.params.id, current)
        .then(data => {
          if (data.list === null || data.list === undefined) {
            this.$mmToast('没有更多歌曲啦！')
          }
          if (data.current === 1) {
            this.list = data.list
          } else {
            this.list = [...this.list, ...data.list]
          }
          this._hideLoad()
        }).catch(() => {
          this._hideLoad()
        })
    },
    // 获取下一页
    getNext() {
      this.current++
      this.getMusicList(this.current)
    },
    ...mapActions(['selectPlay'])
  }
}
</script>

<style lang="less" scoped>
.details {
  position: relative;
  width: 100%;
  height: 100%;
  .musicList {
    width: 100%;
    height: 100%;
  }
}
</style>
