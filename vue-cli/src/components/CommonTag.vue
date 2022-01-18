<template>
    <div class="tabs">
        <el-tag
        v-for="(tag) in tag"
        :key="tag.name"
        >
        {{ tag.label }}
        </el-tag>
    </div>
</template>

<script>
import { mapState } from "vuex"
    export default {
        computed: {
            ...mapState({
            taps: (state) => state.tab.tabsList
            })
        }
    }
</script>
