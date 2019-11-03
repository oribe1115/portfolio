<template lang="pug">
    .content-card
      .image
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
.content-card {
  border: 1px solid black;
  padding: 10px;
  width: 150px;
  height: 200px;
  margin: 15px;

  .image {
    height: 140px;
    width: 140px;
    margin: auto;
  }

  img {
    height: 100%;
  }
}
</style>
