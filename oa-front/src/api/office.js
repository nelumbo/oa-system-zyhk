import request from './base'

export const addOffice = (office) => {
    return request({
        url: '/office',
        method: 'POST',
        data: office
    })
}

export const delOffice = (id) => {
    return request({
        url: '/office/' + id,
        method: 'DELETE',
    })
}

export const editOffice = (office) => {
    return request({
        url: '/office',
        method: 'PUT',
        data: office
    })
}

export const queryOffice = (id) => {
    return request({
        url: '/office/' + id,
        method: 'GET',
    })
}

export const queryOffices = (model, pageData) => {
    return request({
        url: '/offices',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryAllOffice = () => {
    return request({
        url: '/offices',
        method: 'Get',
    })
}
