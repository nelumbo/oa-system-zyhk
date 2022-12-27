import request from './base'

export const addVendor = (vendor) => {
    return request({
        url: '/vendor',
        method: 'POST',
        data: vendor
    })
}

export const editVendor = (vendor) => {
    return request({
        url: '/vendor',
        method: 'PUT',
        data: vendor
    })
}

export const queryVendors = (model, pageData) => {
    return request({
        url: '/vendors',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryAllVendor = () => {
    return request({
        url: '/vendors',
        method: 'GET',
    })
}