<template>
    <div>
        <el-divider content-position="left">
            <h2>报销管理</h2>
        </el-divider>
        <expenseManage />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>采购管理</h2>
        </el-divider>
        <purchasingManage />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>投标保证金管理</h2>
        </el-divider>
        <bidBondManage />
    </div>
</template>

<script setup>
import expenseManage from './ExpenseManage.vue'
import bidBondManage from './BidBondManage.vue'
import purchasingManage from './PurchasingManage.vue'
</script>