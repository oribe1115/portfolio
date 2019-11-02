<template lang="pug">
    .main-category
        h1 Main Category
        ContentsList(:contents="contents")
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
      contents: null
    };
  },
  mounted() {
    axios
      .get("/api/category/content/" + this.$route.params.mainID)
      .then(res => {
        this.contents = res.data;
      });
  }
};
</script>
