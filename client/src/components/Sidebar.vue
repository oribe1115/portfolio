<template lang="pug">
    .sidebar
        h1 sidebar
        p {{ fromStore }}
        .main-category(v-for="mainCategory in mainCategories" :key = "mainCategory.id")
            .main-item(@click="clickMainCategory(mainCategory.id)")
                p {{ mainCategory.name }}
            .sub-category(v-for="subCategory in mainCategory.sub_categories" :key "subCategory.id")
                .sub-item(@click="clickSubCategory(subCategory.id)")
                    p {{ subCategory.name }}
</template>
<script>
import axios from "axios";
import store from "../store.js";

export default {
  name: "Sidebar",
  data() {
    return {
      mainCategories: [],
      fromStore: ""
    };
  },
  methods: {
    clickMainCategory(id) {
      store.setToMainCategory(id);
    },
    clickSubCategory(id) {
      store.setToSubCategory(id);
    }
  },
  mounted() {
    axios.get("/api/category").then(res => (this.mainCategories = res.data));
    store.setToMainCategory("id");
    this.fromStore = store.state;
  }
};
</script>
