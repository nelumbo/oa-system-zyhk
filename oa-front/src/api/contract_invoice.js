import request from './base'

export const addInvoice = (invoice) => {
    return request({
        url: '/invoice',
        method: 'POST',
        data: invoice
    })
}

export const delInvoice = (id) => {
    return request({
        url: '/invoice/' + id,
        method: 'DELETE',
    })
}

export const editInvoice = (invoice) => {
    return request({
        url: '/invoice',
        method: 'PUT',
        data: invoice
    })
}