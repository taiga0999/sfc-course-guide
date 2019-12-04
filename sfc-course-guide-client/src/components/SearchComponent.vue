<template>
  <div>
    <logo></logo>
    <search-bar @search="foundCourseResults" ref="searchbar"></search-bar>
    <course-list :search-results="searchResults"></course-list>
  </div>
</template>

<style>
.dark-mode {
  background-color: black;
}
</style>

<script>
export default {
  data() {
    return {
      searchResults: null
    };
  },
  methods: {
    foundCourseResults(result) {
      this.searchResults = result;
    }
  },
  mounted() {
    if (!this.$setting.showCourseWhenPageLoad) return;

    new this.Request({
      method: "Get",
      url: this.$setting.url.search
    }).perform({
      onThen: response => {
        this.searchResults = response.data;
      },
      onFinally: () => {
        this.$refs.searchbar.$el.focus();
      }
    });
  }
};
</script>
