import request from './base'

export const savePurchasing = (purchasing) => {
    return request({
        url: '/purchasing/save',
        method: 'POST',
        data: purchasing
    })
}

export const submitPurchasing = (purchasing) => {
    return request({
        url: '/purchasing/submit',
        method: 'PUT',
        data: purchasing
    })
}

export const addPurchasing = (purchasing) => {
    return request({
        url: '/purchasing',
        method: 'POST',
        data: purchasing
    })
}

export const delPurchasing = (id) => {
    return request({
        url: '/purchasing/' + id,
        method: 'DELETE',
    })
}

export const approvePurchasing = (purchasing) => {
    return request({
        url: '/purchasing/approve',
        method: 'PUT',
        data: purchasing
    })
}

export const finalPurchasingProduct = (purchasing) => {
    return request({
        url: '/purchasing/product/final',
        method: 'PUT',
        data: purchasing
    })
}

export const finalPurchasingPay = (purchasing) => {
    return request({
        url: '/purchasing/pay/final',
        method: 'PUT',
        data: purchasing
    })
}

export const finalPurchasingInvoice = (purchasing) => {
    return request({
        url: '/purchasing/invoice/final',
        method: 'PUT',
        data: purchasing
    })
}

export const finalPurchasing = (purchasing) => {
    return request({
        url: '/purchasing/final',
        method: 'PUT',
        data: purchasing
    })
}

export const queryPurchasing = (id) => {
    return request({
        url: '/purchasing/' + id,
        method: 'GET',
    })
}

export const queryPurchasings = (model, pageData) => {
    return request({
        url: '/purchasings',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryAllPurchasing = (model) => {
    return request({
        url: '/purchasing/all',
        method: 'POST',
        data: model,
    })
}