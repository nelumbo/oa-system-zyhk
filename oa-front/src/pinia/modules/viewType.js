import { defineStore } from 'pinia'
import { ref } from 'vue'

const WIDTH = 720

export const viewType = defineStore('viewType', () => {
    const isMobbile = ref(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini|Mobile/i.test(navigator.userAgent) ? true : false)
    function compareWidth() {
        if (document.body.clientWidth > WIDTH) {
            isMobbile.value = false
        } else {
            isMobbile.value = true
        }
    }

    return { isMobbile, compareWidth }
})