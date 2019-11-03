<template lang="pug">
  .app
    .header
      Header
    .container
      .sidebar-space
        Sidebar
      .view
        router-view
</template>

<script>
import Sidebar from "@/components/Sidebar.vue";
import Header from "@/components/Header.vue";

export default {
  components: {
    Sidebar,
    Header
  }
};
</script>

<style>
.app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  padding: 10px;
}

.container {
  display: flex;
}

.sidebar-space {
  width: 200px;
  margin-left: 30px;
}

.view {
  width: 900px;
  padding: 10px;
}
</style>
