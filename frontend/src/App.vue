<template>
  <div id="app">
    

    <transition name="fade" mode="out-in">
      <router-view />
    </transition>

    <!-- <Footer /> -->
  </div>
</template>

<script>
export default {
  computed: {
    notIsLoginPage() {
      return this.$route.name !== "login";
    }
  }
};
</script>

<style>

</style>
