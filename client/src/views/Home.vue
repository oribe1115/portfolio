<template lang="pug">
  .home
    About
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
import Sidebar from "@/components/Sidebar.vue";
import About from "@/components/About.vue";

export default {
  name: "home",
  components: {
    // HelloWorld,
    Sidebar,
    About
  }
};
</script>
