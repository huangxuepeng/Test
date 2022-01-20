<template>
    <div class="content">
        防抖指令的使用
        <p>数字num: {{tnum}}</p>
        <van-button type="primary"
                    v-debounce="debounceTest">按钮</van-button>
    </div>
</template>
<script>
import { mixin } from '@/mixins/instructions.js';
export default {
    mixins: [mixin],
    data() {
        return {
            tnum: 0
        };
    },
    created() {},
    methods: {
        debounceTest() {
            console.log('打印文字', this.tnum++);
        }
    }
};
</script>

<style scoped>
.content {
    font-size: 0.16rem;
    font-style: normal;
    font-weight: 500;
    line-height: 0.3rem;
}
</style>