<template lang="pug">
  .sub-category
    .path
      router-link(:to="{ name: 'mainCategory', params: { mainID: main.id } }")
        .main
          h1 {{ main.name }}
      h1 >
      router-link(:to="{ name: 'subCategory', params: { subID: sub.id } }")
        .sub
          h1 {{ sub.name }}
    .description
      markdown-it-vue(:content="String(sub.description)")
    ContentsList(:contents="contentsList.contents")
</template>

<script>
import ContentsList from "@/components/ContentsList.vue";
import axios from "axios";
import MarkdownItVue from "markdown-it-vue";
import "markdown-it-vue/dist/markdown-it-vue.css";

export default {
  name: "SubCategory",
  components: {
    ContentsList,
    MarkdownItVue
  },
  data() {
    return {
      contentsList: Array,
      main: Object,
      sub: Object
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
      axios
        .get("/api/category/content/sub/" + this.$route.params.subID)
        .then(res => {
          this.contentsList = res.data;
          this.main = this.contentsList.main_category;
          this.sub = this.contentsList.sub_category;
        });
    }
  }
};
</script>

<style lang="scss">
.sub-category {
  padding: 0px 30px;
}

.path {
  display: flex;

  a {
    &:hover {
      color: gray;
    }
  }
}

h1 {
  margin: 5px;
}
</style>
