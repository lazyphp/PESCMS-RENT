<template>
  <component :is="autoLayout"></component>
</template>

<script>
import layout from './views/layout.vue';
import login from './views/login.vue';
import logo from './views/logo.vue';
import findpw from './views/findpw.vue';
import install from './views/Install.vue';
import notfound from './views/404.vue';

export default {

  computed: {
    autoLayout() {
      const route = this.$route;
      if (route.meta.layout === 'layout') {
        return layout;
      } else if (route.meta.layout === 'login') {
        return login;
      } else if (route.meta.layout === 'install') {
        return install;
      } else if (route.meta.layout === 'findpw') {
        return findpw;
      } else if (route.meta.layout === 'notfound') {
        return notfound;
      }
      return logo;
    }
  },
  mounted() {
    if (localStorage.getItem('install') == 1) {
      return;
    }

    this.$axios.get('/home/setting/install').then(response => {
      if (response.data.code == 0) {
        if (response.data.data.install == 0) {
          this.$router.push({ name: 'install' });
        } else {
          localStorage.setItem('install', 1);
        }
      }
    })
  }
};  
</script>
