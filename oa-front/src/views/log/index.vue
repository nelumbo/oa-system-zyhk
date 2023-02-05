<template>
    <div v-if="user().my.pids.includes('85') || user().my.pids.includes('86')">
        <el-divider content-position="left">
            <h2>办事处金额变更日志</h2>
        </el-divider>
        <officeLog />
    </div>
    <div style="margin-top: 30px;" v-if="user().my.pids.includes('87') || user().my.pids.includes('88')">
        <el-divider content-position="left">
            <h2>员工金额变更日志</h2>
        </el-divider>
        <employeeLog />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>产品数量变更日志</h2>
        </el-divider>
        <productLog />
    </div>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import officeLog from './OfficeLog.vue'
import employeeLog from './EmployeeLog.vue'
import productLog from './ProductLog.vue'
</script>