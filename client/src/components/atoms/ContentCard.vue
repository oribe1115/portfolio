<template lang="pug">
    .content-card
        img(:src="content.image")
        h3 {{ content.title }}
</template>

<script>
export default {
  name: "ContentCard",
  props: {
    content: Object
  }
};
</script>

<style lang="scss">
img {
  width: 100px;
}
</style>
