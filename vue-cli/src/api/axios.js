//二次封装
import axios from 'axios';
const baseUrl = '/api'

class httpRequest {

  constructor(baseUrl) {
    this.baseUrl = baseUrl
  }
  getInsideConfig() {
    const config = {
      baseUrl: this.baseUrl
    }
  }
    interceptors(instance) {
        // 添加请求拦截器
   instance.interceptors.request.use(function (config) {
      // 在发送请求之前做些什么
      return config;
    }, function (error) {
      console.log(error)
      // 对请求错误做些什么
      return Promise.reject(error)
    });

      // 添加响应拦截器
    instance.interceptors.response.use(function (response) {
      // 对响应数据做点什么
      return response;
    }, function (error) {
      // 对响应错误做点什么
      return Promise.reject(error);
    });
  }
  request () {
    const instance = axios.create()
    this.interceptors(instance)
    return instance()
  }
}