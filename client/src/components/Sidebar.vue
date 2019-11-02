<template lang="pug">
    .sidebar
        .mainCategory(v-for="mainCategory in categories" :key="mainCategory.id")
          router-link(:to="{ name: 'mainCategory', params: { mainID: mainCategory.id }}")
            .category
              p {{ mainCategory.name }}
          .subCategory(v-for="subCategory in mainCategory.sub_categories" :key="subCategory.id")
            router-link(:to="{ name: 'subCategory', params: { subID: subCategory.id }}")
              .category
                p {{ subCategory.name }}
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

<style lang="scss">
.sidebar {
  font-size: 25px;
  text-align: left;
  padding: 10px;
}

.mainCategory {
  padding: 10px;
  border: 1px solid black;
}

.subCategory {
  padding-left: 10px;
}
</style>
