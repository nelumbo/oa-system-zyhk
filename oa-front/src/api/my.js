import request from './base'

export const queryMe = () => {
    return request({
        url: '/me',
        method: 'GET',
    })
}

export const queryMyExpenses = (model, pageData) => {
    return request({
        url: '/my/expenses',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryMyBidbonds = (model, pageData) => {
    return request({
        url: '/my/bidbonds',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}