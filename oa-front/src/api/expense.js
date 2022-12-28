import request from './base'

export const addExpense = (expense) => {
    return request({
        url: '/expense',
        method: 'POST',
        data: expense
    })
}

export const delExpense = (id) => {
    return request({
        url: '/expense/' + id,
        method: 'DELETE',
    })
}

export const approveExpense = (expense) => {
    return request({
        url: '/expense',
        method: 'PUT',
        data: expense
    })
}

export const queryExpense = (id) => {
    return request({
        url: '/expense/' + id,
        method: 'GET',
    })
}

export const queryExpenses = (model, pageData) => {
    return request({
        url: '/expenses',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}