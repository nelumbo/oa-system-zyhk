import request from './base'

export const addRegion = (region) => {
    return request({
        url: '/region',
        method: 'POST',
        data: region
    })
}

export const editRegion = (region) => {
    return request({
        url: '/region',
        method: 'PUT',
        data: region
    })
}

export const queryRegions = (model, pageData) => {
    return request({
        url: '/regions',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryAllRegion = () => {
    return request({
        url: '/regions',
        method: 'GET',
    })
}