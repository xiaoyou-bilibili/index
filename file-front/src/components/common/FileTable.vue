<template>
	<div class="file-table-wrapper">
		<!-- 文件表格 -->
		<el-table
			class="file-table file-type-share"
			ref="multipleTable"
			fit
			v-loading="loading"
			element-loading-text="文件加载中……"
			tooltip-effect="dark"
			:data="fileList"
			:highlight-current-row="true"
			@selection-change="handleSelectRow"
			@sort-change="handleSortChange"
			@row-contextmenu="handleContextMenu"
		>
			<el-table-column
				type="selection"
				key="selection"
				width="56"
				align="center"
			></el-table-column>
			<el-table-column
				label
				prop="file_type"
				key="file_type"
				:width="screenWidth <= 768 ? 40 : 56"
				align="center"
				class-name="file-icon-column"
			>
				<template slot-scope="scope">
					<img
						:src="scope.row.preview"
						title="点击预览"
						style="width: 30px; max-height: 30px; cursor: pointer"
						@click="
							$file.handleFileNameClick(scope.row, scope.$index, sortedFileList)
						"
					/>
				</template>
			</el-table-column>
			<el-table-column
				prop="name"
				key="name"
				:sort-by="['name']"
				sortable
				show-overflow-tooltip
			>
				<template slot="header">
					<span>文件名</span>
				</template>
				<template slot-scope="scope">
					<span
						@click="$file.handleFileNameClick(scope.row, scope.$index, sortedFileList)">
						<span
							class="file-name"
							style="cursor: pointer"
							v-html="scope.row.name"
						></span>
					</span>
				</template>
			</el-table-column>
			<el-table-column
				label="大小"
				width="100"
				prop="fileSize"
				key="fileSize"
				:sort-by="['isDir', 'fileSize']"
				sortable
				align="right"
				v-if="selectedColumnList.includes('fileSize') && screenWidth > 768"
			>
				<template slot-scope="scope">
          {{ scope.row.file_type === 1 ? $file.calculateFileSize(scope.row.info.file_size) : '' }}
				</template>
			</el-table-column>
			<el-table-column
				label="修改日期"
				prop="createTime"
				key="createTime"
				width="160"
				sortable
				align="center"
			>
        <template slot-scope="scope">
          {{ scope.row.create_time }}
        </template>
      </el-table-column>
			<el-table-column
				label=""
				key="operation"
				width="48"
				v-if="screenWidth <= 768"
			>
				<template slot-scope="scope">
					<i
						class="file-operate el-icon-more"
						:class="`operate-more-${scope.$index}`"
						@click="handleClickMore(scope.row, $event)"
					></i>
				</template>
			</el-table-column>
		</el-table>
	</div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
	name: 'FileTable',
	props: {
		// 父文件夹
		fileParent: {
			required: true,
			type: Number
		},
		// 文件列表
		fileList: {
			required: true,
			type: Array
		},
		// 文件加载状态
		loading: {
			required: true,
			type: Boolean
		}
	},
	data() {
		return {
			officeFileType: ['ppt', 'pptx', 'doc', 'docx', 'xls', 'xlsx'],
			sortedFileList: [] //  排序后的表格数据
		}
	},
	computed: {
		//  selectedColumnList: 判断当前用户设置的左侧栏是否折叠
		...mapGetters(['selectedColumnList']),
		// 路由名称
		routeName() {
			return this.$route.name
		},
		// 屏幕宽度
		screenWidth() {
			return this.$store.state.common.screenWidth
		}
	},
	watch: {
		/**
		 * 文件路径变化时清空表格已选行
		 */
    fileParent() {
			this.clearSelectedTable()
			this.$refs.multipleTable.clearSort()
		},
		/**
		 * 文件列表变化时清空表格已选行
		 */
		fileList() {
			this.clearSelectedTable()
			this.$refs.multipleTable.clearSort()
			this.sortedFileList = this.fileList
		}
	},
	methods: {
		/**
		 * 当表格的排序条件发生变化的时候会触发该事件
		 */
		handleSortChange() {
			this.sortedFileList = this.$refs.multipleTable.tableData
		},
		/**
		 * 表格某一行右键事件
		 * @description 打开右键菜单
		 * @param {object} row 当前行数据
		 * @param {object} column 当前列数据
		 * @param {object} event 当前右键元素
		 */
		handleContextMenu(row, column, event) {
			// 阻止右键事件冒泡
			event.cancelBubble = true
			// xs 以上的屏幕
			if (this.screenWidth > 768) {
				event.preventDefault()
				this.$refs.multipleTable.setCurrentRow(row) //  选中当前行
				this.$openBox
					.contextMenu({
						selectedFile: row,
						domEvent: event
					})
					.then((res) => {
						this.$refs.multipleTable.setCurrentRow() //  取消当前选中行
						if (res === 'confirm') {
							this.$emit('getTableDataByType') //  刷新文件列表
						}
					})
			}
		},
		/**
		 * 清空表格已选行
		 * @description 用于父组件调用 | 本组件调用，请勿删除
		 */
		clearSelectedTable() {
			this.$refs.multipleTable.clearSelection()
		},
		/**
		 * 表格选择项发生变化时的回调函数
		 * @param {[]} selection 勾选的行数据
		 */
		handleSelectRow(selection) {
			this.$store.commit('changeSelectedFiles', selection)
			this.$store.commit('changeIsBatchOperation', selection.length !== 0)
		},
		/**
		 * 更多图标点击事件
		 * @description 打开右键菜单
		 * @param {object} row 当前行数据
		 * @param {object} event 当前右键元素
		 */
		handleClickMore(row, event) {
			this.$refs.multipleTable.setCurrentRow(row) //  选中当前行
			this.$openBox
				.contextMenu({
					selectedFile: row,
					domEvent: event
				})
				.then((res) => {
					this.$refs.multipleTable.setCurrentRow() //  取消当前选中行
					if (res === 'confirm') {
						this.$emit('getTableDataByType') //  刷新文件列表
					}
				})
		}
	}
}
</script>

<style lang="stylus" scoped>
@import '~_a/styles/varibles.styl';
@import '~_a/styles/mixins.styl';

.file-table-wrapper {
  margin-top: 2px;

  .file-type-0 {
    height: calc(100vh - 206px) !important;

    >>> .el-table__body-wrapper {
      height: calc(100vh - 262px) !important;
    }
  }

  .file-type-6 {
    height: calc(100vh - 211px) !important;

    >>> .el-table__body-wrapper {
      height: calc(100vh - 263px) !important;
    }
  }

  .file-table.share {
    height: calc(100vh - 109px) !important;

    >>> .el-table__body-wrapper {
      height: calc(100vh - 161px) !important;
    }
  }

  .file-table {
    width: 100% !important;
    height: calc(100vh - 203px);

    >>> .el-table__header-wrapper {
      th {
        // background: $tabBackColor;
        padding: 4px 0;
        color: $RegularText;
      }

      .el-icon-circle-plus, .el-icon-remove {
        margin-left: 6px;
        cursor: pointer;
        font-size: 16px;

        &:hover {
          color: $Primary;
        }
      }
    }

    >>> .el-table__body-wrapper {
      height: calc(100vh - 255px);
      overflow-y: auto;
      setScrollbar(6px, transparent, #C0C4CC);

      td {
        padding: 8px 0;
        .file-name {
          .keyword {
            color: $Danger;
          }
        }
      }

      .el-icon-warning {
        font-size: 16px;
        color: $Warning;
      }

      .el-icon-time {
        font-size: 16px;
        color: $Success;
      }
    }
  }
}
.right-menu-list {
  position: fixed;
  display: flex;
  flex-direction: column;
  background: #fff;
  border: 1px solid $BorderLighter;
  border-radius: 4px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  z-index: 2;
  padding: 4px 0;
  color: $RegularText;

  .right-menu-item,
  .unzip-item {
    padding: 0 16px;
    height: 36px;
    line-height: 36px;
    cursor: pointer;
    &:hover {
      background: $PrimaryHover;
      color: $Primary;
    }
    i {
      margin-right: 8px;
    }
  }

  .unzip-menu-item {
    position: relative;
    &:hover {
      .unzip-list {
        display: block;
      }
    }
    .unzip-list {
      position: absolute;
      display: none;
      .unzip-item {
        width: 200px;
        setEllipsis(1)
      }
    }
  }
}
.right-menu-list,
.unzip-list {
  background: #fff;
  border: 1px solid $BorderLighter;
  border-radius: 4px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  z-index: 2;
  padding: 4px 0;
  color: $RegularText;
}
</style>
