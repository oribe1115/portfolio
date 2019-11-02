<template lang="pug">
    .content
        h1 Content
        p {{ content }}
</template>

<script>
import axios from "axios";

export default {
  name: "Content",
  data() {
    return {
      content: null
    };
  },
  mounted() {
    axios.get("/api/content/" + this.$route.params.contentID).then(res => {
      this.content = res.data;
    });
  }
};
</script>
