<template lang="pug">
    .sidebar
        h1 sidebar
        p {{ this.test }}
</template>
<script>
import axios from "axios";

export default {
  name: "Sidebar",
  data() {
    return {
      test: []
    };
  },
  mounted() {
    axios.get("/api/category").then(res => (this.test = res.data));
  }
};
</script>
