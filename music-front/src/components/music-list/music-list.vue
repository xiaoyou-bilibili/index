<template>
  <!--歌曲列表-->
  <div class="musicList">
    <template v-if="list.length > 0">
      <mm-dialog
        ref="loginDialog"
        head-text="提示"
        confirm-btn-text="确定"
        cancel-btn-text="关闭"
        @confirm="addFavorite"
      >
        <div class="mm-dialog-text">
          <input
            v-model.trim="albumID"
            class="mm-dialog-input"
            type="number"
            autofocus
            placeholder="请输入专辑ID"
            @keyup.enter="addFavorite"
          />
        </div>
      </mm-dialog>
      <div class="list-item list-header">
        <span class="list-name">歌曲</span>
        <span class="list-artist">歌手</span>
      </div>
      <div ref="listContent" class="list-content" @scroll="listScroll($event)">
        <div
          v-for="(item, index) in list"
          :key="item.id"
          class="list-item"
          :class="{ on: playing && currentMusic.id === item.id }"
          @dblclick="selectItem(item, index, $event)"
        >
          <span class="list-num" v-text="index + 1"></span>
          <div class="list-name">
            <span>{{ item.name }}</span>
            <div class="list-menu">
              <mm-icon
                class="hover"
                :type="getPlayIconType(item)"
                :size="40"
                @click.stop="selectItem(item, index)"
              />
              <!-- 添加到专辑按钮-->
              <mm-icon
                class="hover"
                type="ic_favorite"
                :size="40"
                @click.stop="openDialog(0, 'add', item.id)"
              />
              <!-- 从专辑中删除-->
              <mm-icon
                class="hover"
                type="delete-mini"
                :size="40"
                @click.stop="openDialog(0, 'delete', item.id)"
              />
            </div>
          </div>
          <span class="list-artist">{{ item.singer }}</span>
        </div>
        <slot name="listBtn"></slot>
      </div>
    </template>
    <mm-no-result v-else title="内容为空~" />
  </div>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
import { format } from '@/utils/util'
import { AlbumAddMusic, AlbumDeleteMusic } from '@/api'
import MmNoResult from 'base/mm-no-result/mm-no-result'
import MmDialog from 'base/mm-dialog/mm-dialog'

const LIST_TYPE_ALBUM = 'album'
const LIST_TYPE_PULLUP = 'pullup'

// 触发滚动加载的阈值
const THRESHOLD = 10

export default {
  name: 'MusicList',
  components: {
    MmNoResult,
    MmDialog
  },
  filters: {
    format
  },
  props: {
    // 歌曲数据
    list: {
      type: Array,
      default: () => []
    },
    /**
     * 列表类型
     * album: 显示专辑栏目（默认）
     * duration: 显示时长栏目
     * pullup: 开启上拉加载
     */
    listType: {
      type: String,
      default: LIST_TYPE_ALBUM
    }
  },
  data() {
    return {
      lockUp: true, // 是否锁定滚动加载事件,默认锁定
      albumID: '', // 专辑ID
      musicId: 0, // 当前音乐ID
      option: 'add' // 默认操作为添加
    }
  },
  computed: {
    ...mapGetters(['playing', 'currentMusic'])
  },
  watch: {
    list(newList, oldList) {
      if (this.listType !== LIST_TYPE_PULLUP) {
        return
      }
      if (newList.length !== oldList.length) {
        this.lockUp = false
      } else if (
        newList[newList.length - 1].id !== oldList[oldList.length - 1].id
      ) {
        this.lockUp = false
      }
    }
  },
  activated() {
    this.scrollTop &&
      this.$refs.listContent &&
      (this.$refs.listContent.scrollTop = this.scrollTop)
  },
  methods: {
    // 滚动事件
    listScroll(e) {
      const scrollTop = e.target.scrollTop
      this.scrollTop = scrollTop
      if (this.listType !== LIST_TYPE_PULLUP || this.lockUp) {
        return
      }
      const { scrollHeight, offsetHeight } = e.target
      if (scrollTop + offsetHeight >= scrollHeight - THRESHOLD) {
        this.lockUp = true // 锁定滚动加载
        this.$emit('pullUp') // 触发滚动加载事件
      }
    },
    addFavorite() { // 添加到我的收藏
      if (this.albumID === '') {
        this.$mmToast('专辑 ID不能为空')
        this.openDialog(0, '')
        return
      }
      this.$mmToast(this.albumID)
      // 判断操作类型
      if (this.option === 'add') {
        AlbumAddMusic({
          'album_id': Number(this.albumID),
          'music_id': [this.musicId]
        }).then((res) => {
          if (res.code === 200) {
            this.$mmToast('收藏成功！')
          }
        })
      } else {
        AlbumDeleteMusic({
          'album_id': Number(this.albumID),
          'music_id': [this.musicId]
        }).then((res) => {
          if (res.code === 200) {
            this.$mmToast('删除成功，重新加载后生效！')
          }
        })
      }
    },
    openDialog(key, option, id = 0) {
      this.option = option
      if (id !== 0) {
        this.musicId = id
      }
      this.$refs.loginDialog.show()
    },
    // 播放暂停事件
    selectItem(item, index, e) {
      if (e && /list-menu-icon-del/.test(e.target.className)) {
        return
      }
      if (this.currentMusic.id && item.id === this.currentMusic.id) {
        this.setPlaying(!this.playing)
        return
      }

      /**
       * 为了修复 safari、 ios 微信、安卓 UC 无法播放问题，暂时移除接口校验直接播放
       */
      this.$emit('select', item, index) // 触发点击播放事件
    },
    // 获取播放状态 type
    getPlayIconType({ id: itemId }) {
      const {
        playing,
        currentMusic: { id }
      } = this
      return playing && id === itemId ? 'pause-mini' : 'play-mini'
    },
    // 删除事件
    deleteItem(index) {
      this.$emit('del', index) // 触发删除事件
    },
    ...mapMutations({
      setPlaying: 'SET_PLAYING'
    })
  }
}
</script>

<style lang="less" scoped>
.list-header {
  border-bottom: 1px solid @list_head_line_color;
  color: @text_color_active;

  .list-name {
    padding-left: 40px;
    user-select: none;
  }
}

.list-content {
  width: 100%;
  height: calc(~'100% - 60px');
  overflow-x: hidden;
  overflow-y: auto;
  -webkit-overflow-scrolling: touch;
}

.list-no {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  color: @text_color;
}

.list-item {
  display: flex;
  width: 100%;
  height: 50px;
  border-bottom: 1px solid @list_item_line_color;
  line-height: 50px;
  overflow: hidden;

  &.list-item-no {
    justify-content: center;
    align-items: center;
  }

  &.on {
    color: #fff;

    .list-num {
      font-size: 0;
      background: url('~assets/img/wave.gif') no-repeat center center;
    }
  }

  &:hover {
    .list-name {
      padding-right: 80px;

      .list-menu {
        display: block;
      }
    }
  }

  &:not([class*='list-header']):hover {
    .list-name {
      padding-right: 80px;

      .list-menu {
        display: block;
      }
    }

    .list-time {
      font-size: 0;

      .list-menu-icon-del {
        display: block;
      }
    }
  }

  .list-num {
    display: block;
    width: 30px;
    margin-right: 10px;
    text-align: center;
  }

  .list-name {
    position: relative;
    flex: 1;
    box-sizing: border-box;

    & > span {
      text-overflow: ellipsis;
      overflow: hidden;
      display: -webkit-box;
      -webkit-line-clamp: 1;
      -webkit-box-orient: vertical;
    }

    small {
      margin-left: 5px;
      font-size: 12px;
      color: rgba(255, 255, 255, 0.5);
    }

    /*hover菜单*/

    .list-menu {
      display: none;
      position: absolute;
      top: 50%;
      right: 10px;
      height: 40px;
      font-size: 0;
      transform: translateY(-50%);
    }
  }

  .list-artist,
  .list-album {
    display: block;
    width: 300px;
    .no-wrap();
    @media (max-width: 1440px) {
      width: 200px;
    }
    @media (max-width: 1200px) {
      width: 150px;
    }
  }

  .list-time {
    display: block;
    width: 60px;
    position: relative;

    .list-menu-icon-del {
      display: none;
      position: absolute;
      top: 50%;
      left: 0;
      transform: translateY(-50%);
    }
  }
}

@media (max-width: 960px) {
  .list-item .list-name {
    padding-right: 70px;
  }
}

@media (max-width: 768px) {
  .list-item {
    .list-name .list-menu {
      display: block;
    }

    .list-artist,
    .list-album {
      width: 20%;
    }
  }
}

@media (max-width: 640px) {
  .list-item {
    .list-artist {
      width: 80px;
    }

    .list-album,
    .list-time {
      display: none;
    }
  }
}
</style>
