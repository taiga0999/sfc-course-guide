<template>
  <div :class="{'dark': dark}" class="index">
    <button @click="themeSwitcher">{{ this.switcherName }}</button>
    <router-view></router-view>
  </div>
</template>

<style>
/* global styles */
body {
  margin: 0;
}

body.dark {
  background-color: rgb(31, 31, 31);
}

.aligner {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
}
button {
  margin: 20px 0% 0% 70%;
}
</style>

<style scoped>
/* local styles */
</style>

<script>
export default {
  data() {
    return {
      switcherName: "dark",
      dark: true
    };
  },
  methods: {
    themeSwitcher() {
      // TODO(1) manually trigger darkmode for body element
      document.body.classList[this.dark ? "add" : "remove"]("dark");
      this.switcherName = this.dark ? "light" : "dark";
      this.dark = !this.dark;
    }
  },
  mounted() {
    // TODO(1)
    this.dark && this.themeSwitcher();
  }
};
</script>
