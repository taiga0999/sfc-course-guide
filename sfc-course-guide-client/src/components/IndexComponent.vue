<template>
  <div>
    <button @click="darkMode" id="dark-mode">Dark</button>
    <button @click="lightMode" id="light-mode">Light</button>
    <router-view></router-view>
  </div>
</template>

<style>
/* global styles */
body {
  margin: 0;
}

.aligner {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
}
</style>

<style scoped>
/* local styles */
</style>

<script>
export default {
  methods: {
    darkMode() {
      document.body.classList.add("dark-mode");
    },
    lightMode() {
      document.body.classList.remove("dark-mode");
    }
  }
};
</script>
