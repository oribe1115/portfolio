<template lang="pug">
  .app
    .sidebar
      Sidebar
    .view
      router-view
</template>

<script>
import Sidebar from "@/components/Sidebar.vue";

export default {
  components: {
    Sidebar
  }
};
</script>

<style>
.app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  display: flex;
}

.sidebar {
  width: 150px;
}

.view {
  width: 800px;
}
</style>
