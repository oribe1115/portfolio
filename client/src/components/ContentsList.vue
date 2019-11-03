<template lang="pug">
    .contents-list
        .list(v-for="content in contents" :key="content.id")
          router-link(:to="{ name: 'content', params: { contentID: content.id } }")
            ContentCard(:content="content")
</template>

<script>
import ContentCard from "@/components/atoms/ContentCard.vue";

export default {
  name: "Contentslist",
  components: {
    ContentCard
  },
  props: {
    contents: Array
  }
};
</script>

<style lang="scss">
.contents-list {
  display: flex;
}
</style>
