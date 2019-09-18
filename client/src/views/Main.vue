<template lang="pug">
    .main
        h1 main
        Sidebar
        .content-list(v-if="isMain")
            ContentsList(text="main")
        .content-list(v-if="isSub")
            ContentsList(text="sub")
        p isMain: {{ isMain }}
        p isSub: {{ isSub }}

</template>
<script>
import Sidebar from "@/components/Sidebar.vue";
import ContentsList from "@/components/ContentsList.vue";

import store from "../store.js";

export default {
  name: "Main",
  components: {
    Sidebar,
    ContentsList
  },
  data() {
    return {
      isMain: false,
      isSub: false
    };
  },
  mounted() {
    this.isMain = store.isMainCategory();
    this.isSub = store.isSubCategory();
  }
};
</script>
