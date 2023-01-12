import request from './base'

export const queryHistoryOffices = (model, pageData) => {
    return request({
        url: '/historyOffices',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryHistoryEmployees = (model, pageData) => {
    return request({
        url: '/historyEmployees',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}
