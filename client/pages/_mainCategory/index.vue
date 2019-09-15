<template lang="pug">
    .mainCategory
        h1 here is main category
        h2 mainCategory: {{ mainID }}
        .sidebar
            Sidebar
</template>

<script>

import Sidebar from "@/components/Sidebar.vue"

export default {
    name: "MainCategory",
    components: {
        Sidebar
    },
    mounted(){
        const mainID = this.$route.params.mainCategory
    }
}
</script>