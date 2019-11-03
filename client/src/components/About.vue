<template lang="pug">
  .about
    h1 about
    .content
      markdown-it-vue(:content="about.content")
</template>

<script>
import axios from "axios";
import MarkdownItVue from "markdown-it-vue";
import "markdown-it-vue/dist/markdown-it-vue.css";

export default {
  name: "About",
  components: {
    MarkdownItVue
  },
  data() {
    return {
      about: null
    };
  },
  created() {
    this.fetchData();
  },
  watch: {
    $route: "fetchData"
  },
  methods: {
    fetchData: function() {
      axios.get("/api/generalData/about").then(res => {
        this.about = res.data;
      });
    }
  }
};
</script>

<style lang="scss">
.content {
  text-align: left;
  padding: 30px;
}
</style>
