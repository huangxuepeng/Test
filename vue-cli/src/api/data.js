import axios from './axios'

export const getMenu = () => {
    return axios.request({
        url: '/u/v1/user/list',
        method: 'get'
    })
}
