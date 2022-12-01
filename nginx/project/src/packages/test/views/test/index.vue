<template>
    <div>
        <div class="desc">
            <button @click="test0">按钮一</button>
            <button @click="test1">按钮二</button>
        </div>
    </div>
</template>

<script>
import { test8080, test8081 } from '../../server/api.js';
export default {
    data() {
        return {};
    },
    methods: {
      async test0() {
         const res = await test8080();
         console.log(res);
        }, 
        async test1() {
         const res = await test8081();
         console.log(res);
        }
    }   
};
</script>
<style lang="less" scoped>
</style>
