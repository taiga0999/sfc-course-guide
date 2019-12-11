<template>
  <div :class="{ dark: dark }" class="index">
    <span class="themeSwitcher" @click="themeSwitcher">
      <svg
        v-if="this.dark"
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="feather feather-sunrise"
      >
        <path d="M17 18a5 5 0 0 0-10 0" />
        <line x1="12" y1="2" x2="12" y2="9" />
        <line x1="4.22" y1="10.22" x2="5.64" y2="11.64" />
        <line x1="1" y1="18" x2="3" y2="18" />
        <line x1="21" y1="18" x2="23" y2="18" />
        <line x1="18.36" y1="11.64" x2="19.78" y2="10.22" />
        <line x1="23" y1="22" x2="1" y2="22" />
        <polyline points="8 6 12 2 16 6" />
      </svg>
      <svg
        v-else
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="feather feather-sunset"
      >
        <path d="M17 18a5 5 0 0 0-10 0" />
        <line x1="12" y1="9" x2="12" y2="2" />
        <line x1="4.22" y1="10.22" x2="5.64" y2="11.64" />
        <line x1="1" y1="18" x2="3" y2="18" />
        <line x1="21" y1="18" x2="23" y2="18" />
        <line x1="18.36" y1="11.64" x2="19.78" y2="10.22" />
        <line x1="23" y1="22" x2="1" y2="22" />
        <polyline points="16 5 12 9 8 5" />
      </svg>
    </span>
    <router-view></router-view>
  </div>
</template>

<style>
/* global styles */
body {
  margin: 0;
  font-family: Helvetica, Arial, sans-serif;
}

body.dark {
  background-color: #282c34;
  color: rgb(221, 221, 221);
}

.aligner {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
}
.themeSwitcher {
  position: absolute;
  color: #ff6200;
  top: 1.5em;
  right: 1.5em;
  cursor: pointer;
}
svg {
  width: 2em;
  height: 2em;
}
</style>

<style scoped>
/* local styles */
</style>

<script>
export default {
  data() {
    return {
      switcherName: 'dark',
      dark: true,
    };
  },
  methods: {
    themeSwitcher() {
      // TODO(1) manually trigger darkmode for body element
      document.body.classList[this.dark ? 'add' : 'remove']('dark');
      this.switcherName = this.dark ? 'light' : 'dark';
      this.dark = !this.dark;
    },
  },
  mounted() {
    // TODO(1)
    if (this.dark) this.themeSwitcher();
  },
};
</script>
