<template>
  <div>
    <h1>Hello Vue Router</h1>
    <p>
      <router-link to="/foo">Go to Foo</router-link>
      <router-link to="/bar">Go to Bar</router-link>
    </p>
    <router-view></router-view>
  </div>
</template>

<style scoped>
</style>

<script>
export default {
  mounted() {
    console.log('Component mounted.');
  },
};
</script>
