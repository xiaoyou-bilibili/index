<template>
	<div class="header-wrapper">
		<img class="logo" :src="logoUrl" />
		<el-menu
			:default-active="activeIndex"
			class="top-menu-list"
			mode="horizontal"
			router
		>
			<el-menu-item
				index="File"
				:route="{ name: 'File', query: { fileType: 0, filePath: '/' } }"
				>网盘</el-menu-item
			>
		</el-menu>
	</div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
	name: 'Header',
	data() {
		return {
			logoUrl: require('_a/images/common/logo_header.png'),
			logoUrlXs: require('_a/images/common/logo_header_xs.png')
		}
	},
	computed: {
		...mapGetters(['isLogin', 'username']),
		// 当前激活菜单的 index
		activeIndex() {
			return this.$route.name || 'Home' //  获取当前路由名称
		},
		isProductEnv() {
			return (
				process.env.NODE_ENV !== 'development' &&
				location.origin === 'https://pan.qiwenshare.com'
			)
		},
		// 屏幕宽度
		screenWidth() {
			return this.$store.state.common.screenWidth
		}
	},
	methods: {
		// 奇文社区生产环境账户网址
		getAccountHref(path) {
			return `https://account.qiwenshare.com${path}?Rurl=${location.href}`
		},
		/**
		 * 退出登录
		 * @description 清除 cookie 存放的 token  并跳转到登录页面
		 */
		exitButton() {
			this.$message.success('退出登录成功！')
			this.$common.removeCookies(this.$config.tokenKeyName)
			this.$store.dispatch('getUserInfo').then(() => {
				this.$router.push({ name: 'Home' })
			})
		}
	}
}
</script>

<style lang="stylus" scoped>
@import '~_a/styles/varibles.styl';

.header-wrapper {
  width: 100%;
  padding: 0 20px;
  box-shadow: $tabBoxShadow;
  display: flex;

  .logo {
    margin: 14px 24px 0 24px;
    display: inline-block;
    height: 40px;
    cursor: pointer;
  }

  .logo-xs {
    display: none;
  }

  >>> .el-menu--horizontal {
    .el-menu-item:not(.is-disabled):hover {
      border-bottom-color: $Primary !important;
      background: $tabBackColor;
    }

    .external-link {
      padding: 0;
      a {
        display: block;
        padding: 0 20px;
      }
    }
  }

  .top-menu-list {
    flex: 1;

    .login, .register, .username, .exit, .user-exit-submenu {
      float: right;
    }
  }
}
</style>
