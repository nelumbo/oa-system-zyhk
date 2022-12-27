import request from './base'

export const addCustomerCompany = (customerCompany) => {
    return request({
        url: '/customerCompany',
        method: 'POST',
        data: customerCompany
    })
}

export const delCustomerCompany = (id) => {
    return request({
        url: '/customerCompany/' + id,
        method: 'DELETE',
    })
}

export const editCustomerCompany = (customerCompany) => {
    return request({
        url: '/customerCompany',
        method: 'PUT',
        data: customerCompany
    })
}

export const queryCustomerCompanys = (model, pageData) => {
    return request({
        url: '/customerCompanys',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}