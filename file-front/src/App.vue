<template>
	<div id="app">
		<router-view class="main-content"></router-view>
		<el-backtop class="backtop" target="#app" title="回到顶部"></el-backtop>
	</div>
</template>

<script>
export default {
	name: 'App',
	mounted() {
		const that = this
		window.addEventListener('resize', function () {
			return (() => {
				that.$store.commit('changeScreenWidth', document.body.clientWidth)
			})()
		})
	}
}
</script>
<style lang="stylus" scoped>
@import '~_a/styles/varibles.styl';

#app {
  height: 100%;
  overflow-x: hidden;
  -webkit-text-size-adjust: none;
  overflow-y: auto;
  >>> .el-backtop {
    background-color: $Success;
    color: #fff;
    z-index: 3;
  }
  .main-content {
    flex: 1;
    width: 90%;
    min-height: calc(100vh);
    margin: 0 auto;
    display: flex;
  }
}
</style>
