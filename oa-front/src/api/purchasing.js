import request from './base'

export const addPurchasing = (purchasing) => {
    return request({
        url: '/purchasing',
        method: 'POST',
        data: purchasing
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