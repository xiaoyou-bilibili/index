<template>
	<div class="file-list-wrapper">
		<!-- 操作按钮 -->
		<el-header height="auto">
			<OperationMenu
				:file-parent="fileParent"
				@getSearchFileList="getSearchFileList"
				@getTableDataByType="getTableDataByType"
			></OperationMenu>
		</el-header>
<!--		<div class="middle-wrapper">-->
<!--			&lt;!&ndash; 面包屑导航栏 &ndash;&gt;-->
<!--			<BreadCrumb-->
<!--				class="breadcrumb"-->
<!--				:fileType="fileType"-->
<!--				:filePath="filePath"-->
<!--				@getTableDataByType="getTableDataByType"-->
<!--			></BreadCrumb>-->
<!--		</div>-->
		<!-- 文件列表-表格模式 -->
		<FileTable
			:file-parent="fileParent"
			:fileList="fileList"
			:loading.sync="loading"
			v-if="fileModel === 0"
			@getTableDataByType="getTableDataByType"
			@click.native.right="handleClickRight"
		></FileTable>
		<!-- 文件列表-网格模式 -->
		<FileGrid
			:file-parent="fileParent"
			:fileList="fileList"
			:loading="loading"
			v-if="fileModel === 1"
			@getTableDataByType="getTableDataByType"
			@click.native.right="handleClickRight"
		></FileGrid>
		<!-- 图片-时间线模式 -->
		<FileTimeLine
			class="image-model"
			:fileList="fileList"
			:loading.sync="loading"
			v-if="fileModel === 2"
			@getTableDataByType="getTableDataByType"
			@click.native.right="handleClickRight"
		></FileTimeLine>
		<div class="pagination-wrapper">
			<div class="current-page-count">当前页{{ fileList.length }}条</div>
			<!-- 回收站不展示分页组件 -->
			<el-pagination
				:current-page="pageData.currentPage"
				:page-size="pageData.pageCount"
				:total="pageData.total"
				:page-sizes="[10, 50, 100, 200]"
				:layout="
					screenWidth <= 768
						? 'total, prev, next, jumper'
						: 'sizes, total, prev, pager, next'
				"
				@current-change="handleCurrentChange"
				@size-change="handleSizeChange"
			>
			</el-pagination>
		</div>
	</div>
</template>

<script>
import OperationMenu from './components/OperationMenu.vue'
import FileTable from '_c/common/FileTable.vue'
import FileGrid from './components/FileGrid.vue'
import FileTimeLine from './components/FileTimeLine.vue'

import {
	getFileList,
	searchFile
} from '_r/file.js'

export default {
	name: 'FileList',
	components: {
		OperationMenu,
		FileTable,
		FileGrid,
		FileTimeLine
	},
	data() {
		return {
			fileNameSearch: '', // 搜索的文件夹名字
			loading: true, //  表格数据-loading
			fileList: [], //  表格数据-文件列表
			// 分页数据
			pageData: {
				currentPage: 1,
				pageCount: 50,
				total: 0
			}
		}
	},
	computed: {
		// 父文件夹夹
		fileParent() {
			return Number(this.$route.query.parent)
		},
		// 文件查看模式 0列表模式 1网格模式 2 时间线模式
		fileModel() {
			return this.$store.getters.fileModel
		},
		// 屏幕宽度
		screenWidth() {
			return this.$store.state.common.screenWidth
		}
	},
	watch: {
    fileParent() {
			// 当左侧菜单选择“全部”或“我的分享”，文件路径发生变化时，再重新获取文件列表
      this.setPageCount()
      this.getTableDataByType()
		},
		// 监听文件查看模式
		fileModel() {
			this.setPageCount()
		}
	},
	created() {
		this.setPageCount()
		this.getTableDataByType()
	},
	methods: {
		/**
		 * 文件展示区域的空白处右键事件
		 * @param {Document} event 右键事件对象
		 */
		handleClickRight(event) {
			event.preventDefault()
			// 只有在全部页面才可以进行以下操作
			if (this.fileParent >= 0) {
				this.$openBox
					.contextMenu({
						selectedFile: undefined,
						domEvent: event,
						serviceEl: this
					})
					.then((res) => {
						if (res === 'confirm') {
							this.getTableDataByType() //  刷新文件列表
						}
					})
			}
		},
		/**
		 * 表格数据获取相关事件 | 调整分页大小
		 */
		setPageCount() {
			this.pageData.currentPage = 1
			if (this.fileModel === 0) {
				this.pageData.pageCount = 50
			}
			if (this.fileModel === 1) {
				this.pageData.pageCount = 100
			}
		},
		/**
		 * 表格数据获取相关事件 | 获取文件列表数据
		 */
		getTableDataByType() {
			this.loading = true
      // 全部文件
      this.showFileList()
		},
		/**
     * 表格数据获取相关事件 | 获取当前路径下的文件列表
     */
    async showFileList() {
      // 判断一下是否为搜索模式
      let res = null
      if (this.fileNameSearch !== '') {
        res = await searchFile({
          current: this.pageData.currentPage,
          size: this.pageData.pageCount,
          search_keyword: this.fileNameSearch,
          search_type: 'name'
        })
      } else {
        res = await getFileList(this.fileParent || 891, {
          current: this.pageData.currentPage,
          size: this.pageData.pageCount
        })
      }
      if (res.code === 200) {
        this.fileList = res.data.list
        this.pageData.total = Number(res.data.total)
        this.loading = false
      } else {
        this.$message.error(res.msg)
      }
    },
		/**
		 * 表格数据获取相关事件 | 分页组件 | 当前页码改变
		 */
		handleCurrentChange(currentPage) {
			this.pageData.currentPage = currentPage
			this.getTableDataByType()
		},
		/**
		 * 表格数据获取相关事件 | 分页组件 | 页大小改变时
		 */
		handleSizeChange(pageCount) {
			this.pageData.pageCount = pageCount
			this.getTableDataByType()
		},
		/**
		 * 获取搜索文件结果列表
		 * @param {string} fileName 文件名称
		 */
		getSearchFileList(fileName) {
      // 每次搜索就重置current
      this.pageData.currentPage = 1
			this.loading = true
      this.fileNameSearch = fileName
      this.showFileList()
		}
	}
}
</script>

<style lang="stylus" scoped>
@import '~_a/styles/varibles.styl';

.file-list-wrapper {
  >>> .el-header {
    padding: 0;
  }

  .middle-wrapper {
    margin-bottom: 8px;
  }

  .pagination-wrapper {
    position: relative;
    border-top: 1px solid $BorderBase;
    height: 44px;
    line-height: 44px;
    text-align: center;

    .current-page-count {
      position: absolute;
      left: 16px;
      height: 32px;
      line-height: 32px;
      font-size: 13px;
      color: $RegularText;
    }
  }
}
</style>
