<template lang="pug">
    .contents-list
        h1 ContentsList
        p {{ contentsList }}
</template>

<script>
import axios from "axios";

export default {
  name: "Contentslist",
  props: {
    path: String
  },
  data() {
    return {
      contentsList: null
    };
  },
  mounted() {
    axios.get(this.path).then(res => {
      this.contentsList = res.data;
    });
  }
};
</script>
