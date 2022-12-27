import request from './base'

export const login = (queryObject) => {
    return request({
        url: '/login',
        method: 'POST',
        data: queryObject,
    })
}

export const tokenLogin = () => {
    return request({
        url: '/tokenLogin',
        method: 'GET',
    })
}