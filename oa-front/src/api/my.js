import request from './base'

export const queryTopList = () => {
    return request({
        url: '/topList',
        method: 'GET',
    })
}

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

export const queryMyPredesigns = (model, pageData) => {
    return request({
        url: '/my/predesigns',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryMyPredesignTasks = (model, pageData) => {
    return request({
        url: '/my/predesignTasks',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryMyContracts = (model, pageData) => {
    return request({
        url: '/my/contracts',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryMySaveContracts = (model, pageData) => {
    return request({
        url: '/my/contracts/save',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryMyTasks = (model, pageData) => {
    return request({
        url: '/my/tasks',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryMySavePurchasings = (model) => {
    return request({
        url: '/my/purchasings/save',
        method: 'POST',
        data: model
    })
}

export const queryMyHistorys = (model, pageData) => {
    return request({
        url: '/my/historys',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryMyProductTrials = (model, pageData) => {
    return request({
        url: '/my/productTrials',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const updatePwd = (model) => {
    return request({
        url: '/my/pwd',
        method: 'POST',
        data: model,
    })
}