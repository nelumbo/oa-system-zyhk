<template>
    <div>
        <el-divider content-position="left" v-if="user().my.pids.includes('51')">
            <h2>办事处管理</h2>
        </el-divider>
        <officeManage />
    </div>
    <div style="margin-top: 30px;" v-if="user().my.pids.includes('58') || user().my.pids.includes('59')">
        <el-divider content-position="left">
            <h2>员工管理</h2>
        </el-divider>
        <employeeManage />
    </div>
    <div style="margin-top: 30px;" v-if="user().my.pids.includes('66')">
        <el-divider content-position="left">
            <h2>职位管理</h2>
        </el-divider>
        <roleManage />
    </div>
    <div style="margin-top: 30px;" v-if="user().my.pids.includes('69')">
        <el-divider content-position="left">
            <h2>区域管理</h2>
        </el-divider>
        <regionManage />
    </div>
    <div style="margin-top: 30px;" v-if="user().my.pids.includes('72')">
        <el-divider content-position="left">
            <h2>合同签订单位管理</h2>
        </el-divider>
        <vendorManage />
    </div>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import officeManage from './OfficeManage.vue'
import employeeManage from './EmployeeManage.vue'
import roleManage from './RoleManage.vue'
import regionManage from './RegionManage.vue'
import vendorManage from './VendorManage.vue'
</script>