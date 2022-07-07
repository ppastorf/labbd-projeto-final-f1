<template>
  <div id="app">
    
    <transition name="fade" mode="out-in">
      <router-view />
    </transition>

    <Login />
    <Overview />
    <Relatorio />

    <!-- <Footer /> -->
  </div>
</template>

<script>
import Login from './views/Login.vue';
import Login from './views/Overview.vue';
import Login from './views/Relatorio.vue';
export default {
  computed: {
    notIsLoginPage() {
      return this.$route.name !== "login";
    }
  },
  components: {
    Login,
    Overview,
    Relatiorio
  }
};

</script>

<style>

</style>
