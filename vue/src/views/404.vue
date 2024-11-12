<template>
  <el-result title="404" sub-title="页面丢失拉~">
    <template #icon>
      <el-image src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png" />
    </template>
    <template #extra>
      <el-button type="primary" @click="goHome">返回房租管理</el-button>
    </template>
  </el-result>
</template>
<script>
export default {
  methods: {
    goHome() {
      this.$router.push('/house/list');
    }
  }
}

</script>