<template>
  <div>
    <el-container class="home-container">
      <!-- 头部区 -->
  <el-header>
    <div>
      <img src="../assets/logo.png" alt="">
      <span>作业</span>
    </div>
    <el-button type="info" @click="logout">退出</el-button>
  </el-header>
  <!-- 页面主体区 -->
  <el-container>
    <!-- 侧边栏 -->
    <el-aside width="200px">
      <!-- 侧边栏菜单 -->
       <el-menu
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      :router="true"
      default-active="item.path"
      >
      <el-submenu index="1">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span>导航一</span>
        </template>
        <el-menu-item-group>
          <el-menu-item :index= "item.path + ''" v-for="item in menulist" path="item.path" :key="item.name">{{ item.name }}</el-menu-item>
        </el-menu-item-group>
      </el-submenu>
    </el-menu>
    </el-aside>
    <!-- 右侧内容主体 -->
    <el-main>
      <!-- 路由占位符 -->
      <router-view></router-view>
    </el-main>
  </el-container>
</el-container>
  </div>
</template>

<script>
export default {
  data () {
    return {
        menulist: [
          {
            name: '查询列表',
            path: '/list'
          },
          {
            name: '删除信息',
            path: '/delete'
          },
          {
            name: '添加信息',
            path: '/add'
          }
        ],
        // 被激活的地址
        activePath: ''
    }
  },
  create () {
    this.activePath = window.sessionStorage.getItem('activePath')
  },
  methods: {
    logout () {
      window.sessionStorage.clear()
      this.$router.push('/login')
    },
    // 保存连接的激活状态
    saveNavState(activePath){
      window.sessionStorage.setItem('activePath', activePath)
      this.activePath = activePath
    }
  }
}
</script>

<style scoped>
.home-container {
  position:absolute;
  height: 100%;
  width:100%;
}
.el-header {
  background-color: #373d41;
  display: flex;
  justify-content: space-between;
  color: #fff;
  font-size: 20px;
}
.el-aside {
  background-color: #333744;
}
.el-main {
  background-color: #eaedf1;
}
 img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  background-color: #eee;
}
</style>
