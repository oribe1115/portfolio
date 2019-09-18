<template lang="pug">
    .sidebar
        h1 sidebar
        .main-categories
            .main-item(v-for="mainCategory in mainCategories" :key = "mainCategory.id")
                p {{ mainCategory.name }}
                .sub-categoies(v-for="subCategory in mainCategory.sub_categories" :key "subCategory.id")
                    p {{ subCategory.name }}
</template>
<script>
import axios from "axios";

export default {
  name: "Sidebar",
  data() {
    return {
      mainCategories: []
    };
  },
  mounted() {
    axios.get("/api/category").then(res => (this.mainCategories = res.data));
  }
};
</script>
