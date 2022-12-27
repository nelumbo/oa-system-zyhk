<template>
    <div>
        <el-divider content-position="left">
            <h2>办事处金额变更日志</h2>
        </el-divider>
        <officeLog />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>员工金额变更日志</h2>
        </el-divider>
        <employeeLog />
    </div>
</template>

<script setup>
import officeLog from './OfficeLog.vue'
import employeeLog from './EmployeeLog.vue'
</script>