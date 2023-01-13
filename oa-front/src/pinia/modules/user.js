import { defineStore } from 'pinia'
import { reactive } from 'vue'

export const user = defineStore('user', () => {
    const my = reactive({
        id: 0,
        name: "空白",
        officeID: null,
        roles: [],
        pids: []
    })

    function setMy(obj) {
        my.id = obj.id
        my.name = obj.name
        my.officeID = obj.officeID
        my.roles = obj.roles
        my.pids = obj.pids
    }
    return { my, setMy }
})