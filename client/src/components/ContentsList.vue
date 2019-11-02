<template lang="pug">
    .contents-list
        h1 ContentsList
        p {{ category }}
        p {{ categoryID }}
        p {{ basePath}}
        p {{ contentsList }}
</template>

<script>
import axios from "axios";

export default {
  name: "Contentslist",
  props: {
    category: String,
    categoryID: String,
    addPath: String
  },
  data() {
    return {
      basePath: "/api/category/content/",
      contentsList: null
    };
  },
  mounted() {
    this.contentsList = null;
    if (this.category == "sub") {
      this.basePath = "/api/category/content/sub/";
    }
    axios.get(this.basePath + this.categoryID).then(res => {
      this.contentsList = res.data;
    });
  }
};
</script>
