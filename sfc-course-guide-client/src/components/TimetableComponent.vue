<template>
  <div class="timetable">
    <table id="test" border="1">
      <thead>
        <tr>
          <th style="width: 100px;"></th>
          <th style="width: 200px;">月</th>
          <th style="width: 200px;">火</th>
          <th style="width: 200px;">水</th>
          <th style="width: 200px;">木</th>
          <th style="width: 200px;">金</th>
          <th style="width: 200px;">土</th>
        </tr>
      </thead>
      <tbody>
        <!-- <tr v-for="n in 6">
          <th>{{ n }}</th>
          <td v-for="i in 6">{{ empty }}</td>
        </tr>-->
        <tr>
          <th>１</th>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
        </tr>
        <tr>
          <th>２</th>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
        </tr>
        <tr>
          <th>３</th>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
        </tr>
        <tr>
          <th>４</th>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
        </tr>
        <tr>
          <th>５</th>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
        </tr>
        <tr>
          <th>６</th>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
        </tr>
      </tbody>
    </table>
    <button @click="reset_timetable">Reset</button>
  </div>
</template>

<style scoped>
table {
  color: rgb(75, 75, 75);
  border: 0.5px solid #aaa;
  border-collapse: separate;
  border-spacing: 0;
  border-radius: 6px;
  overflow: hidden;
  background-color: rgb(255, 255, 255);
  width: 800px;
  margin-left: auto;
  margin-right: auto;
}
table thead th,
table tbody th,
table tbody td {
  /* padding: 0.6em 3em; */
  border-bottom: 0.5px solid #aaa;
}
table thead th {
  background-color: rgb(79, 192, 141);
}
table tbody th {
  background-color: #eee;
}
table thead th + th,
table tbody td {
  border-left: 0.5px solid #aaa;
}
table tbody tr:last-child th,
table tbody tr:last-child td {
  border-bottom: none;
}

body.dark table {
  background-color: #282c34;
  color: rgb(221, 221, 221);
}

body.dark table thead th {
  color: #282c34;
  background-color: rgb(192, 147, 79);
}

body.dark table tbody th {
  color: #282c34;
}
</style>

<script>
export default {
  data: function() {
    return {
      days: ["月", "火", "水", "木", "金", "土"],
      empty: "&nbsp;"
    };
  },
  mounted() {
    if (localStorage.local_array) {
      var table = document.getElementById("test");
      var array = JSON.parse(localStorage.local_array);
      for (var i = 0; i <= array.length; i++) {
        var titles_en = array[i]["title"];
        var x = array[i]["x"];
        var y = array[i]["y"];
        var td = table.rows[y].cells[x];
        td.firstChild.nodeValue = titles_en;
      }
    }
  },
  methods: {
    reset_timetable() {
      localStorage.clear();
      location.reload();
    }
  }
};
</script>