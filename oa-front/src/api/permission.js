import request from './base'

export const queryAllPermission = () => {
    return request({
        url: '/permissions',
        method: 'Get',
    })
}