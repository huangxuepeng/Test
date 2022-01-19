<template>
    <div>
        <el-row class="home" :getter="20">
            <el-col :span="8">
                <el-card shadow="hover">
                    <div class="user">
                        <img :src="userImg" alt="图片">
                        <div class="userInfo">
                            <p class="name">Admin</p>
                            <p class="access">超级管理员</p>
                        </div>
                    </div>
                    <div class="login-info">
                        <p>上次登录时间: <span>2021-7-19</span></p>
                        <p>上次登录地点: <span>你家</span></p>
                    </div>
                </el-card>
                <el-card style="height: 470px; margin-top: 20px">
                    <el-table :data="tableData">
                            <el-table-column
                            show-overflow-tooltip
                            v-for="(val, key) in tableLable"
                            :key="key"
                            :prop="key"
                            :lable="val"
                            >
                            </el-table-column>
                    </el-table>
                </el-card>
            </el-col>
            <el-col :span="16">
            </el-col>
        </el-row>
    </div>
</template>

<script>
import { getMenu } from '../../api/data'
    export default {
        data() {
            return {
                userImg: require("../../assets/logo.png"),
                tableData: [
                    {
                        name: "kjsi",
                        todyBuy: 100,
                        monthBuy: 300,
                        totalBuy: 800
                    },
                    {
                        name: "kjssssi",
                        todyBuy: 100,
                        monthBuy: 300638888,
                        totalBuy: 800
                    },
                    {
                        name: "kjvbi",
                        todyBuy: 100,
                        monthBuy: 300,
                        totalBuy: 800
                    },
                    {
                        name: "vbi",
                        todyBuy: 100,
                        monthBuy: 300,
                        totalBuy: 800
                    },
                    {
                        name: "ffgf",
                        todyBuy: 100,
                        monthBuy: 300,
                        totalBuy: 800
                    },
                    {
                        name: "kf",
                        todyBuy: 100,
                        monthBuy: 300,
                        totalBuy: 800
                    }
                ],
                tableLable: {
                    name: "课程",
                    todayBuy: "今日购买",
                    monthBuy: "本月购买",
                    totalBuy: "总购买"
                }
            }
        },
        mounted() {
           getMenu().then((res) => {
               console.log(res)
           })
        }
    }
</script>
