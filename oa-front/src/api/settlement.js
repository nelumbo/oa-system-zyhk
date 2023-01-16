import request from './base'

export const iceGet = () => {
    return request({
        url: '/ice',
        method: 'GET',
    })
}

export const iceStart = () => {
    return request({
        url: '/ice/start',
        method: 'PUT',
    })
}

export const iceEnd = () => {
    return request({
        url: '/ice/end',
        method: 'PUT',
    })
}

export const iceSubmit = (model) => {
    return request({
        url: '/ice/submit',
        method: 'PUT',
        data: model
    })
}

export const settSubmit = (model) => {
    return request({
        url: '/sett/submit',
        method: 'PUT',
        data: model
    })
}

export const settApprove = (model) => {
    return request({
        url: '/sett/approve',
        method: 'PUT',
        data: model
    })
}