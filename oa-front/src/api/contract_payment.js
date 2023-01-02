import request from './base'

export const addPayment = (payment) => {
    return request({
        url: '/payment',
        method: 'POST',
        data: payment
    })
}

export const editPayment = (payment) => {
    return request({
        url: '/payment',
        method: 'PUT',
        data: payment
    })
}

export const finalPayment = (payment) => {
    return request({
        url: '/payment/final',
        method: 'PUT',
        data: payment
    })
}