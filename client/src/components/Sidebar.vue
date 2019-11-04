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
      categories: Array
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
  background-color: #ffffff;
  border-radius: 4px;
  border: 1px solid #f5f5f5;
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.1);
  color: #2c3e50;

  a {
    text-decoration: none;

    &:visited {
      color: #2c3e50;
    }

    &:link {
      color: #2c3e50;
    }
  }
}

p {
  margin: 0;
}

.mainCategory {
  padding: 10px 0px;
}

.subCategory {
  margin-left: 25px;
  border-left: 1px solid black;
}

.category {
  padding: 5px 5px 5px 15px;

  &:hover {
    background-color: lightblue;
  }
}
</style>
