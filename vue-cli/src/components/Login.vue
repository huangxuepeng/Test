<template>
    <div class="login-container">
        <div class="login-box">
            <!-- 头像区 -->
            <div class="avator-box">
                <img src="../assets/logo.png" alt="">
            </div>
            <!-- 登录表单区 -->
            <el-form ref="loginFormRef" :model="loginForm" :rules="loginFormRules" label-width="0px" class="login-form">
                <!-- 用户名 -->
                <el-form-item prop="mobile">
                    <el-input v-model="loginForm.mobile"></el-input>
                </el-form-item>
                <!-- 密码 -->
                <el-form-item prop="password">
                    <el-input v-model="loginForm.password"></el-input>
                </el-form-item>
                <!-- 按钮区域 -->
                 <el-form-item class="btns">
                    <el-button type="primary" @click="login">登录</el-button>
                    <el-button type="info" @click="resetLoginForm">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script>
export default {
    data () {
        return {
         // 这是登录表单的数据绑定对象
            loginForm: {
                mobile: '18088630924',
                password: 'i * i = 4'
            },
            // 这是表单验证规则对象
            loginFormRules: {
                // 验证用户名
                mobile: [
                    { require: true, message: '请输入手机', trigger: 'blur' },
                    { min: 11, max: 11, message: '手机的长度必须是11位', trigger: 'blur' }
                ],
                // 验证密码
                password: [
                    { require: true, message: '请输入密码', trigger: 'blur' },
                    { min: 6, max: 15, message: '密码的长度必须是6-15位', trigger: 'blur' }
                ]
            }
        }
    },
    methods: {
        // 点击重置按钮
        resetLoginForm () {
            this.$refs.loginFormRef.resetFields()
        },
        login () {
            var params = new URLSearchParams()
            // 参数保存
            params.append('mobile', this.loginForm.mobile)
            params.append('password', this.loginForm.password)
            this.$refs.loginFormRef.validate(async valid => {
                if (!valid) return
                const { data: res } = await this.$http.post('/user/login', params)
                if (res.code !== 200) {
                    return this.$message.error('登录失败')
                } else {
                    this.$message.success('登录成功')
                    // console.log(res)
                    // 将登陆成功后的token, 保存到客户端 sessionStorage 中
                    // 项目中除了登录之外的其他的API接口, 必须在登录之后才能访问
                    // token只应在当前网站打开期间生效, 所以将token保存在sessionStorage中
                    window.sessionStorage.setItem('token', res.token)
                    // 通过函数式编程跳转到后台主页, 路由地址是 /home
                    this.$router.push('/home')
                }
            })
        }
    }
}
</script>

<style scoped>
    .login-container {
        background-color: #2b4b6b;
        height: 100%;
    }
    .login-box {
        width: 450px;
        height: 300px;
        background-color: #fff;
        border-radius: 3px;
        position: absolute;
        left: 50%;
        top: 50%;
        transform: translate(-50%, -50%);
    }
    .avator-box{
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
    }
    img {
            width: 100%;
            height: 100%;
            border-radius: 50%;
            background-color: #eee;
        }
        .login-form {
            position: absolute;
            bottom: 0;
            width: 100%;
            padding: 0 20px;
            box-sizing: border-box;
        }
</style>
