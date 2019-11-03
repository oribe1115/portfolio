<template lang="pug">
    .main-category
      .path
        router-link(:to="{ name: 'mainCategory', params: { mainID: contentsList.main_category.id } }")
          .main
            h1 {{ contentsList.main_category.name }}
      .description
        p {{ contentsList.main_category.description }}
      ContentsList(:contents="contentsList.contents")
</template>

<script>
import ContentsList from "@/components/ContentsList.vue";
import axios from "axios";

export default {
  name: "MainCategory",
  components: {
    ContentsList
  },
  data() {
    return {
      contentsList: null
    };
  },
  mounted() {
    axios
      .get("/api/category/content/" + this.$route.params.mainID)
      .then(res => {
        this.contentsList = res.data;
      });
  }
};
</script>

<style lang="scss">
.main-category {
  padding: 0px 30px;
}
</style>
