import request from './base'

export const saveContract = (bidbond) => {
    return request({
        url: '/contract/save',
        method: 'POST',
        data: bidbond
    })
}

export const addContract = (contract) => {
    return request({
        url: '/contract',
        method: 'POST',
        data: contract
    })
}

export const delContract = (id) => {
    return request({
        url: '/contract/' + id,
        method: 'DELETE',
    })
}

export const approveContract = (contract) => {
    return request({
        url: '/contract/approve',
        method: 'PUT',
        data: contract
    })
}

export const finalContract = (contract) => {
    return request({
        url: '/contract/final',
        method: 'PUT',
        data: contract
    })
}

export const resetContract = (contract) => {
    return request({
        url: '/contract/reset',
        method: 'PUT',
        data: contract
    })
}

export const rejectContract = (contract) => {
    return request({
        url: '/contract/reject',
        method: 'PUT',
        data: contract
    })
}

export const queryContract = (id) => {
    return request({
        url: '/contract/' + id,
        method: 'GET',
    })
}

export const queryContracts = (model, pageData) => {
    return request({
        url: '/contracts',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}