<template>
    <div>
        <el-divider content-position="left">
            <h2>产品管理</h2>
        </el-divider>
        <productManage />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>供应商管理</h2>
        </el-divider>
        <supplierManage />
    </div>
    <!-- <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>采购计划</h2>
        </el-divider>
        <planManage />
    </div> -->
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>产品类型管理</h2>
        </el-divider>
        <productTypeManage />
    </div>
</template>

<script setup>
import productManage from './ProductManage.vue'
import supplierManage from './SupplierManage.vue'
import planManage from './PlanManage.vue'
import productTypeManage from './ProductTypeManage.vue'
</script>