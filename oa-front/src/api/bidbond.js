import request from './base'

export const addBidbond = (bidbond) => {
    return request({
        url: '/bidbond',
        method: 'POST',
        data: bidbond
    })
}

export const delBidbond = (id) => {
    return request({
        url: '/bidbond/' + id,
        method: 'DELETE',
    })
}

export const editBidbond = (bidbond) => {
    return request({
        url: '/bidbond',
        method: 'PUT',
        data: bidbond
    })
}

export const approveBidbond = (bidbond) => {
    return request({
        url: '/bidbond/approve',
        method: 'PUT',
        data: bidbond
    })
}

export const queryBidbonds = (model, pageData) => {
    return request({
        url: '/bidbonds',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}