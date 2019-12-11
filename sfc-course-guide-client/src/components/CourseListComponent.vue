<template>
  <div class="course-list">
    <div class="menu">
      <h2 class="result-stats">{{ resultStats }}</h2>
      <select id="showing-limit" v-model="pagination.perPage">
        <option v-for="(option, index) in pagination.options" :key="index">{{
          option
        }}</option>
      </select>
    </div>
    <div class="course-wrapper" v-if="pagination.total">
      <course
        v-for="(result, index) in resultsShowing"
        :key="index"
        :search-result="result"
        :query="query"
      ></course>
    </div>
  </div>
</template>

<style scoped>
.course-list {
  height: 50vh;
}
.menu {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-around;
}
.result-stats {
  margin-top: 1%;
  color: rgb(127, 127, 127);
  /* text-shadow: 0px 0px 4px #c4c4c4; */
}
#showing-limit {
  margin-top: 1.5em;
}
.course-wrapper {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: center;
}
</style>

<!--<style scoped class="dark">
body.dark .result-stats {
  color: white;
}
</style> -->

<script>
export default {
  data() {
    return {
      pagination: {
        options: [10, 20, 50, 100],
        total: 0,
        count: 0,
        perPage: 10,
      },
    };
  },
  props: ['searchResults'],
  watch: {
    searchResults() {
      this.pagination.total = this.searchResults
        ? this.searchResults.Stat.Total
        : 0;
    },
    count() {
      this.pagination.count = this.count;
    },
  },
  computed: {
    count() {
      return Math.min(this.pagination.perPage, this.pagination.total);
    },
    resultStats() {
      let searchMsg = '';
      if (this.searchResults && this.pagination.count > 0) {
        searchMsg = `(${
          this.searchResults.Stat.Latency > 0
            ? this.searchResults.Stat.Latency / 1000
            : 'less then 0.001'
        } seconds)`;
      }

      return `${this.pagination.total} results ${searchMsg}`;
    },
    query() {
      return this.searchResults.Query || '';
    },
    resultsShowing() {
      return this.searchResults.Hits.slice(0, this.pagination.perPage);
    },
  },
};
</script>
