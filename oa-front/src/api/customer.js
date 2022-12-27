import request from './base'

export const addCustomer = (customer) => {
    return request({
        url: '/customer',
        method: 'POST',
        data: customer
    })
}

export const delCustomer = (id) => {
    return request({
        url: '/customer/' + id,
        method: 'DELETE',
    })
}

export const editCustomer = (customer) => {
    return request({
        url: '/customer',
        method: 'PUT',
        data: customer
    })
}

export const queryCustomers = (model, pageData) => {
    return request({
        url: '/customers',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}