<template lang="pug">
    .header
        .site-name
            router-link(:to="{ name: 'home' }")
                h1 {{ name.content }}のポートフォリオ
</template>

<script>
import axios from "axios";

export default {
  name: "Header",
  data() {
    return {
      name: Object
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData: function() {
      axios.get("/api/generalData/name").then(res => {
        this.name = res.data;
      });
    }
  }
};
</script>

<style lang="scss">
.header {
  text-align: left;
  padding: 10px;

  h1 {
    font-size: 30px;
  }

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
</style>
