<template>
    <div class="index-container animate__animated animate__bounceInUp">
        <header class="header">
            <div class="plate-title">This is a app is created with this
                <a href="https://github.com/hhcol620/create-vue-app">cva-cli</a>
            </div>
        </header>
        <nav class="nav">
            <router-link class="nav-item"
                         v-for="(item,index) in navList"
                         :key="item.label+index"
                         :to="item.url">
                <svg class="svg-icon"
                     aria-hidden="true">
                    <use :xlink:href="item.icon"></use>
                </svg>
                <div class="icon-text">{{item.label}}</div>
            </router-link>
        </nav>
        <section>
            <div class="plate-title">Example</div>
            <div class="show-content">
                <router-view></router-view>
            </div>
        </section>
    </div>
</template>

<script>
export default {
    data() {
        return {
            // 导航入口
            navList: [
                {
                    label: 'GetQuery',
                    icon: '#iconliaotian',
                    url: '01'
                },
                {
                    label: 'timeFormat',
                    icon: '#iconshuqian',
                    url: '02'
                },
                {
                    label: '防抖指令',
                    icon: '#iconzixun',
                    url: '03'
                },
                {
                    label: '封装请求',
                    icon: '#iconquanxuan',
                    url: '04'
                },
                {
                    label: '全局Store',
                    icon: '#icondingwei',
                    url: '05'
                },
                {
                    label: 'pxBorder',
                    icon: '#iconliaotian',
                    url: '06'
                }
            ]
        };
    }
};
</script>
<style lang="less" scoped>
.index-container {
    padding: 0 0.2rem;
    .header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        font-size: 0.16rem;
        font-weight: 500;
        font-style: normal;
        line-height: 0.24rem;
        color: #4d4d4d;
    }
    .plate-title {
        color: #4d4d4d;
        font-size: 0.36rem;
        font-weight: 600;
        line-height: 1.5;
        padding-top: 0.2rem;
    }
    .nav {
        display: flex;
        flex-wrap: wrap;
        .nav-item {
            width: 20%;
            margin-top: 0.1rem;
            display: flex;
            flex-direction: column;
            align-items: center;
            .icon-text {
                color: #4d4d4d;
                font-size: 0.16rem;
                font-weight: 500;
                line-height: 1.5;
            }
        }
    }
    .show-content {
        height: 3.4rem;
        background-color: #ffab73;
        border-radius: 0.1rem;
        margin-top: 0.1rem;
        padding: 0.1rem;
    }
}
</style>
