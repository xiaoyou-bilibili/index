<template>
  <!--搜索-->
  <div class="search">
    <mm-loading v-model="mmLoadShow" />
    <div class="search-head">
      <!--      <span-->
      <!--        v-for="(item, index) in Artists"-->
      <!--        :key="index"-->
      <!--        @click="clickHot(item.first)"-->
      <!--      >-->
      <!--        {{ item.first }}-->
      <!--      </span>-->
      <input
        v-model.trim="searchValue"
        class="search-input"
        type="text"
        placeholder="音乐/歌手"
        @keyup.enter="onEnter"
      />
    </div>
    <music-list
      ref="musicList"
      :list="list"
      list-type="pullup"
      @select="selectItem"
    />
  </div>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from 'vuex'
import { search } from 'api'
import MmLoading from 'base/mm-loading/mm-loading'
import MusicList from 'components/music-list/music-list'
import { loadMixin } from '@/utils/mixin'

export default {
  name: 'Search',
  components: {
    MmLoading,
    MusicList
  },
  mixins: [loadMixin],
  data() {
    return {
      searchValue: '', // 搜索关键词
      Artists: [], // 热搜数组
      list: [], // 搜索数组
      page: 0, // 分页
      lockUp: true, // 是否锁定上拉加载事件,默认锁定
      mmLoadShow: false
    }
  },
  computed: {
    ...mapGetters(['playing', 'currentMusic'])
  },
  methods: {
    listAppend(data) {
      if (data !== null && data.length > 0) {
        this.list = [...this.list, ...data]
      }
    },
    // 搜索事件
    async onEnter() {
      if (this.searchValue.replace(/(^\s+)|(\s+$)/g, '') === '') {
        this.$mmToast('搜索内容不能为空！')
        return
      }
      this.mmLoadShow = true
      // 先把数据清空
      this.list = []
      await search('name', this.searchValue, 1).then(({ data }) => this.listAppend(data.list))
      await search('singer', this.searchValue, 1).then(({ data }) => this.listAppend(data.list))
      this._hideLoad()
    },
    // 播放歌曲
    async selectItem(music) {
      try {
        this.selectAddPlay(music)
      } catch (error) {
        this.$mmToast('哎呀，出错啦~')
      }
    },
    ...mapMutations({
      setPlaying: 'SET_PLAYING'
    }),
    ...mapActions(['selectAddPlay'])
  }
}
</script>

<style lang="less" scoped>
.search {
  position: relative;
  width: 100%;
  height: 100%;
  .search-head {
    display: flex;
    height: 40px;
    padding: 10px 15px;
    overflow: hidden;
    background: @search_bg_coloe;
    span {
      line-height: 40px;
      margin-right: 15px;
      cursor: pointer;
      &:hover {
        color: @text_color_active;
      }
      @media (max-width: 640px) {
        & {
          display: none;
        }
      }
    }
    .search-input {
      flex: 1;
      height: 40px;
      box-sizing: border-box;
      padding: 0 15px;
      border: 1px solid @btn_color;
      outline: 0;
      background: transparent;
      color: @text_color_active;
      font-size: @font_size_medium;
      box-shadow: 0 0 1px 0 #fff inset;
      &::placeholder {
        color: @text_color;
      }
    }
  }
  .musicList {
    width: 100%;
    height: calc(~'100% - 50px');
  }
}
</style>
