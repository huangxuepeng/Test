import axios from './axios'

export const getMenu = () => {
    return axios.request({
        url: 'http://192.168.10.12:8080/u/v1/user/list',
        method: 'get'
    })
}
