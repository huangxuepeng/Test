<template>
    <div class="content">
        获取url的中参数工具
        <p>当前url地址中的query为<span class="query-text">{{queryString}}</span></p>
    </div>
</template>

<script>
import { _getUrlQuery } from '@/utils/getUrlQuery.js';
export default {
    data() {
        return {
            queryString: ''
        };
    },
    created() {},
    mounted() {
        this.getQueryString();
    },
    computed: {},
    watch: {},
    methods: {
        getQueryString() {
            let res = _getUrlQuery();
            this.queryString = JSON.stringify(res);
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

.query-text {
    color: #7158e2;
    margin-left: 0.2rem;
}
</style>