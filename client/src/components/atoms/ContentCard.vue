<template lang="pug">
    .content-card
      img(:src="content.image")
      .title
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
  width: 210px;
  height: 280px;
  margin: 30px;
  background-color: #ffffff;
  border-radius: 4px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.4);
  transition: all 0.5s cubic-bezier(0.23, 1.5, 0.32, 1);

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.4);
  }

  img {
    height: 210px;
    width: 210px;
    object-fit: cover;
    border-radius: 4px 4px 0px 0px;
  }

  a {
    color: #24292e;
  }

  .title {
    max-width: 210px;
    margin: auto;
    text-align: center;
    padding: 10px;

    h3 {
      padding: 0px;
      margin: 0px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
  }
}
</style>
