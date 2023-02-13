import axios from "axios"

// 让请求在浏览器中允许跨域携带cookie
// axios.defaults.withCredentials = true

const service = axios.create({
    //公共路由
    // baseURL: 'http://127.0.0.1:8088/api',
    baseURL: 'http://101.201.51.200:8080/api',
    // 请求最大响应时间
    // timeout: 10000
    timeout: 1000000
})

//请求拦截器
service.interceptors.request.use(config => {
    //请求头带上token
    if (localStorage.getItem('token')) {
        config.headers.Authorization = localStorage.getItem('token')
    }
    return config
}, error => {
    console.log(error)
    return Promise.reject()
})

// 响应拦截器
service.interceptors.response.use(
    response => {
        if (response.status === 200) {
            return Promise.resolve(response.data)
        } else {
            return Promise.reject(response)
        }
    }, error => {
        alert("服务器开小差了，请稍后再试！");
        // localStorage.clear()
        setTimeout(() => {
            window.location.href = '/'
        }, 300)
    }
)

export default service