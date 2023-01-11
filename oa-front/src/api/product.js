import request from './base'

export const addProduct = (product) => {
    return request({
        url: '/product',
        method: 'POST',
        data: product
    })
}

export const editProductBase = (product) => {
    return request({
        url: '/product/base',
        method: 'PUT',
        data: product
    })
}

export const editProductAttribute = (product) => {
    return request({
        url: '/product/attribute',
        method: 'PUT',
        data: product
    })
}

export const editProductNumber = (product) => {
    return request({
        url: '/product/number',
        method: 'PUT',
        data: product
    })
}

export const queryProduct = (id) => {
    return request({
        url: '/product/' + id,
        method: 'GET',
    })
}

export const queryProducts = (model, pageData) => {
    return request({
        url: '/products',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}