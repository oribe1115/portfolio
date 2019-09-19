<template lang="pug">
    .content-detail
        h1 Content Detail
        p contentID: {{ contentID }}
</template>
<script>
export default {
  name: "ContentDetail",
  data() {
    return {
      contentID: this.$route.params.contentID
    };
  }
};
</script>
