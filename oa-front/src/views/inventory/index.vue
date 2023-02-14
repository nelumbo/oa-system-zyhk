<template>
    <div>
        <el-divider content-position="left">
            <h2>任务管理</h2>
        </el-divider>
        <taskManage />
    </div>
    <div>
        <el-divider content-position="left">
            <h2>采购入库管理</h2>
        </el-divider>
        <purchasingManage />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>库存警报</h2>
        </el-divider>
        <productCallManage />
    </div>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import taskManage from './TaskManage.vue'
import purchasingManage from './PurchasingManage.vue'
import productCallManage from './ProductCallManage.vue'
</script>