import request from './base'


export const queryProductCalls = (model, pageData) => {
    return request({
        url: '/productCalls',
        method: 'POST',
        data: model,
        params: {
            "pageSize": pageData.pageSize,
            "pageNo": pageData.pageNo,
        }
    })
}