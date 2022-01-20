<template>
    <div class="content">
        1px边框问题兼容
        <div class="box">
            <div class="normal">上边框-正常</div>
            <div class="compatible">上边框-兼容</div>
        </div>
    </div>
</template>


<script>
</script>

<style lang="less" scoped>
@import "@/styles/borderPX.less";
// .borderPX;    (@position: border-bottom, @width: 1px, @color: #eaeaec)
.content {
    font-size: 0.16rem;
}

.normal {
    border-top: 1px solid #fff;
}
.compatible {
    .borderPX(border-top,1px,#fff);
}
</style>