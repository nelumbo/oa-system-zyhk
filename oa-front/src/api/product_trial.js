import request from './base'

export const addProductTrial = (productTrial) => {
    return request({
        url: '/productTrial',
        method: 'POST',
        data: productTrial
    })
}

export const delProductTrial = (id) => {
    return request({
        url: '/productTrial/' + id,
        method: 'DELETE',
    })
}

export const editProductTrial = (productTrial) => {
    return request({
        url: '/productTrial',
        method: 'PUT',
        data: productTrial
    })
}

export const nextProductTrial = (productTrial) => {
    return request({
        url: '/productTrial/next',
        method: 'PUT',
        data: productTrial
    })
}

export const queryProductTrials = (model, pageData) => {
    return request({
        url: '/productTrials',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}