import request from './base'

export const addTask = (task) => {
    return request({
        url: '/task',
        method: 'POST',
        data: task
    })
}

export const distributeTask = (task) => {
    return request({
        url: '/task/distribute',
        method: 'PUT',
        data: task
    })
}

export const nextTask = (task) => {
    return request({
        url: '/task/next',
        method: 'PUT',
        data: task
    })
}

export const resetTask = (task) => {
    return request({
        url: '/task/reset',
        method: 'PUT',
        data: task
    })
}

export const rejectTask = (task) => {
    return request({
        url: '/task/reject',
        method: 'PUT',
        data: task
    })
}

export const delTask = (id) => {
    return request({
        url: '/task/' + id,
        method: 'DELETE',
    })
}