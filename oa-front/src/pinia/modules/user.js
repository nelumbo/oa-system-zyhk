import { defineStore } from 'pinia'
import { reactive } from 'vue'

export const user = defineStore('user', () => {
    const my = reactive({
        id: 0,
        name: "空白",
        roles: [],
        permissions: []
    })

    function setMy(obj) {
        my.id = obj.id
        my.name = obj.name
    }
    return { my, setMy }
})