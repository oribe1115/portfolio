<template lang="pug">
    .contents-list
        h1 ContentsList
        p {{ category }}
        p {{ categoryID }}
</template>

<script lang="ts">
export default {
  name: "Contentslist",
  props: {
    category: String,
    categoryID: String
  }
};
</script>
