<template>
    <div>
        <el-divider content-position="left">
            <h2>办事处管理</h2>
        </el-divider>
        <officeManage />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>员工管理</h2>
        </el-divider>
        <employeeManage />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>职位管理</h2>
        </el-divider>
        <roleManage />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>区域管理</h2>
        </el-divider>
        <regionManage />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>合同签订单位管理</h2>
        </el-divider>
        <vendorManage />
    </div>
</template>

<script setup>
import officeManage from './OfficeManage.vue'
import employeeManage from './EmployeeManage.vue'
import roleManage from './RoleManage.vue'
import regionManage from './RegionManage.vue'
import vendorManage from './VendorManage.vue'
</script>