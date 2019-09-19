<template lang="pug">
    .all-contents-list
        h1 All Contents List
        p {{ contents }}
</template>
<script>
import axios from "axios";

export default {
  name: "AllContentsList",
  data() {
    return {
      contents: []
    };
  },
  mounted() {
    axios.get("/api/content").then(res => (this.contents = res.data));
  }
};
</script>
