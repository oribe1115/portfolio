<template lang="pug">
    .main-category
        h1 Main Category
        ContentsList(:path="path")
</template>

<script>
import ContentsList from "@/components/ContentsList.vue";

export default {
  name: "MainCategory",
  components: {
    ContentsList
  },
  data() {
    return {
      path: ""
    };
  },
  mounted() {
    this.path = "/api/category/content/" + this.$route.params.mainID;
  }
};
</script>
