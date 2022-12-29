import request from './base'

export const addPredesign = (predesign) => {
    return request({
        url: '/predesign',
        method: 'POST',
        data: predesign
    })
}

export const delPredesign = (id) => {
    return request({
        url: '/predesign/' + id,
        method: 'DELETE',
    })
}

export const editPredesign = (predesign) => {
    return request({
        url: '/predesign',
        method: 'PUT',
        data: predesign
    })
}

export const approvePredesign = (predesign) => {
    return request({
        url: '/predesign/approve',
        method: 'PUT',
        data: predesign
    })
}

export const queryPredesign = (id) => {
    return request({
        url: '/predesign/' + id,
        method: 'GET',
    })
}

export const queryPredesigns = (model, pageData) => {
    return request({
        url: '/predesigns',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}