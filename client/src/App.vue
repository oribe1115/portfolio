<template lang="pug">
  .app
    router-link(:to="{ name: 'home' }")
      p Home
    .container
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
}

.container {
  display: flex;
}

.sidebar {
  width: 150px;
}

.view {
  width: 800px;
}
</style>
