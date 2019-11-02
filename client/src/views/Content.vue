<template lang="pug">
    .content
      .title
        h1 {{ content.title }}
      .image
        img(:src="content.image")
      .sub-image-list
        .sub-image(v-for="subImage in content.sub_images" :key="subImage.id")
          img(:src="subImage.url")
      .created-at
        p {{ content.date | date }}制作
      .description
        markdown-it-vue(:content="content.description")
</template>

<script>
import axios from "axios";
import MarkdownItVue from "markdown-it-vue";
import "markdown-it-vue/dist/markdown-it-vue.css";
import moment from "moment";

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
  created() {
    this.fetchData();
  },
  watch: {
    $route: "fetchData"
  },
  methods: {
    fetchData: function() {
      axios.get("/api/content/" + this.$route.params.contentID).then(res => {
        this.content = res.data;
      });
    }
  },
  filters: {
    date: function(created_at) {
      return (
        moment(created_at).get("year") +
        "年" +
        moment(created_at).get("month") +
        "月"
      );
    }
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

.sub-image-list {
  display: flex;
  justify-content: center;
  padding: 20px;
}

.sub-image {
  height: 100px;
}

.created-at {
  text-align: right;
}

.description {
  text-align: left;
  padding: 20px;
}
</style>
