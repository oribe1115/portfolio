<template lang="pug">
    .sub-category
        h1 Sub Category
        ContentsList(:contents="contents")
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
      contents: null
    };
  },
  mounted() {
    axios
      .get("/api/category/content/sub/" + this.$route.params.subID)
      .then(res => {
        this.contents = res.data;
      });
  }
};
</script>
