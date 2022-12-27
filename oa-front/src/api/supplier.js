import request from './base'

export const addSupplier = (supplier) => {
    return request({
        url: '/supplier',
        method: 'POST',
        data: supplier
    })
}

export const editSupplier = (supplier) => {
    return request({
        url: '/supplier',
        method: 'PUT',
        data: supplier
    })
}

export const querySuppliers = (model, pageData) => {
    return request({
        url: '/suppliers',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}