<template lang="pug">
    .main-category
      .path
        router-link(:to="{ name: 'mainCategory', params: { mainID: main.id } }")
          .main
            h1 {{ main.name }}
      .description
        p {{ main.description }}
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
      contentsList: Array,
      main: Object
    };
  },
  mounted() {
    axios
      .get("/api/category/content/" + this.$route.params.mainID)
      .then(res => {
        this.contentsList = res.data;
        this.main = this.contentsList.main_category;
      });
  }
};
</script>

<style lang="scss">
.main-category {
  padding: 0px 30px;
}
</style>
