<template lang="pug">
    .content
      .title
        h1 {{ content.title }}
      .image
        img(:src="content.image")
      .description
        markdown-it-vue(:content="content.description")
</template>

<script>
import axios from "axios";
import MarkdownItVue from "markdown-it-vue";
import "markdown-it-vue/dist/markdown-it-vue.css";

export default {
  name: "Content",
  components: {
    MarkdownItVue
  },
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
