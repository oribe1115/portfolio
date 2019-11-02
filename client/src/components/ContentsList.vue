<template lang="pug">
    .contents-list
        h1 ContentsList
        p {{ category }}
</template>

<script lang="ts">
export default {
  name: "Contentslist",
  props: {
    category: String
  }
};
</script>
