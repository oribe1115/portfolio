<template lang="pug">
    .content
      .title
        h1 {{ content.title }}
      .image
        img(:src="content.image")
      .description
        p {{ content.description }}
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

<style lang="scss">
.content {
  padding: 30px;
}

.title {
  font-size: 20px;
  text-align: left;
}

.image {
  height: 200px;
  margin: auto;
}

img {
  height: 100%;
  width: auto;
}

.description {
  text-align: left;
  padding: 20px;
}
</style>
