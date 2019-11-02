<template lang="pug">
    .sidebar
        h1 Sidebar
        p {{ categories }}
</template>

<script>
import axios from "axios";

export default {
  name: "Sidebar",
  data() {
    return {
      categories: null
    };
  },
  mounted() {
    axios.get("/api/category").then(res => {
      this.categories = res.data;
    });
  }
};
</script>
