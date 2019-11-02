<template lang="pug">
    .contents-list
        h1 ContentsList
        p {{ contents }}
</template>

<script>
export default {
  name: "Contentslist",
  props: {
    contents: Array
  }
};
</script>
