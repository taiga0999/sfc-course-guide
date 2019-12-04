<template>
  <div class="course-box" @click="loadModal">
    <div class="course-info">
      <div class="subject-sort" v-html="course.Subject_sort"></div>
      <div class="title-memo-en" v-html="course.Title_memo_en"></div>
      <div class="language-en" v-html="course.Language_en"></div>
    </div>
    <div class="title-en" v-html="course.Title_en"></div>
    <div class="semester-en" v-html="course.Semester_en"></div>
    <div class="days-en" v-html="course.Days_en"></div>
    <!-- <div class="semester-days-en" v-html="course.Semester_days_en"></div> -->
    <div class="faculty-in-charge" v-html="course.Faculty_in_charge"></div>
    <course-modal v-if="showModal" @close="closeModal" :link="course.Link">
      <div slot="header" v-html="header"></div>
      <div slot="body" v-html="body"></div>
      <div slot="footer" v-html="footer"></div>
    </course-modal>
  </div>
</template>

<style scoped>
.course-box {
  border-radius: 25px;
  padding: 20px;
  width: 25em;
  margin: 1%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  line-height: 2em;
  box-shadow: 0px 2px 20px 0px #e5ddf5;
}

.course-box:hover {
  cursor: pointer;
  background: Highlight;
}

.course-box .highlight {
  color: red;
}

.course-info {
  display: inline-flex;
  justify-content: space-evenly;
}

.subject-sort {
  float: left;
}

.language-en {
  float: right;
}
</style>

<script>
export default {
  data() {
    return {
      showModal: false
    };
  },
  props: ["searchResult", "query"],
  watch: {
    // "searchResult.Highlight"() {
    //   let highlight = this.searchResult.Highlight;
    //   if (highlight !== null) {
    //     Object.entries(highlight).forEach(([k, v]) => {
    //       const regexp = new RegExp("<highlight>(.*?)</highlight>", "g");
    //       v = Array.from(
    //         new Set(
    //           [].concat(
    //             ...v.map(e =>
    //               e.match(regexp).map(match => match.replace(regexp, "$1"))
    //             )
    //           )
    //         )
    //       );
    //       console.log(v);
    //       v.forEach(e => {
    //         this.course[k] = this.course[k].replace(
    //           e,
    //           `<span class='highlight'>${e}</span>`
    //         );
    //         this.course[k] = v;
    //       });
    //     });
    //   }
    // }
  },
  computed: {
    course() {
      const course = this.searchResult.Course;

      // Text Processing
      const match = course.Title_memo_en.match(/(\(.*?\))/);
      course.Title_memo_en = (match && match[1]) || "";

      course.Language_en = course.Language_en || "TBD";
      course.Semester_en = course.Semester_en || "TBD";

      if (!course.Days_en) {
        course.Days_en = "TBD";
      } else if (typeof course.Days_en !== "string") {
        course.Days_en = course.Days_en.join(", ");
      }
      course.Semester_days_en = `${course.Semester_en}: ${course.Days_en}`;

      if (!course.Faculty_in_charge) {
        course.Faculty_in_charge = "TBD";
      } else if (typeof course.Faculty_in_charge === "string") {
        course.Faculty_in_charge = course.Faculty_in_charge;
      } else if (course.Faculty_in_charge.length > 3) {
        course.Faculty_in_charge = `${course.Faculty_in_charge.slice(0, 3).join(
          ", "
        )} ...`;
      } else {
        course.Faculty_in_charge = course.Faculty_in_charge.join(", ");
      }

      // highlight
      this.query.split(" ").forEach(query => {
        const regexp = new RegExp(query, "gi");
        Object.keys(course).forEach(key => {
          course[key] = course[key].replace(
            regexp,
            `<span class="highlight">$&</span>`
          );
        });
      });

      return course;
    },
    header() {
      return `<strong>${[
        this.course.Subject_sort,
        this.course.Title_en,
        `<span class="highlight">${this.course.Title_memo_en}</span>`
      ].join(" ")}</strong>`;
    },
    body() {
      return [
        this.course.Faculty_in_charge,
        this.course.Semester_en,
        this.course.Days_en,
        this.course.Language_en
      ].join("<br>");
    },
    footer() {
      return this.course.Description;
    }
  },
  methods: {
    loadModal() {
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
    }
  }
};
</script>
