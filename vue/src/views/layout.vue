<template>
  <el-container direction="vertical">
    <v-header></v-header>
  
    <el-main><router-view></router-view></el-main>
  </el-container>
</template>

<script>
import header from './header.vue'

export default {
  components: {
    'v-header': header,
  },
}

</script>