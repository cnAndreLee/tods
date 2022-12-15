import axios from "axios"
import storageService from '../service/storageService'
import config from "../config.js"

const requests = axios.create({
    // baseURL: "http://10.0.0.10:8000/",
    baseURL: config.backend,
    timeout: 2000,
})

// 请求拦截器
requests.interceptors.request.use( (config) => {
    let token = storageService.get(storageService.USER_TOKEN)
    if (token != null) {
        Object.assign(config.headers, { Authorization: `Bearer ${token}`})
    }
    return config;
}, (error) => {
    return Promise.reject(error);
});

// requests.interceptor.response.use( (res) => {
//   
//     return res.data;
// },(error) => {
//     return Promise.reject(new Error('faile'))
// })

export default requests;