import request from './base'

export const addEmployee = (employee) => {
    return request({
        url: '/employee',
        method: 'POST',
        data: employee
    })
}

export const delEmployee = (id) => {
    return request({
        url: '/employee/' + id,
        method: 'DELETE',
    })
}

export const editEmployeeBase = (employee) => {
    return request({
        url: '/employee/base',
        method: 'PUT',
        data: employee
    })
}

export const editEmployeeExpense = (employee) => {
    return request({
        url: '/employee/expense',
        method: 'PUT',
        data: employee
    })
}

export const editEmployeeOffice = (employee) => {
    return request({
        url: '/employee/office',
        method: 'PUT',
        data: employee
    })
}

export const queryEmployee = (id) => {
    return request({
        url: '/employee/' + id,
        method: 'GET',
    })
}

export const queryEmployees = (model, pageData) => {
    return request({
        url: '/employees',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}

export const queryAllEmployee = (model) => {
    return request({
        url: '/employee/all',
        method: 'POST',
        data: model,
    })
}

export const resetPwd = (id) => {
    return request({
        url: '/employee/reset/' + id,
        method: 'PUT',
    })
}