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

.content-card {
  border: 1px solid black;
  padding: 10px;
  width: 150px;
  margin: 15px;
}
</style>
