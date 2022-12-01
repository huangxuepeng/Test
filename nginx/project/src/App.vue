<template>
    <div class="main-container">
        <router-view></router-view>
    </div>
</template>
<script>
export default {
    name: 'App',
    components: {},
    created() {
        this.test();
    },
    methods: {
        test() {
            console.log('./App.vue');
        }
    }
};
</script>
<style scoped>
.main-container {
    height: 100%;
    width: 100%;
    overflow: hidden;
}
</style>