<template lang="pug">
  .sub-category
    .path
      router-link(:to="{ name: 'mainCategory', params: { mainID: contentsList.main_category.id } }")
        .main
          h1 {{ contentsList.main_category.name }}
      h1 >
      router-link(:to="{ name: 'subCategory', params: { subID: contentsList.sub_category.id } }")
        .sub
          h1 {{ contentsList.sub_category.name }}
    .description
      p {{ contentsList.sub_category.description }}
    ContentsList(:contents="contentsList.contents")
</template>

<script>
import ContentsList from "@/components/ContentsList.vue";
import axios from "axios";

export default {
  name: "SubCategory",
  components: {
    ContentsList
  },
  data() {
    return {
      contentsList: null
    };
  },
  created() {
    this.fetchData();
  },
  watch: {
    $route: "fetchData"
  },
  methods: {
    fetchData: function() {
      axios
        .get("/api/category/content/sub/" + this.$route.params.subID)
        .then(res => {
          this.contentsList = res.data;
        });
    }
  }
};
</script>

<style lang="scss">
.sub-category {
  padding: 30px;
}

.path {
  display: flex;
}

h1 {
  margin: 5px;
}
</style>
