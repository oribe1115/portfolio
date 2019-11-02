<template lang="pug">
    .sidebar
        .mainCategory(v-for="mainCategory in categories" :key="mainCategory.id")
            | {{ mainCategory.name }}
            .subCategory(v-for="subCategory in mainCategory.sub_categories" :key="subCategory.id")
                | {{ subCategory.name }}
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
