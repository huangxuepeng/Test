<template>
    <div class="login_container">
       <div class="login_box">

           <!-- 头像盒子 -->
           <div class="avatar_box">
               <img src="../assets/logo.png" alt="头像">
           </div>

           <!-- 登录表单区 -->
           <el-form :model="loginFrom" label-width="0px" class="login_from">
               <!-- 账户 -->
                <el-form-item label="" >
                    <el-input v-model="loginFrom.username"></el-input>
                </el-form-item>

                 <!-- 密码 -->
                <el-form-item label="">
                    <el-input v-model="loginFrom.password" type="password"></el-input>
                </el-form-item>

                <!-- 按钮 -->
                <el-from-item class="btns">
                    <el-button type="primary">登录</el-button>
                    <el-button type="info">重置</el-button>
                </el-from-item>
           </el-form>
       </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            //这是登录表单的数据绑定对象
            loginFrom:{
                username: '', 
                password: '',
            }
        }
        
    }
}
</script>
<style lang="scss" scoped>
    .login_container {
        background: #2b4b6b;
        height: 100%;
    }
    .login_box {
        width: 450px;
        height: 300px;
        background-color: #fff;
        border-radius: 5px;
        position: absolute;
        left: 50%;
        top: 50%;
        transform: translate(-50%, -50%);
        .avatar_box {
            height: 130px;
            width: 130px;
            border: 1px solid #eee;
            border-radius: 50%;
            padding: 10px;
            box-shadow: 0 0 10px #ddd;
            position: absolute;
            left: 50%;
            transform: translate(-50%, -50%);
            background-color: #fff;
            img {
                height: 100%;
                width: 100%;
                border-radius: 50%;
                background: #eee;
            }
        }
    }

    .login_from {
        position: absolute;
        bottom: 0;
        width: 100%;
        padding: 0 10px;
        box-sizing: border-box;
    }

    .btns {
        display: flex;
        justify-content: flex-end;
    }
</style>