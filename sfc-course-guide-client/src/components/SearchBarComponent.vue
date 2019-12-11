<template>
  <input type="text" class="search-bar" v-model="input" @input="search" />
</template>

<style scoped>
input.search-bar {
  border: 0px solid #73ad21;
  border-radius: 25px;
  margin: 0px 16px 3em;
  padding: 20px;
  width: 50vw;
  font-size: 1em;
  box-shadow: 0px 0px 4px #c4c4c4;
  outline: none;
}
body.dark .search-bar {
  background-color: rgb(221, 221, 221);
}
</style>



<script>
export default {
  data() {
    return {
      input: '',
      query: '',
      searchResult: null,
      timer: null,
    };
  },
  methods: {
    search() {
      clearTimeout(this.timer);

      this.timer = setTimeout(() => {
        // Filter key like shift or fn
        if (this.query === this.input) return;

        this.query = this.input;

        new this.Request({
          method: 'Get',
          url: `${this.$setting.url.search}?query=${this.query}`,
        }).perform({
          onThen: (response) => {
            this.$emit('search', response.data);
          },
        });
      }, this.$setting.search.delay);
    },
  },
};
</script>
