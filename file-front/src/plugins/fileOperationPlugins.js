/**
 * 以服务的方式，将对文件的一些操作挂载到 Vue 上
 * @description 各个服务的参数传递查看服务封装对应目录下的 index.js 文件，里面注明了需要传递的参数
 * @author 小鲤鱼听听
 */

/**
 * 动态引入操作文件的弹窗组件
 */
const fileOperateDialog = require.context(
	'_c/file/dialog',
	true,
	/\.\/(.*)\/index.js$/
)
/**
 * 服务挂载到 Vue - $openDialog，各个服务的名称为其对应的文件夹的名称
 * @description 例如，创建文件夹功能：服务封装路径 '_c/file/dialog/addFolder' ，则如下：
 * 在 *.vue 文件中，使用 this.$openDialog.addFolder 调用创建文件夹服务
 * 在 *.js 文件中，需要先在文件顶部引入 Vue ，即 import Vue from 'vue' ，然后使用 Vue.property.$openDialog.addFolder 调用创建文件夹服务
 */
const openDialog = fileOperateDialog
	.keys()
	.map((key) => {
		return { [key.split('/')[1]]: fileOperateDialog(key).default }
	})
	.reduce((pre, next) => {
		return { ...pre, ...next }
	}, {})

/**
 * 动态添加操作文件的遮罩或浮层组件
 */
const fileOperateBox = require.context(
	'_c/file/box',
	true,
	/\.\/(.*)\/index.js$/
)
/**
 * 服务挂载到 Vue - $openBox，各个服务的名称为其对应的文件夹的名称
 * @description 例如，图片预览功能：服务封装路径 '_c/file/box/imgPreview' ，则如下：
 * 在 *.vue 文件中，使用 this.$openDialog.imgPreview 调用图片预览服务
 * 在 *.js 文件中，需要先在文件顶部引入 Vue ，即 import Vue from 'vue' ，然后使用 Vue.property.$openDialog.imgPreview 调用图片预览服务
 */
const openBox = fileOperateBox
	.keys()
	.map((key) => {
		return { [key.split('/')[1]]: fileOperateBox(key).default }
	})
	.reduce((pre, next) => {
		return { ...pre, ...next }
	}, {})

export default { openDialog, openBox }
