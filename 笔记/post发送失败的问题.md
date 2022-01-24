# 1. 前言
对于一个只会用postman调试后端接口的菜鸟来说, 突然使用前端代码去调试还是挺痛苦的,虽然别人说vue的代码格式简单, 但是一言难尽....
# 2. 问题
这个问题是我在解决了[跨域](https://github.com/huangxuepeng/Test/blob/main/%E7%AC%94%E8%AE%B0/%E8%B7%A8%E5%9F%9F.md)问题之后发现的.发现GET请求就可以发送到前端数据, 但是等到了POST请求之后, 发现后端的虽然是写的接收数据, 但是后端打印却没有打印出前端发送的数据.   
后端代码: 
```
mobiles := ctx.PostForm("mobile") //拿到前端发送的手机号码
passwords := ctx.PostForm("password")
fmt.Println(mobiles, passwords)
```
打印的结果为空.   

  
我开始猜测出现这个问题的原因:
1. 后端的问题 ?   
   但是使用postman测试的接口就是没有问题呀, 无论是post, 还是get请求都是没有问题的呀, 所以排除了后端接口的问题;
2. 跨域问题没有彻底解决?   
    只有前端向后端发送数据的时候才会出现, 数据传送失败的现象, 当后端向前端传送数据的时候, 前端就可以正常打印, 而后尝试了delete等都是不能提交数据.所以排除跨域未解决的问题;  
3. 前端的问题?  
    因为不擅长前端, 所以一开始就不敢想, 但是最后实在没有办法, 从前端下手, 仔细检查, 百度... 应该是axios的地方出现了错误.   
    ```
    const { data: res } = await this.$http.post('http://localhost:8080/u/v1/user/login', 
    this.loginForm)
    ```
其实this.loginForm 打印出来是undefind, 所以换了它, 换成前后端传递的参数即可.    
解决方案:
```
 var params = new URLSearchParams()
            // 参数保存
            params.append('mobile', this.loginForm.mobile)
            params.append('password', this.loginForm.password)
            this.$refs.loginFormRef.validate(async valid => {
                if (!valid) return
                const { data: res } = await this.$http.post('http://localhost:8080/u/v1/user/login', params)
```
这样就可以解决问题.   
解决了问题我就去请教我哥, 这样改的原理是啥?
![图片](bcd7c2e81a3b4509f815b9c2b4b8dab.png)  

然后就直接把问题找出来, 可以一瞬间解决的, 哎...大佬就是大佬呀!
