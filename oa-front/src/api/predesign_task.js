import request from './base'

export const submitPredesignTask = (predesignTask) => {
    return request({
        url: '/predesignTask/submit',
        method: 'PUT',
        data: predesignTask
    })
}

export const approvePredesignTask = (predesignTask) => {
    return request({
        url: '/predesignTask/approve',
        method: 'PUT',
        data: predesignTask
    })
}

export const queryPredesignTasks = (model, pageData) => {
    return request({
        url: '/predesignTasks',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}
