import { defineStore } from 'pinia'
import { reactive } from 'vue'

export const user = defineStore('user', () => {
    const my = reactive({
        id: 0,
        name: "空白",
        officeID: null,
        roles: [],
        pids: [],
        urls: [],
    })

    function setMy(obj) {
        my.id = obj.id
        my.name = obj.name
        my.officeID = obj.officeID
        my.roles = obj.roles
        my.pids = obj.pids
        my.urls = obj.urls
    }
    return { my, setMy }
})