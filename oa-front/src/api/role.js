import request from './base'

export const addRole = (role) => {
    return request({
        url: '/role',
        method: 'POST',
        data: role
    })
}

export const delRole = (id) => {
    return request({
        url: '/role/' + id,
        method: 'DELETE',
    })
}

export const editRole = (role) => {
    return request({
        url: '/role',
        method: 'PUT',
        data: role
    })
}

export const queryRole = (id) => {
    return request({
        url: '/role/' + id,
        method: 'GET',
    })
}

export const queryRoles = (model, pageData) => {
    return request({
        url: '/roles',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryAllRole = () => {
    return request({
        url: '/roles',
        method: 'GET',
    })
}