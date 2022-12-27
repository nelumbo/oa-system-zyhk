import request from './base'

export const addProductType = (productType) => {
    return request({
        url: '/productType',
        method: 'POST',
        data: productType
    })
}

export const editProductType = (productType) => {
    return request({
        url: '/productType',
        method: 'PUT',
        data: productType
    })
}

export const queryProductTypes = (model, pageData) => {
    return request({
        url: '/productTypes',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryAllProductType = () => {
    return request({
        url: '/productTypes',
        method: 'Get',
    })
}