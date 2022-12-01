<template>
    <div class="content">
        axios封装的使用
        <p class="result-text">请求返回的数据是 -- {{infoData}}</p>
    </div>
</template>

<script>
import { _exampleAPI } from '../server/api.js';
export default {
    data() {
        return {
            infoData: ''
        };
    },
    mounted() {
        this.initResponseData();
    },
    methods: {
        async initResponseData() {
            let res = await _exampleAPI();
            if (res.code !== 200) {
                console.log('请求出错了');
            }
            console.log(JSON.stringify(res));
            this.infoData = JSON.stringify(res.data);
            console.log(
                '请求完成,返回数据是-------' + JSON.stringify(res.data)
            );
        }
    }
};
</script>

<style scoped>
.content {
    font-size: 0.16rem;
    line-height: 0.32rem;
    font-style: normal;
    color: #333;
}

.result-text {
    color: #3c40c6;
    font-weight: 500;
}
</style>