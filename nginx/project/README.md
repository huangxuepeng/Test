# vue-template-pc
vue脚手架-模板

### 启动方式:

* 将项目拉到本地,安装依赖(npm i); 建议装上yarn 和 nrm
* 开发环境启动项目  npm run dev
* 项目打包 npm run build
* 项目指令化创建页面 npm run create-page  详细脚本代码 /template/script.js  和/template/util.js



2020/1/8 防抖 | 节流 指令和函数 和使用参考

2020/1/10 axios二次封装 和 webpack-dev-server配置proxy实现开发模式下的本地代理(可解决跨域问题)

2020/1/11 vuex简单配置和使用 vuex的模块的使用 

2020/1/16 补充时间格式化函数和千位打点工具函数

2020/1/17 优化按钮防抖函数和节流函数
        防抖函数是首部触发还是尾部触发,根据常规交互逻辑,默认置为首部触发
        节流函数

2020/1/26 完成指令化创建页面,并将新页面的路由自动的导入全局,创建新页面完成之后重新获取所有的路由配置文件,并将配置文件整理输出到page_router.json文件中,方便维护的时候进行查看


### 下阶段需要完成的
* 命令化生成页面模板(包括页面根据用户输入将模板进行插值配置,并增加对应的路由和配置)    √
* ~~多个多个请求重叠只出现一个loading(axios的封装优化)~~                            
* 增加常见指令(动画交替,复制指令,图片预加载和懒加载)                                待补充+
* webpack 按需加载页面组件                                                        √
* 根据全局状态管理权限(两种接口模拟 yapi或者是本地的json文件导入(这个需要加一个setTimeout进行延时,模拟接口的延迟))                                                               √
* 上传组件开发
* 拖放组件开发
* 针对大量数据的渲染 比如10w条的渲染  如果每条数据都创建了dom节点,浏览器肯定是承受不了的(chrome 推荐的是页面不超过1500个)   时间切片方案  虚拟滚动方案
* 增加两种移动端调试工具 vconsole 和 chii
    vconsole需要在根目录下面的.env下配置 IS_VCONSOLE=true
    chii需要在根目录下面的.env下配置IS_CHII=true 然后运行npm run chii 即可本地启动调试服务

```
sass-loader和node-sass版本对应
sass-loader 4.1.1，node-sass 4.3.0

sass-loader 7.0.3，node-sass 4.7.2

sass-loader 7.3.1，node-sass 4.7.2

sass-loader 7.3.1，node-sass 4.14.1

sass-loader 10.0.1，node-sass 6.0.1
```