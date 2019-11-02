<template lang="pug">
    .contents-list
        h1 ContentsList
        p {{ contents }}
        .list(v-for="content in contents" :key="content.id")
          router-link(:to="{ name: 'content', params: { contentID: content.id } }")
            .content
              p {{ content.title }}
</template>

<script>
export default {
  name: "Contentslist",
  props: {
    contents: Array
  }
};
</script>
