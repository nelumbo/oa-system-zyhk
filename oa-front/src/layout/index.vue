<template>
    <el-container>
        <el-aside class="aside" :style="{ width: asideWidth }">
            <Sidebar />
        </el-aside>
        <el-main>
            <router-view />
        </el-main>
    </el-container>
</template>

<script setup>
import { queryMe } from "@/api/my";
import { onMounted, computed } from 'vue'
import { viewType } from '@/pinia/modules/viewType'
import { user } from '@/pinia/modules/user'
import Sidebar from './Sidebar/index.vue'

let asideWidth = computed(() => {
    if (viewType().isMobbile) {
        return '70px'
    } else {
        return '200px'
    }
})

onMounted(() => {
    queryMe().then((res) => {
        if(res.status == 1){
            user().setMy(res.data)
        }else{
            localStorage.removeItem("token");
            window.location.href = '/login'
        }
    })
    window.onresize = () => {
        viewType().compareWidth()
    }
})


</script>

<style>
.aside {
    background-color: #545c64;
}
</style>