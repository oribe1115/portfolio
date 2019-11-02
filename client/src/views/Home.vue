<template lang="pug">
  .home
    <img alt="Vue logo" src="../assets/logo.png" />
    Sidebar
    .main(v-if="category == 'main'")
      ContentsList(:category="category" :categoryID="categoryID")
    .main(v-else)
      About
    
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
import Sidebar from "@/components/Sidebar.vue";
import About from "@/components/About.vue";
import ContentsList from "@/components/ContentsList.vue";

export default {
  name: "home",
  components: {
    // HelloWorld,
    Sidebar,
    About,
    ContentsList
  },
  data() {
    return {
      category: "",
      categoryID: ""
    };
  }
};
</script>
