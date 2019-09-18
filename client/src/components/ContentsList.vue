<template lang="pug">
    .content-list
        h1 Content List
        p {{ text }}
        p {{ contents }}
</template>

<script>
import axios from "axios";
import store from "../store.js";

export default {
  name: "ContentsList",
  props: {
    text: String
  },
  data() {
    return {
      contents: []
    };
  },
  mounted() {
    if (store.isMainCategory()) {
      axios
        .get("/api/category/content/" + store.state.categoryID)
        .then(res => (this.contents = res.data));
    } else if (store.isSubCategory()) {
      axios
        .get("/api/category/content/sub/" + store.state.categoryID)
        .then(res => (this.contents = res.data));
    }
  }
};
</script>
