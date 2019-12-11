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
  box-shadow: 0px 2px 20px 0px #e5ddf5;
  outline: none;
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
