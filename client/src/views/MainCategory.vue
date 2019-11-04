<template lang="pug">
    .main-category
      .path
        router-link(:to="{ name: 'mainCategory', params: { mainID: main.id } }")
          .main
            h1 {{ main.name }}
      .description
        markdown-it-vue(:content="String(main.description)")
      ContentsList(:contents="contentsList.contents")
</template>

<script>
import ContentsList from "@/components/ContentsList.vue";
import axios from "axios";
import MarkdownItVue from "markdown-it-vue";
import "markdown-it-vue/dist/markdown-it-vue.css";

export default {
  name: "MainCategory",
  components: {
    ContentsList,
    MarkdownItVue
  },
  data() {
    return {
      contentsList: Array,
      main: Object
    };
  },
  mounted() {
    axios
      .get("/api/category/content/" + this.$route.params.mainID)
      .then(res => {
        this.contentsList = res.data;
        this.main = this.contentsList.main_category;
      });
  }
};
</script>

<style lang="scss">
.main-category {
  padding: 0px 30px;
  color: #2c3e50;

  .path {
    a {
      text-decoration: none;

      &:visited {
        color: #2c3e50;
      }

      &:link {
        color: #2c3e50;
      }

      &:hover {
        color: gray;
      }
    }
  }
}
</style>
