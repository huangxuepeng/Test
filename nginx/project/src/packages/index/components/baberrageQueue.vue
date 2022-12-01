<template>
    <div class="baberrage-box">
        HELLO
    </div>
</template>

<script>
export default {
    name: 'v-baberrage',
    props: {
        liNum: {
            type: Number,
            default: 5
        },
        liHeight: {
            type: Number,
            default: 50
        }
    }
};
</script>

<style lang="less" scoped>
.baberrage-box {
}
</style>