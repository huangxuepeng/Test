<template>
    <div>
        <!-- 面包屑 -->
        <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>用户管理</el-breadcrumb-item>
        <el-breadcrumb-item>用户列表</el-breadcrumb-item>
        </el-breadcrumb>
            <!-- 卡片视图 -->
        <el-card >
            <!-- 搜索区 -->
            <el-row :gutter="20">
            <el-form ref="loginFormRef" :model="selectForm" :rules="selectFormRules" label-width="0px">
                <!-- 用户名 -->
                <el-form-item prop="mobile">
                    <el-input v-model="selectForm.content" placeholder="请输入学号"></el-input>
                </el-form-item>
                <!-- 按钮区域 -->
                 <el-form-item class="btns">
                    <el-button type="primary" @click="getUserList">查询</el-button>
                </el-form-item>
            </el-form>
            </el-row>
            <el-descriptions title="用户信息" v-for="item in listStudent" :key="item.name">
                <el-descriptions-item label="用户名">{{ item.Name }}</el-descriptions-item>
                <el-descriptions-item label="手机号">{{ item.Mobile }}</el-descriptions-item>
                <el-descriptions-item label="学号">{{ item.StuNum }}</el-descriptions-item>
                <el-descriptions-item label="学校">
                <el-tag size="small">长春大学</el-tag>
                </el-descriptions-item>
                <el-descriptions-item label="联系地址">长春大学校</el-descriptions-item>
            </el-descriptions>
        </el-card>
    </div>
</template>

<script>
export default {
    data () {
        return {
            // 学生数据
            listStudent: [],

            selectForm: {
                content: ''
            },
            selectFormRules: {
                content: [
                    { require: true, message: '请输入要查询的学号', trigger: 'blur' },
                    { max: 9, message: '学号的长度最多是9位', trigger: 'blur' }
                ]
            }
        }
    },
    mounted () {
        this.getUserList()
    },
    methods: {
        // 获取所有的学生信息
       getUserList () {
           this.$http.get(`/user/list?query=${this.selectForm.content}`).then((response) => {
               this.listStudent = response.data.data.studentList
           })
    
       }
    }
}
</script>

<style  scoped>

</style>
