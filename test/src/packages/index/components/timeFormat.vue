<template>
    <div class="content">
        时间格式化工具使用
        <p>当前时间为<span class="time-text">{{getDateNow('YY-MM-DD hh:mm:ss')}}</span></p>
        <p>当前时间为<span class="time-text">{{getDateNow('YY-MM-DD hh:mm')}}</span></p>
        <p>当前时间为<span class="time-text">{{getDateNow('YY-MM-DD')}}</span></p>
        <p>当前时间为<span class="time-text">{{getDateNow('YY/MM/DD hh:mm:ss')}}</span></p>
    </div>
</template>


<script>
import { _timeFormat } from '@/utils/timeFormat.js';
export default {
    data() {
        return {};
    },
    created() {},
    mounted() {},
    computed: {},
    watch: {},
    methods: {
        getDateNow(type) {
            let timeStamp = Date.now() || new Date().getTime();
            return _timeFormat(timeStamp, type);
        }
    }
};
</script>

<style scoped>
.content {
    font-size: 0.16rem;
}
.time-text {
    margin-left: 0.2rem;
    font-size: 0.16rem;
    line-height: 1.5;
    font-style: normal;
    font-weight: 500;
    color: #c56cf0;
}
</style>