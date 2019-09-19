<template lang="pug">
    .content-detail
        h1 Content Detail
        p contentID: {{ contentID }}
        p content: {{ content.title }}
        img(src=mainImage)
</template>
<script>
import axios from "axios";

export default {
  name: "ContentDetail",
  data() {
    return {
      contentID: this.$route.params.contentID,
      content: null,
      mainImage: ""
    };
  },
  mounted() {
    axios
      .get("/api/content/" + this.contentID)
      .then(res => (this.content = res.data));
    this.mainImage = this.content.image;
  }
};
</script>
