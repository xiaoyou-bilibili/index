<template>
	<!-- 删除文件对话框 -->
	<el-dialog
		title="删除文件"
		:visible.sync="visible"
		:close-on-click-modal="false"
		width="550px"
		@close="handleDialogClose"
	>
		<div >此操作将永久删除该文件, 是否继续？</div>
		<div slot="footer" class="dialog-footer">
			<el-button @click="handleDialogClose">取 消</el-button>
			<el-button
				type="primary"
				:loading="sureBtnLoading"
				@click="handleDialogSure"
				>确 定</el-button
			>
		</div>
	</el-dialog>
</template>

<script>
import {
	batchDeleteFile
} from '_r/file.js'

export default {
	name: 'DeleteFileDialog',
	data() {
		return {
			visible: false, //  对话框是否可见
			sureBtnLoading: false
		}
	},
	methods: {
		/**
		 * 删除文件对话框 | 对话框关闭的回调
		 * @description 关闭对话框，重置表单
		 */
		handleDialogClose() {
			this.visible = false
			this.callback('cancel')
		},
		/**
		 * 删除文件对话框 | 确定按钮点击事件
		 * @description 区分 删除到回收站中 | 在回收站中彻底删除，调用相应的删除文件接口
		 */
		async handleDialogSure() {
			this.sureBtnLoading = true
			let res = null
			// 批量删除模式
			if (this.isBatchOperation) {
        let ids = []
        this.fileInfo.forEach((item)=>{ids.push(item.id)})
        res = await batchDeleteFile(ids.join(","))
			} else {
				// 单文件删除模式
        res = await batchDeleteFile(this.fileInfo.id)
			}
			if (res.code === 200) {
				this.sureBtnLoading = false
				this.$message.success('删除成功')
				this.visible = false
				this.callback('confirm')
			} else {
				this.sureBtnLoading = false
				this.$message.error('删除失败，' + res.msg)
			}
		}
	}
}
</script>
