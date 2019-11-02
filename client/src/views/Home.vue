<template lang="pug">
  .home
    <img alt="Vue logo" src="../assets/logo.png" />
    Sidebar
    .main(v-if="category == ''")
      About
    .main(v-else)
      ContentsList(:category="category" :categoryID="categoryID" :addPath="addPath")
    
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
      categoryID: "",
      addPath: ""
    };
  }
};
</script>
