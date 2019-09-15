<template lang="pug">
    h1 here is mainCategory
    //- .div
    //- <sidebar></sidebar>
    .sidebar
        sidebar()
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import Sidebar from '../../components/Sidebar.vue'

@Component({
    components: {
        Sidebar
    }
})
export default class MainCategory extends Vue {
}

// export default {
//   components: {
//     Sidebar
//   }
// }
</script>
