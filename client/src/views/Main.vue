<template lang="pug">
    .main
        h1 main
        Sidebar
        ContentsList
</template>
<script>
import Sidebar from "@/components/Sidebar.vue";
import ContentsList from "@/components/ContentsList.vue";

export default {
  name: "Main",
  components: {
    Sidebar,
    ContentsList
  }
};
</script>
