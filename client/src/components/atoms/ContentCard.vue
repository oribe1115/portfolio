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
.image {
  height: 50px;
  width: 50px;
}

img {
  height: 100%;
  width: 100%;
}

.content-card {
  border: 1px solid black;
  padding: 10px;
  width: 150px;
  margin: 15px;
}
</style>
