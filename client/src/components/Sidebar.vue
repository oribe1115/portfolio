<template lang="pug">
    .sidebar
        .mainCategory(v-for="mainCategory in categories" :key="mainCategory.id")
            | {{ mainCategory.name }}
            .category( @click="showMainCategory(mainCategory)")
                p {{ mainCategory.name }}
            .subCategory(v-for="subCategory in mainCategory.sub_categories" :key="subCategory.id" @click="showSubCategory(subCategory)")
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
  methods: {
    showMainCategory(mainCategory) {
      this.$parent.category = "main";
      this.$parent.categoryID = mainCategory.id;
      this.$parent.addPath = "";
    },
    showSubCategory(subCategory) {
      this.$parent.category = "sub";
      this.$parent.categoryID = subCategory.id;
      this.$parent.addPath = "sub/";
    }
  },
  mounted() {
    axios.get("/api/category").then(res => {
      this.categories = res.data;
    });
  }
};
</script>

<style lang="scss">
.sidebar {
  font-size: 25px;
  text-align: left;
  padding: 10px;
  width: 20%;
}

.mainCategory {
  padding: 10px;
  border: 1px solid black;
}

.subCategory {
  padding-left: 10px;
}
</style>
