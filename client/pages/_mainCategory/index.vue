<template lang="pug">
    .mainCategory
        h1 here is main category
        .sidebar
            Sidebar
</template>

<script>

import Sidebar from "@/components/Sidebar.vue"

export default {
    name: "MainCategory",
    components: {
        Sidebar
    }
}
</script>