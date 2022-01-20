<template>
    <div class="content">
        全局状态管理vuex-demo
        <p>官方文档&nbsp;<a href="https://vuex.vuejs.org/zh/">https://vuex.vuejs.org/zh/</a></p>
        <p style="height: 40px;line-height: 40px;display:flex;align-item:center; justify-content:center"><span style="font-size:16px">全局的count值: </span><span style="font-size:40px;marginLeft:10px">{{count}}</span></p>
        <button class="btn increment-btn"
                @click="increment">+1</button>
        <button class="btn subtract-btn"
                @click="subtract">-1</button>
        <button class="btn increment-btn"
                @click="incrementStep(payload)">+10</button>
        <button class="btn subtract-btn"
                @click="subtractStep(payload)">-10</button>
        <button class="btn increment-btn"
                @click="incrementStepAsync(payload)">Async +10</button>
        <button class="btn subtract-btn"
                @click="subtractStepAsync(payload)">Async -10</button>
    </div>
</template>

<script>
import { mapState, mapMutations, mapActions } from 'vuex';

export default {
    data() {
        return {
            payload: {
                step: 10
            }
        };
    },
    created() {},
    mounted() {},
    computed: {
        ...mapState('indexStore', {
            count: state => state.count
        })
    },
    watch: {},
    methods: {
        ...mapMutations('indexStore', [
            'increment',
            'subtract',
            'incrementStep',
            'subtractStep'
        ]),
        ...mapActions('indexStore', ['incrementStepAsync', 'subtractStepAsync'])
    }
};
</script>

<style scoped>
.content {
    font-size: 0.16rem;
    font-weight: 500;
    font-style: normal;
    color: #fff;
    line-height: 1.5;
}

.btn {
    border: none;
    margin: 0.2rem 0.2rem 0 0;
    height: 0.5rem;
    min-width: 1rem;
    background-color: #037fff;
    outline: none;
}
</style>