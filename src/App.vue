<template>
  <el-container>
    <el-main>
      <router-view />
    </el-main>
    <el-footer>Powered by send.ykmeiz.dev</el-footer>
  </el-container>
</template>

<script>

export default {
  name: 'App',
};
</script>

<style>
#app {
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB',
  'Microsoft YaHei', '微软雅黑', Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.el-main{
  margin: auto;
}
.el-footer {
  margin: 10em 0 5em 0;
  color: #909399;
}

</style>
