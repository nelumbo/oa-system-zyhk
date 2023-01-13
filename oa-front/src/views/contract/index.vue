<template>
    <div>
        <el-divider content-position="left" v-if="user().my.pids.includes('17') || user().my.pids.includes('18')">
            <h2>合同管理</h2>
        </el-divider>
        <contractManage />
    </div>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import contractManage from './ContractManage.vue'
</script>